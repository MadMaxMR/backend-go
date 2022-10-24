package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	//"io/ioutil"
	"mime/multipart"
	"net/http"

	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gorilla/schema"

	"reflect"
)

func UploadImages(w http.ResponseWriter, req *http.Request) {
	pregunta := modelos.PreguntaExamens{}

	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	fmt.Println("Form del request: ", req.Form.Get("data"))
	jsonDecoder := json.NewDecoder(req.Body)
	fmt.Println("Body del request: ", req.Body)
	err = jsonDecoder.Decode(&pregunta)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "error en body"+err.Error())
		return
	}
	//err = json.Unmarshal(req.Form.Get("data"), &pregunta)
	decoder := schema.NewDecoder()

	err = decoder.Decode(&pregunta, req.Form)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, ("error decoder " + err.Error()))
	}
	fmt.Println("La estruct pregunta es: ", pregunta)
	fmt.Println("*****************************************")
	fmt.Println("la pregunta impresa es :", pregunta.Enunciado1)
	fmt.Println("tipo de dato pregunta : ", reflect.TypeOf(pregunta))

	file, _, err := req.FormFile("image")
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	UpImage(file, "10")
}

func UpImage(image multipart.File, id string) {
	fmt.Println("UpImages tipo: ", reflect.TypeOf(image))
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
