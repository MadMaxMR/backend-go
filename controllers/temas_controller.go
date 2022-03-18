package controllers

import (
	"API/auth"
	"API/database"
	"API/handler"
	"API/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetAllTemas(w http.ResponseWriter, req *http.Request) {
	temas := []models.Temas{}
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&temas, page)
	if err != nil {

		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetTema(w http.ResponseWriter, req *http.Request) {
	tema := models.Temas{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Get(&tema, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func SaveTema(w http.ResponseWriter, req *http.Request) {
	tema := models.Temas{}
	err := auth.ValidateBody(req, &tema)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateTema(&tema)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	modelo, err := database.Create(&tema)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func DeleteTema(w http.ResponseWriter, req *http.Request) {
	tema := models.Temas{}
	id := mux.Vars(req)["id"]
	message, err := database.Delete(&tema, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, message)
}

func UpdateTema(w http.ResponseWriter, req *http.Request) {
	tema := models.Temas{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Update(req, &tema, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetTemaByCurso(w http.ResponseWriter, req *http.Request) {
	temas := []models.Temas{}
	curso := models.Cursos{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	_, err := database.Get(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	err = db.Where("cursos_id = ?", id).Find(&temas).Error
	if err != nil {
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, temas)
}
