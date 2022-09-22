package controllers

import (
	"fmt"

	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//SaveExamens controller para crear y guardar un nuevo examen
func SaveExamens(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	err := auth.ValidateBody(req, &examen)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateExamen(&examen)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	_, err = database.Create(&examen)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusCreated, examen)
}

func GetAllExamens(w http.ResponseWriter, req *http.Request) {
	examen := []modelos.Examens{}
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&examen, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetExamen(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Get(&examen, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func DeleteExamen(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	id := mux.Vars(req)["id"]
	message, err := database.Delete(&examen, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, message)
}

//GetExamensPregByArea retorna todos los examenes de un area con sus preguntas y alternativas
func GetExamensPregByArea(w http.ResponseWriter, req *http.Request) {
	examen := []modelos.Examens{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	result := db.Model(&examen).Where("areas_id = ?", id).Preload("PreguntaExamens", func(db *gorm.DB) *gorm.DB {
		return db.Order("pregunta_examens.id ASC")
	}).Preload("PreguntaExamens.RespuestaExs", func(db *gorm.DB) *gorm.DB {
		return db.Order("respuesta_exs.id ASC")
	}).Find(&examen)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró examenes para el area: "+id)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, examen)
}

func SavePreguntaResp(w http.ResponseWriter, req *http.Request) {
	pregunta := modelos.PreguntaExamens{}
	err := auth.ValidateBody(req, &pregunta)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	modelo, err := database.Create(&pregunta)
	if err != nil {
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetPoints(w http.ResponseWriter, req *http.Request) {
	points := modelos.Result{Resultado: make(map[string]string), Solucion: make(map[string]uint)}
	result := map[string]interface{}{}
	var solution, answers string

	db := database.GetConnection()
	defer db.Close()
	correct, incorrect := 0, 0

	err := auth.ValidateBody(req, &result)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}

	for i := 1; i < (len(result)/2 + 1); i++ {
		respuesta := modelos.RespuestaExs{}
		val := strconv.Itoa(i)
		db.Model(&respuesta).Where("pregunta_examens_id = ? and valor = 'true'", result["id_pregunta"+val]).Find(&respuesta)
		solution += strconv.Itoa(int(respuesta.ID)) + "-"
		answers += fmt.Sprintf("%v", result["id_respuesta"+val]) + "-"
		if result["id_respuesta"+val] != float64(0) {
			if result["id_respuesta"+val] == float64(respuesta.ID) {
				points.Resultado["pregunta"+val] = "Correcto"
				points.Solucion["pregunta"+val] = respuesta.ID
				correct++
			} else {
				points.Resultado["pregunta"+val] = "Incorrecto"
				points.Solucion["pregunta"+val] = respuesta.ID
				incorrect++
			}
		}
		if result["id_respuesta"+val] == float64(0) {
			points.Resultado["pregunta"+val] = "No contestada"
			points.Solucion["pregunta"+val] = respuesta.ID
			incorrect++
		}
		// respuesta.ID = 0
	}
	points.Correct = correct
	points.Incorrect = incorrect
	points.Nota = float64(correct) / float64(correct+incorrect)
	fmt.Println("answers: ", answers)
	fmt.Println("solution: ", solution)
	handler.SendSuccess(w, req, http.StatusOK, points)
}

// a := "10-15-119-5-8-10-55-"
// 	cadena := strings.Split(a, "-")
// 	for i := 0; i < len(cadena)-1; i++ {
// 		fmt.Print("\n valor ", i, ":", cadena[i])
// 	}
