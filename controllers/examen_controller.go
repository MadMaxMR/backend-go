package controllers

import (
	//"fmt"

	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

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
	}
	modelo, err := database.Create(&pregunta)
	if err != nil {
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetPoints(w http.ResponseWriter, req *http.Request) {
	points := modelos.Result{Solucion: make(map[string]string)}
	result := map[string]interface{}{}
	//id := mux.Vars(req)["id"]
	db := database.GetConnection()
	correct := 0
	incorrect := 0
	err := auth.ValidateBody(req, &result)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}
	for i := 1; i < 101; i++ {
		respuesta := modelos.RespuestaExs{}
		val := strconv.Itoa(i)
		db.Model(&respuesta).Where("pregunta_examens_id = ? and valor = 'true'", result["id_pregunta"+val+""]).Find(&respuesta)

		if result["respuesta"+val] != float64(0) {
			if result["respuesta"+val] == float64(respuesta.ID) {
				points.Solucion["pregunta"+val] = "Correcto"
				correct++
			} else {
				points.Solucion["pregunta"+val] = "Incorrecto"
				incorrect++
			}
		}
		if result["respuesta"+val] == float64(0) {
			points.Solucion["pregunta"+val] = "Incorrecto"
			incorrect++
		}
		respuesta.ID = 0
	}
	points.Correct = correct
	points.Incorrect = incorrect
	points.Nota = float64(correct) / float64(correct+incorrect)
	handler.SendSuccess(w, req, http.StatusOK, points)
}
