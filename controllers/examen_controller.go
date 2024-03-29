package controllers

import (
	"fmt"
	"math"
	"time"

	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// SaveExamens controller para crear y guardar un nuevo examen con preguntas y respuestas
func SaveExamens(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	err := auth.ValidateBody(req, &examen)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateExamen(&examen)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	examen.FechaCreacion = time.Now().Add(time.Hour - 6)
	_, err = database.Create(&examen)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusCreated, examen)
}

func GetAllExamens(w http.ResponseWriter, req *http.Request) {
	examen := []modelos.Examens{}
	page := req.URL.Query().Get("page")
	pageSizes := req.URL.Query().Get("pageSize")

	if page == "" {
		page = "1"
	}
	if pageSizes == "" {
		pageSizes = "10"
	}

	pageInt, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(pageSizes)

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	type Result struct {
		Page     string
		Prev     bool
		Next     bool
		Total    int
		Examenes []modelos.Examens
	}
	result2 := Result{}

	result2.Page = page
	result2.Next = true
	if pageInt == 1 {
		result2.Prev = false
	}
	if pageInt > 1 {
		result2.Prev = true
	}

	_, _ = database.GetAll(&examen, "")
	if len(examen)%pageSize == 0 {
		result2.Total = len(examen) / pageSize
	} else {
		result2.Total = (len(examen) / pageSize) + 1
	}

	if pageInt == result2.Total {
		result2.Next = false
	}

	result := db.Model(&examen).Select("*").Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("id DESC").Find(&examen)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No se encontró exámenes")
		return
	}

	result2.Examenes = examen
	handler.SendSuccess(w, req, http.StatusOK, result2)
}

func GetExamenById(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	result := db.Model(&examen).Where("id = ?", id).Preload("PreguntaExamens", func(db *gorm.DB) *gorm.DB {
		return db.Order("pregunta_examens.id ASC")
	}).Preload("PreguntaExamens.RespuestaExs", func(db *gorm.DB) *gorm.DB {
		return db.Order("respuesta_exs.id ASC")
	}).Find(&examen)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró examenes con el id: "+id)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, examen)
}

func UpdateExamen(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	id := mux.Vars(req)["id"]

	err := auth.ValidateBody(req, &examen)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	_, err = database.Update(&examen, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, "Examen actualizado correctamente")
}

func DeleteExamen(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	examenPreguntas := []modelos.ExamenPreguntas{}

	id := mux.Vars(req)["id"]

	db := database.GetConnection()

	dbc, _ := db.DB()
	defer dbc.Close()

	db.Raw("SELECT * FROM examen_preguntas where examens_id = ?", id).Scan(&examenPreguntas)
	if len(examenPreguntas) != 0 {
		handler.SendFail(w, req, http.StatusNotAcceptable, "Invalid-El examen tiene preguntas")
		return
	}
	message, err := database.Delete(&examen, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, message)
}

func GetPoints(w http.ResponseWriter, req *http.Request) {

	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	historial := modelos.MisExamenes{}
	points := modelos.Result{Resultado: make(map[string]string), Solucion: make(map[string]uint)}
	result := map[string]interface{}{}
	var solution, answers string
	idExamen := req.URL.Query().Get("idExamen")
	if idExamen == "" {
		idExamen = "2"
	}
	idExamenInt, _ := strconv.Atoi(idExamen)

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()
	correct, incorrect, note := 0, 0, 0.0

	err = auth.ValidateBody(req, &result)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	examen := modelos.Examens{}
	db.Model(&examen).Where("id = ? ", idExamenInt).Find(&examen)

	for i := 1; i < (len(result)/2 + 1); i++ {
		respuesta := modelos.RespuestaExs{}
		ponderado := modelos.Ponderacion{}
		pregunta := modelos.PreguntaExamens{}

		val := strconv.Itoa(i)
		rest := db.Model(&respuesta).Where("pregunta_examens_id = ? and valor = 'true'", result["id_pregunta"+val]).Find(&respuesta)
		if rest.RowsAffected == 0 {
			handler.SendFail(w, req, http.StatusBadRequest, "No hay alternativa correcta")
			return
		}
		solution += strconv.Itoa(int(respuesta.ID)) + "-"
		answers += fmt.Sprintf("%v", result["id_respuesta"+val]) + "-"
		if result["id_respuesta"+val] != float64(0) {
			if result["id_respuesta"+val] == float64(respuesta.ID) {
				db.Model(&pregunta).Where("id = ? ", result["id_pregunta"+val]).Find(&pregunta)
				db.Model(&ponderado).Where("cursos_id = ? and cod_area = ?", pregunta.CursosId, examen.AreasId).Find(&ponderado)
				points.Resultado["pregunta"+val] = "Correcto"
				points.Solucion["pregunta"+val] = respuesta.ID
				note = note + ponderado.Ponderacion
				correct++
			} else {
				points.Resultado["pregunta"+val] = "Incorrecto"
				points.Solucion["pregunta"+val] = respuesta.ID
				incorrect++
			}
		} else if result["id_respuesta"+val] == float64(0) {
			points.Resultado["pregunta"+val] = "No contestada"
			points.Solucion["pregunta"+val] = respuesta.ID
			incorrect++
		}
	}
	points.Correct = correct
	points.Incorrect = incorrect
	points.Nota = math.Round(((note*20)/float64(examen.LimitePreguntas))*100) / 100

	historial.AreasId = examen.AreasId
	historial.UniversidadsId = examen.Id_Uni
	historial.ExamensId = uint(idExamenInt)
	historial.Fecha_Examen = time.Now().Add(time.Hour - 6)
	historial.Nota = points.Nota
	if points.Nota < 10.5 {
		historial.Condicion = "Desaprobado"
	} else {
		historial.Condicion = "Aprobado"
	}

	iduser, _ := strconv.Atoi(tk.Id_Usuario)
	historial.UsuarioId = uint(iduser)

	db.Create(&historial)
	//fmt.Println("answers: ", answers)	fmt.Println("solution: ", solution)
	db.Save(&modelos.HistorialExamens{UsuarioId: uint(iduser), Id_Examen: uint(idExamenInt),
		Fecha_Examen: time.Now().Add(time.Hour - 6), Nota_Max: points.Nota,
		Respuestas: answers, Solucion: solution})

	handler.SendSuccess(w, req, http.StatusOK, points)
}

