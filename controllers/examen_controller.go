package controllers

import (
	"net/http"
	// "github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/gorilla/mux"
)

func GetAllExamens(w http.ResponseWriter, req *http.Request) {
	examen := []modelos.Examens{}
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&examen, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetExamen(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Get(&examen, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}
func GetExamensPregByArea(w http.ResponseWriter, req *http.Request) {
	examen := []modelos.Examens{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	result := db.Model(&examen).Where("areas_id = ?", id).Preload("PreguntaExamens").Preload("RespuestaExs").Find(&examen)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró examenes para el area: "+id)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, examen)
}
