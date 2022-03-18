package controllers

import (
	"API/auth"
	"API/database"
	"API/handler"
	"API/models"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func GetAllCursos(w http.ResponseWriter, req *http.Request) {
	curso := []models.Cursos{}
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&curso, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetCurso(w http.ResponseWriter, req *http.Request) {
	curso := models.Cursos{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Get(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusNotFound, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func SaveCurso(w http.ResponseWriter, req *http.Request) {
	curso := models.Cursos{}
	err := auth.ValidateBody(req, &curso)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateCurso(&curso)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	model, err := database.Create(&curso)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusCreated, model)
}

func DeleteCurso(w http.ResponseWriter, req *http.Request) {
	curso := models.Cursos{}
	id := mux.Vars(req)["id"]
	message, err := database.Delete(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, message)
}

func UpdateCurso(w http.ResponseWriter, req *http.Request) {
	curso := models.Cursos{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Update(req, &curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func UploadImage(w http.ResponseWriter, req *http.Request) {
	curso := models.Cursos{}

	db := database.GetConnection()
	defer db.Close()

	id := mux.Vars(req)["id"]

	db.Find(&curso, id)

	if curso.ID > 0 {
		file, hand, err := req.FormFile("image")
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al leer el archivo - "+err.Error())
			return
		}
		var extension = strings.Split(hand.Filename, ".")[1]
		var archivo string = "media/cursos/" + id + "." + extension

		f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al crear el archivo - "+err.Error())
			return
		}
		_, err = io.Copy(f, file)
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al copiar el archivo - "+err.Error())
			return
		}
		defer f.Close()

		curso.Image = string(id) + "." + extension
		db.Save(&curso)
		handler.SendSuccess(w, req, http.StatusOK, curso)
	} else {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontro el curso")
	}
}

func GetImage(w http.ResponseWriter, req *http.Request) {
	curso := models.Cursos{}
	id := mux.Vars(req)["id"]
	_, err := database.Get(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	OpenFile, err := os.Open("media/cursos/" + curso.Image)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encuentra la imagen -"+err.Error())
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al copiar el archivo - "+err.Error())
		return
	}
}
