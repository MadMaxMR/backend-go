package controllers

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"

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
	usuario := modelos.Usuarios{}
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
	usuario.Password = modelos.BeforeSave(usuario.Password)
	modelo, err := database.Create(&usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println(usuario.ID)
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func DeleteUsuario(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
	id := mux.Vars(req)["id"]
	message, err := database.Delete(&usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, message)
}

func UpdateUsuario(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
	id := mux.Vars(req)["id"]
	err := auth.ValidateBody(req, &usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	if usuario.Password != "" {
		usuario.Password = modelos.BeforeSave(usuario.Password)
	}
	_, err = database.Update(&usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, "Actualización correcta")
}

func VerPerfil(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
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
	user, err := database.Get(&usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, user)
}

func Login(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
	err := auth.ValidateBody(req, &usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}
	err = auth.ValidateLogin(&usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	jwtKey, valid_id, err := SignIn(usuario.Email, usuario.Password)

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

func SignIn(email string, password string) (string, uint, error) {
	var err error
	usuario := modelos.Usuarios{}
	db := database.GetConnection()
	defer db.Close()

	err = db.Where("email = ?", email).Find(&usuario).Error

	if err != nil {
		err = errors.New("email incorrecto")
		return "", 0, err
	}
	err = modelos.VerifyPassword(usuario.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		err = errors.New("contraseña incorrecta")
		return "", 0, err
	}
	Last_Login := time.Now()
	db.Model(&usuario).Update("last_login", Last_Login)

	jwtKey, err := auth.CreateToken(usuario.ID)
	return jwtKey, usuario.ID, err
}

func UpdateAvatar1(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}

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

		usuario.Image = string(id) + "." + extension
		db.Save(&usuario)
		handler.SendSuccess(w, req, http.StatusOK, usuario)
	} else {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al encontrar usuario")
	}
}

func GetAvatar1(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
	id := mux.Vars(req)["id"]
	_, err := database.Get(&usuario, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	OpenFile, err := os.Open("media/usuarios/" + usuario.Image)
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

func SaveAvatar(w http.ResponseWriter, req *http.Request) {
	usuario := modelos.Usuarios{}
	img := modelos.ImageUpdate{}

	file, _, err := req.FormFile("image")
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al leer el archivo - "+err.Error())
		return
	}
	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error en el Token !-"+err.Error())
		return
	}

	cdl, err := cloudinary.NewFromURL("cloudinary://919663283643663:r7-EgFidG0Eu1sFM26ZU1sASIAU@umachayfiles")
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al acceder a cloudinary - "+err.Error())
		return
	}

	var filename string = "user-" + tk.Id_Usuario
	var ctx = context.Background()

	uploadResult, err := cdl.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID:       filename,
		Folder:         "user",
		AllowedFormats: []string{"jpg", "png", "jpeg", "jfif"},
	})
	if uploadResult.AssetID == "" || err != nil {
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al subir la imagen - "+err.Error())
			return
		} else {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al subir la imagen")
			return
		}
	}

	usuario.Image = uploadResult.SecureURL
	_, err = database.Update(&usuario, tk.Id_Usuario)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	img.Image = usuario.Image
	// handler.SendSuccessMessage(w, req, http.StatusOK, usuario.Image)
	handler.SendSuccess(w, req, http.StatusOK, img)
}

func UpdateAvatar(w http.ResponseWriter, req *http.Request) {
	//usuario := modelos.Usuarios{}
	id := mux.Vars(req)["id"]

	cld, err := cloudinary.NewFromURL("cloudinary://919663283643663:r7-EgFidG0Eu1sFM26ZU1sASIAU@umachayfiles")
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al acceder a cloudinary - "+err.Error())
		return
	}
	//file, _, err := req.FormFile("image")
	/*if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al leer el archivo - "+err.Error())
		return
	}*/
	var filename string = "user-" + id
	var ctx = context.Background()

	_, err = cld.Upload.Destroy(ctx, uploader.DestroyParams{
		PublicID: "user/" + filename,
	})
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al borrar la imagen - "+err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, "Borrado de imagen correcta")
}
