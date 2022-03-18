package controllers

import (
	"API/auth"
	"API/database"
	"API/handler"
	"API/modelos"
	"API/models"
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUsuarios(w http.ResponseWriter, req *http.Request) {
	usuarios := []modelos.Usuarios{}
	page := req.URL.Query().Get("page")
	user, err := database.GetAll(&usuarios, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, user)
}

func GetUsuario(w http.ResponseWriter, req *http.Request) {
	usuario := models.Usuarios{}
	id := mux.Vars(req)["id"]

	user, err := database.Get(&usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, user)
}

func SaveUsuario(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
	err := auth.ValidateBody(req, &usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateUsuario(&usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	usuario.Password = models.BeforeSave(usuario.Password)
	modelo, err := database.Create(&usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func DeleteUsuario(w http.ResponseWriter, req *http.Request) {
	usuario := models.Usuarios{}
	id := mux.Vars(req)["id"]
	message, err := database.Delete(&usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, message)
}

func UpdateUsuario(w http.ResponseWriter, req *http.Request) {
	usuario := models.Usuarios{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Update(req, &usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func VerPerfil(w http.ResponseWriter, req *http.Request) {
	usuario := models.Usuarios{}
	id := mux.Vars(req)["id"]
	user, err := database.Get(&usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, user)
}

func Login(w http.ResponseWriter, req *http.Request) {
	usuario := models.Usuarios{}
	err := auth.ValidateBody(req, &usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}
	err = auth.ValidateLogin(&usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	jwtKey, err, valid_id := SignIn(usuario.Email, usuario.Password)

	if err != nil {
		//log.Fatal(err)
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
		return
	}
	/*Llenamos el data con el ID del usuario y el token generado*/
	data := auth.Token{
		Id_Usuario: valid_id,
		Token:      jwtKey,
	}
	handler.SendSuccess(w, req, http.StatusOK, data)
}

func SignIn(email string, password string) (string, error, uint) {
	var err error
	usuario := models.Usuarios{}
	db := database.GetConnection()
	defer db.Close()

	err = db.Where("email = ?", email).Find(&usuario).Error

	if err != nil {
		err = errors.New("email incorrecto")
		return "", err, 0
	}
	err = models.VerifyPassword(usuario.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err = errors.New("contraseña incorrecta")
		return "", err, 0
	}
	Last_Login := time.Now()
	db.Model(&usuario).Update("last_login", Last_Login)

	jwtKey, err := auth.CreateToken(usuario.ID)
	return jwtKey, err, usuario.ID
}

func UpdateAvatar(w http.ResponseWriter, req *http.Request) {
	usuario := models.Usuarios{}

	db := database.GetConnection()
	defer db.Close()

	id := mux.Vars(req)["id"]
	db.Find(&usuario, id)

	if usuario.ID > 0 {
		file, hand, err := req.FormFile("image")
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al leer el archivo - "+err.Error())
			return
		}
		var extension = strings.Split(hand.Filename, ".")[1]
		var archivo string = "media/usuarios/" + id + "." + extension

		f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al crear el archivo - "+err.Error())
			return
		}
		_, err = io.Copy(f, file)
		if err != nil {
			handler.SendFail(w, req, http.StatusInternalServerError, "Error al copiar elarchivo - "+err.Error())
			return
		}
		defer f.Close()

		usuario.Imagen = string(id) + "." + extension
		db.Save(&usuario)
		handler.SendSuccess(w, req, http.StatusOK, usuario)
	} else {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al encontrar usuario")
	}
}

func GetAvatar(w http.ResponseWriter, req *http.Request) {
	usuario := models.Usuarios{}
	id := mux.Vars(req)["id"]
	_, err := database.Get(&usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	OpenFile, err := os.Open("media/usuarios/" + usuario.Imagen)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Imagen no encontrada -"+err.Error())
		return
	}
	/*envio de la imagen*/
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al copiar la imagen - "+err.Error())
		return
	}
	log.Print("'", req.Method, " - ", req.URL.Path, " - ", req.Proto, "' - ", http.StatusOK, " - ", req.RemoteAddr)
}
