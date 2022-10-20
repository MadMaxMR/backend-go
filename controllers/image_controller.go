package controllers

import (
	"context"
	"fmt"
	"mime/multipart"
	"net/http"

	"github.com/MadMaxMR/backend-go/handler"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"

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
}

func UpImage64(w http.ResponseWriter, req *http.Request) {
	handler.SendFail(w, req, http.StatusOK, "dentro de UpImage64")
}

func UpImages(image multipart.File, id string) (string, error) {
	fmt.Println("UpImages tipo: ", reflect.TypeOf(image))

	cdl, err := cloudinary.NewFromURL("cloudinary://919663283643663:r7-EgFidG0Eu1sFM26ZU1sASIAU@umachayfiles")
	if err != nil {
		return "", err
	}

	var filename string = "user-" + id
	var ctx = context.Background()

	uploadResult, err := cdl.Upload.Upload(ctx, image, uploader.UploadParams{
		PublicID:       filename,
		Folder:         "user",
		AllowedFormats: []string{"jpg", "png", "jpeg", "jfif"},
	})
	if uploadResult.AssetID == "" || err != nil {
		if err != nil {

			return "", err
		} else {
			return "", err
		}
	}

	return uploadResult.SecureURL, nil
}
