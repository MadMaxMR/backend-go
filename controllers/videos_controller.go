package controllers

import (
	"net/http"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"

	"github.com/gorilla/mux"
)

func GetAllVideos(w http.ResponseWriter, req *http.Request) {
	videos := []modelos.Videos{}
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&videos, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetVideo(w http.ResponseWriter, req *http.Request) {
	video := modelos.Videos{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Get(&video, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetVideoByTema(w http.ResponseWriter, req *http.Request) {
	video := []modelos.Videos{}
	id := mux.Vars(req)["id"]
	db := database.GetConnection()
	result := db.Where("Temas_Id = ?", id).Find(&video)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró videos para el tema : "+id)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, video)
}

func SaveVideo(w http.ResponseWriter, req *http.Request) {
	video := modelos.Videos{}
	err := auth.ValidateBody(req, &video)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateVideo(&video)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	modelo, err := database.Create(&video)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}
