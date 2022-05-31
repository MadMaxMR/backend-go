package controllers

import (
	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"net/http"
)

func GetEvaluaciones(w http.ResponseWriter, req *http.Request) {
	eval := []modelos.Evaluaciones{}
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&eval, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func SaveEvaluaciones(w http.ResponseWriter, req *http.Request) {
	eval := modelos.Evaluaciones{}
	err := auth.ValidateBody(req, &eval)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateEvaluaciones(&eval)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	modelo, err := database.Create(&eval)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}