// a := "10-15-119-5-8-10-55-"
// 	cadena := strings.Split(a, "-")
// 	for i := 0; i < len(cadena)-1; i++ {
// 		fmt.Print("\n valor ", i, ":", cadena[i])
// 	}

// GetExamensPregByArea retorna todos los examenes de un area con sus preguntas y alternativas
func GetExamensPregByArea(w http.ResponseWriter, req *http.Request) {
	preguntas := []modelos.PreguntaExamens{}
	examen := modelos.Examens{}
	id := mux.Vars(req)["idExamen"]

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	resultQ := db.Model(&examen).Where("id = ?", id).Find(&examen)
	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No existe el examen - "+id)
		return
	}
	if examen.LimitePreguntas != examen.CantidadPreguntas {
		handler.SendFail(w, req, http.StatusBadRequest, "El examen no tiene preguntas completas - "+fmt.Sprint(examen.LimitePreguntas)+"/"+fmt.Sprint(examen.CantidadPreguntas))
		return
	}
	resultQ = db.Preload("RespuestaExs").
		Where("ex.examens_id = ?", examen.ID).
		Select("pregunta_examens.id,ex.examens_id,pregunta_examens.enunciado1,pregunta_examens.grafico," +
			"pregunta_examens.enunciado2,pregunta_examens.enunciado3,row_number() OVER () AS num_question," +
			"pregunta_examens.cursos_id,pregunta_examens.temas_id,pregunta_examens.nivel").
		Joins("INNER JOIN examen_preguntas ex on ex.pregunta_examens_id = pregunta_examens.id").Order("pregunta_examens.id DESC").
		Find(&preguntas)
	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se agregó preguntas al examen")
		return
	}
	examen.PreguntaExamens = preguntas

	handler.SendSuccess(w, req, http.StatusOK, examen)
}

/*
SELECT pre.id,ex.examens_id,pre.enunciado1,pre.grafico,pre.enunciado2,pre.enunciado3,ex.num_question,pre.cursos_id,
pre.temas_id,pre.nivel
FROM examen_preguntas ex
INNER JOIN pregunta_examens  pre on ex.pregunta_examens_id = pre.id
WHERE ex.examens_id = 1
*/
func GetModalidad(w http.ResponseWriter, req *http.Request) {
	type Modalidades struct {
		Name string
		Code string
	}
	modalidad := []Modalidades{}

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	db.Raw("SELECT name,code FROM modalidades").Scan(&modalidad)
	if len(modalidad) == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No hay modalidades registradas")
		return
	}

	handler.SendSuccess(w, req, http.StatusOK, modalidad)
}

