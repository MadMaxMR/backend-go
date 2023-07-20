package controllers

import (
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/jinzhu/gorm"

	"net/http"

	"github.com/gorilla/mux"
)

func GetAllAreas(w http.ResponseWriter, req *http.Request) {
	areas := []modelos.Areas{}
	db := database.GetConnection()
	defer db.Close()
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&areas, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetArea(w http.ResponseWriter, req *http.Request) {
	area := modelos.Areas{}
	id := mux.Vars(req)["id"]

	modelo, err := database.Get(&area, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)

}

func GetAreaByUni(w http.ResponseWriter, req *http.Request) {
	areas := []modelos.Areas{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	result := db.Where("id_uni = ?", id).Find(&areas)
	if result.Error != nil {
		handler.SendFail(w, req, http.StatusBadRequest, result.Error.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, areas)
}

// GetAreaCarrerasByUni retorna todas las areas de una "universidad" incluido las carreras de sus areas
func GetAreaCarrerasByUni(w http.ResponseWriter, req *http.Request) {
	universidads := []modelos.Universidads{}
	//areas := []modelos.Area{}
	idCarrera := req.URL.Query().Get("idCarrera")
	idArea := req.URL.Query().Get("idArea")
	idUniversidad := req.URL.Query().Get("idUniversidad")
	//page := req.URL.Query().Get("page")

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
	defer db.Close()

	db.Debug().Where("id LIKE ?", idUniversidad).Preload("Area", "id LIKE ?", idArea).Preload("Area.Carreras", "id::text LIKE ?", idCarrera, func(db *gorm.DB) *gorm.DB {
		return db.Order("Carreras.nombre_carr ASC")
	}).Preload("Area.Carreras.PerfilPostulante").Find(&universidads)
	// if result.RowsAffected == 0 {
	// 	handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró areas para la universidad solicitada")
	// 	return
	// }

	handler.SendSuccess(w, req, http.StatusOK, universidads)
}
