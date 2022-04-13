package controllers

import (
	"fmt"
	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"net/http"
)

func ResetPassword(w http.ResponseWriter, req *http.Request) {

	tk := req.URL.Query().Get("tk")
	ml := req.URL.Query().Get("ml")

	email, err := auth.ValidateTokenReset(tk)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("ml: " + ml)
	fmt.Println("email: " + email)
	if ml != email {
		handler.SendFail(w, req, http.StatusInternalServerError, "El email no coincide con el token")
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, "El email conincide con el token")
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
	handler.SendSuccessMessage(w, req, http.StatusOK, "Se ha enviado un correo con el link para resetear la contraseña")
}

//func CreateTokenReset()
