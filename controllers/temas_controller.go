package controllers

import (
	"net/http"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"

	"github.com/gorilla/mux"
)

func GetAllTemas(w http.ResponseWriter, req *http.Request) {
	temas := []modelos.Temas{}
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&temas, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetTema(w http.ResponseWriter, req *http.Request) {
	tema := modelos.Temas{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Get(&tema, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func SaveTema(w http.ResponseWriter, req *http.Request) {
	tema := modelos.Temas{}
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
	tema := modelos.Temas{}
	id := mux.Vars(req)["id"]
	message, err := database.Delete(&tema, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, message)
}

func UpdateTema(w http.ResponseWriter, req *http.Request) {
	tema := modelos.Temas{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Update(&tema, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

//GetTemaByCurso retorna todos los temas de un curso
func GetTemaByCurso(w http.ResponseWriter, req *http.Request) {
	temas := []modelos.Temas{}
	curso := modelos.Cursos{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	_, err := database.Get(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	result := db.Where("id_curso = ?", id).Find(&temas)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró temas para el curso: "+curso.Nombre_Curso)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, temas)
}

//GetTemasVideos retorna todos los temas de un curso incluido todos los videos y evaluaciones pertenecientes a los temas
func GetTemasVideos(w http.ResponseWriter, req *http.Request) {
	temas := []modelos.Temas{}
	curso := modelos.Cursos{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	_, err := database.Get(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	result := db.Model(&temas).Where("id_curso = ?", id).Preload("Videos").Preload("Evaluaciones").Find(&temas)

	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró temas para el curso: "+curso.Nombre_Curso)
		return
	}

	handler.SendSuccess(w, req, http.StatusOK, temas)
}
