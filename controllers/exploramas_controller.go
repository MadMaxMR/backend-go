package controllers

import (
	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
)

// GetAreaCarrerasByUni retorna todas las areas de una "universidad" incluido las carreras de sus areas
func GetAreaCarrerasByUni(w http.ResponseWriter, req *http.Request) {

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
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró registros para el filtro seleccionado123")
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
			"per.id as id_perest, per.ptjmin, per.ptjmax, per.anio, per.vacantes, per.modalidad").
		Joins("INNER JOIN areas are on are.id_uni = uni.id").
		Joins("INNER JOIN carreras car on car.cod_area = are.id").
		Joins("INNER JOIN perfil_postulantes per on per.carreras_id = car.id").
		Where("uni.id like ? and are.id like ? and car.id::text like ?", idUniversidad, idArea, idCarrera).
		Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("id_carrera").Find(&result)

	result2.Data = result

	handler.SendSuccess(w, req, http.StatusOK, result2)
}

//GetInfoMas Devuelve la info adicional por carreras
func GetInfoMas(w http.ResponseWriter, req *http.Request) {
	// var result []map[string]interface{}

	// db := database.GetConnection()
	// dbc, _ := db.DB()
	// defer dbc.Close()

}
