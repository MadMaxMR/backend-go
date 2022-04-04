package controllers

import (
	//"errors"
	//"fmt"
	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	//"io"
	//"log"
	"net/http"
	//"os"
	//"strings"
	//"time"

	"github.com/gorilla/mux"
	//"golang.org/x/crypto/bcrypt"
)

func GetStudent(w http.ResponseWriter, req *http.Request) {
	student := modelos.Estudiantes{}
	id := mux.Vars(req)["id"]

	user, err := database.Get(&student, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, user)
}

func GetAllStudent(w http.ResponseWriter, r *http.Request) {
	students := []modelos.Estudiantes{}
	page := r.URL.Query().Get("page")
	user, err := database.GetAll(&students, page)
	if err != nil {
		handler.SendFail(w, r, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, r, http.StatusOK, user)
}

func SaveStudent(w http.ResponseWriter, r *http.Request) {
	student := modelos.Estudiantes{}
	usuario := modelos.Usuarios{}
	err1, err2 := auth.ValidateBody(r, &student), auth.ValidateBody(r, &usuario)
	if err1 != nil || err2 != nil {
		if err1 != nil {
			handler.SendFail(w, r, http.StatusBadRequest, err1.Error())
			return
		} else {
			handler.SendFail(w, r, http.StatusBadRequest, err2.Error())
			return
		}
	}
	err1, err2 = auth.ValidateStudent(&student), auth.ValidateUsuario(&usuario)
	if err1 != nil || err2 != nil {
		if err1 != nil {
			handler.SendFail(w, r, http.StatusBadRequest, err1.Error())
			return
		} else {
			handler.SendFail(w, r, http.StatusBadRequest, err2.Error())
			return
		}
	}
	usuario.Password = modelos.BeforeSave(usuario.Password)
}
