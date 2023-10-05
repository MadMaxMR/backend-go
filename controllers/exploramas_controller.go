package controllers

import (
	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
)

// GetAreaCarrerasByUni retorna todas las areas de una "universidad" incluido las carreras de sus areas
func GetAreaCarrerasByUni(w http.ResponseWriter, req *http.Request) {
	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	var result []map[string]interface{}
	result2 := modelos.Resultado{}
	idCarrera := req.URL.Query().Get("idCarrera")
	idArea := req.URL.Query().Get("idArea")
	idUniversidad := req.URL.Query().Get("idUniversidad")
	page := req.URL.Query().Get("page")
	pageSizes := req.URL.Query().Get("pageSize")

	if page == "" {
		page = "1"
	}
	if pageSizes == "" {
		pageSizes = "20"
	}
	pageInt, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(pageSizes)

	if idUniversidad == "" {
		idUniversidad = "%"
	}
	if idArea == "" {
		idArea = "%"
	}
	if idCarrera == "" {
		idCarrera = "%"
	}

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	resultQ := db.Table("universidads uni").
		Select("uni.id as id_uni, uni.nombre_uni").
		Joins("INNER JOIN areas are on are.id_uni = uni.id").
		Joins("INNER JOIN carreras car on car.cod_area = are.id").
		Joins("INNER JOIN perfil_postulantes per on per.carreras_id = car.id").
		Where("uni.id like ? and are.id like ? and car.id::text like ?", idUniversidad, idArea, idCarrera).Find(&result)

	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró registros para el filtro seleccionado")
		return
	}

	result2.Page = pageInt
	result2.Next = true
	if pageInt == 1 {
		result2.Prev = false
	}
	if pageInt > 1 {
		result2.Prev = true
	}

	if int(resultQ.RowsAffected)%pageSize == 0 {
		result2.Total = int(resultQ.RowsAffected) / pageSize
	} else {
		result2.Total = (int(resultQ.RowsAffected) / pageSize) + 1
	}

	if pageInt == result2.Total {
		result2.Next = false
	}

	result = nil
	db.Table("universidads uni").
		Select("uni.id as id_uni, uni.nombre_uni,"+
			"are.id as id_area, are.nombre_area,"+
			"car.id as id_carrera, car.nombre_carr, "+
			"per.id as id_perest, per.ptjmin, per.ptjmax, per.anio, per.vacantes, per.modalidad, "+
			"case when mex.nota is null then 0 else mex.nota END mi_nota").
		Joins("INNER JOIN areas are on are.id_uni = uni.id").
		Joins("INNER JOIN carreras car on car.cod_area = are.id").
		Joins("INNER JOIN perfil_postulantes per on per.carreras_id = car.id").
		Joins("LEFT JOIN (select usuario_id, areas_id, avg(nota)::float as nota from mis_examenes group by usuario_id,areas_id) mex "+
			"on mex.areas_id = are.id and mex.usuario_id = ?", tk.Id_Usuario).
		Where("uni.id like ? and are.id like ? and car.id::text like ?", idUniversidad, idArea, idCarrera).
		Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("id_carrera").Find(&result)

	if len(result) == 0 {
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró registros para el filtro seleccionado")
		return
	}

	result2.Data = result
	handler.SendSuccess(w, req, http.StatusOK, result2)
}

// GetInfoMas Devuelve la info adicional por carreras
func GetInfoMas(w http.ResponseWriter, req *http.Request) {
	// var result []map[string]interface{}

	// db := database.GetConnection()
	// dbc, _ := db.DB()
	// defer dbc.Close()

}

func GetIndicadores(w http.ResponseWriter, req *http.Request) {
	data := []map[string]interface{}{}
	idUni := req.URL.Query().Get("idUni")
	idArea := req.URL.Query().Get("idArea")
	idCurso := req.URL.Query().Get("idCurso")
	lAnio := req.URL.Query().Get("lAnio")
	rAnio := req.URL.Query().Get("rAnio")

	if idUni == "" || idArea == "" || idCurso == "" || lAnio == "" || rAnio == "" {
		handler.SendFail(w, req, http.StatusNotAcceptable, "Universidad, area, curso o años no pueden estar en blanco")
		return
	}

	leAnio, _ := strconv.Atoi(lAnio)
	riAnio, _ := strconv.Atoi(rAnio)

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	resultQ := db.Table("(?) as dat", db.Table("examens e").
		Select("e.id_uni,e.areas_id ,pe.cursos_id ,pe.temas_id,count(*) as total").
		Joins("INNER JOIN examen_preguntas ep on ep.examens_id  = e.id").
		Joins("INNER JOIN pregunta_examens pe on pe.id = ep.pregunta_examens_id").
		Where("e.tipo_examen  = 'Admision' and e.cantidad_preguntas = e.limite_preguntas and e.id_uni = $1 and pe.cursos_id = $2 "+
			" and e.areas_id = $3 and anio::integer between $4 and $5", idUni, idCurso, idArea, leAnio, riAnio).
		Group(" e.id_uni,e.areas_id,pe.cursos_id,pe.temas_id").
		Order("pe.cursos_id,pe.temas_id")).
		Select("dat.*,te.nombre_tema as nom_tema, SUM(dat.total) OVER (PARTITION BY dat.cursos_id) AS suma,(total/(SUM(dat.total) OVER (PARTITION BY dat.cursos_id)))*100 as prctje").
		Joins("LEFT join temas te on te.id = dat.temas_id").
		Order("dat.total ASC").
		Scan(&data)

	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusNoContent, "No hay datos para el filtro seleccionado")
	}
	handler.SendSuccess(w, req, http.StatusOK, data)
}

/*
select e.anio,e.id_uni,e.areas_id ,pe.cursos_id ,pe.temas_id,count(*) total from examens e
		inner join examen_preguntas ep on ep.examens_id  = e.id
		inner join pregunta_examens pe on pe.id = ep.pregunta_examens_id
		where e.tipo_examen  = 'Admision' and e.cantidad_preguntas = e.limite_preguntas
		group by e.anio, e.id_uni,e.areas_id,pe.cursos_id,pe.temas_id
		order by  e.anio, pe.cursos_id,pe.temas_id
*/