func GetExamensbyAnio(w http.ResponseWriter, req *http.Request) {

	var años []map[string]interface{}
	examenes := make(map[string]interface{})
	id := mux.Vars(req)["id"]
	tipex := req.URL.Query().Get("tipex")

	if tipex == "" {
		tipex = "Admision"
	}

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	db.Table("examens").Select("anio").Where("areas_id = $1 and tipo_examen = $2", id, tipex).Group("anio").Find(&años)

	if len(años) == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No se encontró examenes para el area o tipo ingresado")
		return
	}

	for _, v := range años {
		var anio string
		anio = v["anio"].(string)
		var Examens []modelos.Examens
		res := db.Where("anio= $1 and areas_id = $2 and limite_preguntas = cantidad_preguntas", anio, id).Find(&Examens)
		if res.RowsAffected == 0 {
			continue
		}
		examenes[anio] = Examens
	}

	if len(examenes) == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No se encontró examenes para el area o tipo ingresado")
		return
	}

	handler.SendSuccess(w, req, http.StatusOK, examenes)
}

// GetPreguntasforETA devuelve 10 preguntas para FAS TEST 7 de tipo ETA y 3 de tipo Admision
func GetFastTest(w http.ResponseWriter, req *http.Request) {
	preguntas := []modelos.PreguntaExamens{}
	idCurso := req.URL.Query().Get("idCurso")
	idTema := req.URL.Query().Get("idTema")
	total := req.URL.Query().Get("total")
	if total == "" {
		total = "10"
	}
	totalInt, _ := strconv.Atoi(total)
	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	err := db.Preload("RespuestaExs").Scopes(func(db *gorm.DB) *gorm.DB {
		if idTema == "" && idCurso != "" {
			return db.Where("cursos_id = ?", idCurso)
		} else if idCurso != "" && idTema != "" {
			return db.Where("temas_id = ?", idTema)
		} else {
			return db.Where("cursos_id <> 0")
		}
	}).Select("pregunta_examens.id,pregunta_examens.enunciado1,pregunta_examens.grafico," +
		"pregunta_examens.enunciado2,pregunta_examens.enunciado3,row_number() OVER (order by random()) AS num_question," +
		"pregunta_examens.cursos_id,pregunta_examens.temas_id,pregunta_examens.nivel").
		Limit(totalInt).Order("random()").
		Find(&preguntas).Error
	if err != nil {
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
		return
	}
	if len(preguntas) < totalInt {
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontraron preguntas suficientes para el tema seleccionado")
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, preguntas)
}

func GetPointsFastTest(w http.ResponseWriter, req *http.Request) {
	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	historial := modelos.HistorialFastest{}
	points := modelos.Result{Resultado: make(map[string]string), Solucion: make(map[string]uint)}
	result := map[string]interface{}{}
	var solution, answers string

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()
	correct, incorrect, note := 0, 0, 0.0

	err = auth.ValidateBody(req, &result)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	for i := 1; i < (len(result)/2 + 1); i++ {
		respuesta := modelos.RespuestaExs{}
		ponderado := modelos.PonderacionFastest{}
		pregunta := modelos.PreguntaExamens{}

		val := strconv.Itoa(i)
		rest := db.Model(&respuesta).Where("pregunta_examens_id = ? and valor = 'true'", result["id_pregunta"+val]).Find(&respuesta)
		if rest.RowsAffected == 0 {
			handler.SendFail(w, req, http.StatusBadRequest, "No hay alternativa correcta")
			return
		}
		solution += strconv.Itoa(int(respuesta.ID)) + "-"
		answers += fmt.Sprintf("%v", result["id_respuesta"+val]) + "-"
		if result["id_respuesta"+val] != float64(0) {
			if result["id_respuesta"+val] == float64(respuesta.ID) {
				db.Model(&pregunta).Where("id = ? ", result["id_pregunta"+val]).Find(&pregunta)
				db.Model(&ponderado).Where("cursos_id = ?", pregunta.CursosId).Find(&ponderado)
				points.Resultado["pregunta"+val] = "Correcto"
				points.Solucion["pregunta"+val] = respuesta.ID
				note = note + ponderado.Ponderacion
				correct++
			} else {
				points.Resultado["pregunta"+val] = "Incorrecto"
				points.Solucion["pregunta"+val] = respuesta.ID
				incorrect++
			}
		} else if result["id_respuesta"+val] == float64(0) {
			points.Resultado["pregunta"+val] = "No contestada"
			points.Solucion["pregunta"+val] = respuesta.ID
			incorrect++
		}

		historial.TemasId = pregunta.TemasId
		historial.CursosId = pregunta.CursosId
	}
	points.Correct = correct
	points.Incorrect = incorrect
	points.Nota = note

	historial.Fecha_Examen = time.Now().Add(time.Hour - 6)
	historial.Nota = points.Nota
	iduser, _ := strconv.Atoi(tk.Id_Usuario)
	historial.UsuarioId = uint(iduser)

	db.Create(&historial)

	handler.SendSuccess(w, req, http.StatusOK, points)

}
