package controllers

import (
	"context"
	"errors"
	"net/http"

	//"io/ioutil"
	"mime/multipart"

	"github.com/MadMaxMR/backend-go/handler"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
)

func UploadImages(w http.ResponseWriter, req *http.Request) {
	handler.SendSuccessMessage(w, req, http.StatusOK, "dentro de controller upload Image")
}

func UpImages(image multipart.File, id string, folder string) (string, error) {

	cdl, err := cloudinary.NewFromURL("cloudinary://919663283643663:r7-EgFidG0Eu1sFM26ZU1sASIAU@umachayfiles")
	if err != nil {
		return "", err
	}

	var filename string = folder + "-" + id
	var ctx = context.Background()

	uploadResult, err := cdl.Upload.Upload(ctx, image, uploader.UploadParams{
		PublicID:       filename,
		Folder:         folder,
		AllowedFormats: []string{"jpg", "png", "jpeg", "jfif"},
	})
	if uploadResult.AssetID == "" || err != nil {
		if err != nil {
			return "", errors.New("Error al subir imagen " + err.Error())
		} else {
			return "", errors.New("error al subir imagen")
		}
	}

	return uploadResult.SecureURL, nil
}
