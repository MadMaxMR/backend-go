package controllers

import (
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"

	"net/http"

	"github.com/gorilla/mux"
)

func GetAllCarreras(w http.ResponseWriter, req *http.Request) {
	carreras := []modelos.Carreras{}
	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&carreras, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetCarrera(w http.ResponseWriter, req *http.Request) {
	carrera := modelos.Carreras{}
	id := mux.Vars(req)["id"]

	modelo, err := database.Get(&carrera, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)

}

func GetCarreraByArea(w http.ResponseWriter, req *http.Request) {
	carreras := []modelos.Carreras{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	result := db.Where("cod_area = ?", id).Find(&carreras)
	if result.Error != nil {
		handler.SendFail(w, req, http.StatusBadRequest, result.Error.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, carreras)
}

func GetCarreraUni(w http.ResponseWriter, req *http.Request) {
	carreras := []modelos.Carreras{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	result := db.Where("id_uni = ?", id).Find(&carreras)
	if result.Error != nil {
		handler.SendFail(w, req, http.StatusBadRequest, result.Error.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, carreras)
}
