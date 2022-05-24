package controllers

import (
	//"fmt"
	"net/http"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
)

func ResetPassword(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
	db := database.GetConnection()
	tk := req.URL.Query().Get("tk")
	ml := req.URL.Query().Get("ml")

	email, err := auth.ValidateTokenReset(tk)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	if ml != email {
		handler.SendFail(w, req, http.StatusInternalServerError, "El email no coincide con el token")
		return
	}

	err = auth.ValidateBody(req, &usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateReset(&usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	usuario.Password = modelos.BeforeSave(usuario.Password)
	err = db.Model(&usuario).Where("email = ?", ml).Update("password", usuario.Password).Error
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, "Contraseña cambiada correctamente")
}

func RecoveryPassword(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
	err := auth.ValidateBody(req, &usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateRecovery(&usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	ml := usuario.Email
	db := database.GetConnection()
	err = db.Where("email = ?", ml).First(&usuario).Error
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "No existe el usuario")
		return
	}
	tk, err := auth.CreateTokenReset(ml, usuario.ID)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al crear el token")
		return
	}
	SendMail(ml, tk)
	handler.SendSuccessMessage(w, req, http.StatusOK, "Se ha enviado un correo con el link para restablecer la contraseña")
}
