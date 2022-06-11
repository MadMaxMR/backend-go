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
	"strconv"

	//"os"
	//"strings"
	//"time"

	"github.com/gorilla/mux"
	//"golang.org/x/crypto/bcrypt"
)

func GetStudent(w http.ResponseWriter, req *http.Request) {

	student := modelos.Estudiantes{}
	id := mux.Vars(req)["id"]
	db := database.GetConnection()
	defer db.Close()
	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	if tk.Id_Usuario != id {
		handler.SendFail(w, req, http.StatusBadRequest, "Unauthorized")
		return
	}
	err = db.Where("usuarios_id = ?", id).First(&student).Error
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontro el estudiante - "+err.Error())
		return
	}

	err = db.Model(&student).Related(&student.Usuarios).Find(&student).Error
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, student)
}

func GetAllStudent(w http.ResponseWriter, req *http.Request) {
	students := []modelos.Estudiante{}
	db := database.GetConnection()
	defer db.Close()
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&students, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func SaveStudent(w http.ResponseWriter, req *http.Request) {
	student := modelos.Estudiante{}
	usuario := modelos.Usuarios{}
	db := database.GetConnection()
	err1 := auth.ValidateBody2(req, &usuario, &student)
	if err1 != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err1.Error())
		return
	}
	err1, err2 := auth.ValidateUsuario(&usuario), auth.ValidateStudent(&student)
	if err1 != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err1.Error())
		return
	}
	if err2 != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err2.Error())
		return
	}
	result := db.Where("email = ?", usuario.Email).First(&usuario)
	if result.RowsAffected > 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "El email ya existe")
		return
	}
	result = db.Where("dni = ?", usuario.Dni).First(&usuario)
	if result.RowsAffected > 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "El dni ya existe")
		return
	}

	usuario.Password = modelos.BeforeSave(usuario.Password)
	modelo, err := database.Create(&usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	valu := modelo.(*modelos.Usuarios)
	student.UsuariosId = valu.ID
	estudiante, err := database.Create(&student)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		database.Delete(&usuario, strconv.Itoa(int(valu.ID)))
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, estudiante)
}

func UpdateStudent(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
	estudiante := modelos.Estudiante{}
	db := database.GetConnection()
	defer db.Close()
	id := mux.Vars(req)["id"]
	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	if tk.Id_Usuario != id {
		handler.SendFail(w, req, http.StatusBadRequest, "Unauthorized")
		return
	}
	err = auth.ValidateBody2(req, &usuario, &estudiante)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	if usuario.Password != "" {
		usuario.Password = modelos.BeforeSave(usuario.Password)
	}
	result := db.Where("email = ? and id != ?", usuario.Email, id).First(&usuario)
	if result.RowsAffected > 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "El email ya existe")
		return
	}
	result = db.Where("dni = ? and id != ?", usuario.Dni, id).First(&usuario)
	if result.RowsAffected > 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "El dni ya existe")
		return
	}
	_, err = database.Update(&usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = db.Model(&estudiante).Where("usuarios_id = ?", id).Update(&estudiante).Error
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, "Actualización correcta")
}
