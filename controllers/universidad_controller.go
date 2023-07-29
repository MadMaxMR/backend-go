package controllers

import (
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"

	"net/http"

	"github.com/gorilla/mux"
)

func GetAllUniversidads(w http.ResponseWriter, req *http.Request) {
	universidads := []modelos.Universidads{}
	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&universidads, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetUniversidad(w http.ResponseWriter, req *http.Request) {
	universidad := modelos.Universidads{}
	id := mux.Vars(req)["id"]

	modelo, err := database.Get(&universidad, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)

}
