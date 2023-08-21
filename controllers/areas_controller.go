package controllers

import (
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"

	"net/http"

	"github.com/gorilla/mux"
)

func GetAllAreas(w http.ResponseWriter, req *http.Request) {
	areas := []modelos.Areas{}
	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()
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
	dbc, _ := db.DB()
	defer dbc.Close()

	result := db.Where("id_uni = ?", id).Find(&areas)
	if result.Error != nil {
		handler.SendFail(w, req, http.StatusBadRequest, result.Error.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, areas)
}
