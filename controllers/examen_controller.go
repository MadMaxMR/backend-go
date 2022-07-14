package controllers

import (
	//"fmt"
	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
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

//GetExamensPregByArea retorna todos los examenes de un area con sus preguntas y alternativas
func GetExamensPregByArea(w http.ResponseWriter, req *http.Request) {
	examen := []modelos.Examens{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	result := db.Model(&examen).Where("areas_id = ?", id).Preload("PreguntaExamens", func(db *gorm.DB) *gorm.DB {
		return db.Order("pregunta_examens.id ASC")
	}).Preload("PreguntaExamens.RespuestaExs", func(db *gorm.DB) *gorm.DB {
		return db.Order("respuesta_exs.id ASC")
	}).Find(&examen)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró examenes para el area: "+id)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, examen)
}

func SavePreguntaResp(w http.ResponseWriter, req *http.Request) {
	pregunta := modelos.PreguntaExamens{}
	err := auth.ValidateBody(req, &pregunta)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}
	modelo, err := database.Create(&pregunta)
	if err != nil {
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}
