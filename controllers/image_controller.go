package controllers

import (
	"fmt"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"mime/multipart"
	"net/http"

	"reflect"
)

func UploadImages(w http.ResponseWriter, req *http.Request) {
	file, _, err := req.FormFile("image")
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("tipo de archivo : ", reflect.TypeOf(file))
	fmt.Println("/********************************************/*")
	UpImages(file)
}

func UpImage64(w http.ResponseWriter, req *http.Request) {
	handler.SendFail(w, req, http.StatusOK, "dentro de UpImage64")
}

func UpImages(image multipart.File) string,  error {
	fmt.Println("UpImages tipo: ", reflect.TypeOf(image))

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
}
