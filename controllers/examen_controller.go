package controllers

import (
	"fmt"
	"math"

	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

//SaveExamens controller para crear y guardar un nuevo examen con preguntas y respuestas
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

func GetExamenById(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	result := db.Model(&examen).Where("id = ?", id).Preload("PreguntaExamens", func(db *gorm.DB) *gorm.DB {
		return db.Order("pregunta_examens.id ASC")
	}).Preload("PreguntaExamens.RespuestaExs", func(db *gorm.DB) *gorm.DB {
		return db.Order("respuesta_exs.id ASC")
	}).Find(&examen)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró examenes con el id: "+id)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, examen)
}

func UpdateExamen(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	id := mux.Vars(req)["id"]

	_, err := database.Update(&examen, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, "Examen actualizado correctamente")
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

func GetPoints(w http.ResponseWriter, req *http.Request) {
	points := modelos.Result{Resultado: make(map[string]string), Solucion: make(map[string]uint)}
	result := map[string]interface{}{}
	var solution, answers string

	db := database.GetConnection()
	defer db.Close()
	correct, incorrect, note := 0, 0, 0.0

	err := auth.ValidateBody(req, &result)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}

	for i := 1; i < (len(result)/2 + 1); i++ {
		respuesta := modelos.RespuestaExs{}
		ponderado := modelos.Ponderacion{}
		pregunta := modelos.PreguntaExamens{}
		examen := modelos.Examens{}
		val := strconv.Itoa(i)
		rest := db.Model(&respuesta).Where("pregunta_examens_id = ? and valor = 'true'", result["id_pregunta"+val]).Find(&respuesta)
		if rest.RowsAffected == 0 {
			handler.SendFail(w, req, http.StatusBadRequest, "No hay alternativa correcta")
			return
		}
		solution += strconv.Itoa(int(respuesta.ID)) + "-"
		answers += fmt.Sprintf("%v", result["id_respuesta"+val]) + "-"
		if result["id_respuesta"+val] != float64(0) {
			if result["id_respuesta"+val] == float64(respuesta.ID) {
				db.Model(&pregunta).Where("id = ? ", result["id_pregunta"+val]).Find(&pregunta)
				db.Model(&examen).Where("id = ? ", pregunta.ExamensId).Find(&examen)
				db.Model(&ponderado).Where("cursos_id = ? and cod_area = ?", pregunta.CursosId, examen.AreasId).Find(&ponderado)
				points.Resultado["pregunta"+val] = "Correcto"
				points.Solucion["pregunta"+val] = respuesta.ID
				note = note + ponderado.Ponderacion
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
	points.Nota = math.Round(((note*20)/50)*100) / 100
	fmt.Println("answers: ", answers)
	fmt.Println("solution: ", solution)
	fmt.Println("Nota general", note)
	handler.SendSuccess(w, req, http.StatusOK, points)
}

// a := "10-15-119-5-8-10-55-"
// 	cadena := strings.Split(a, "-")
// 	for i := 0; i < len(cadena)-1; i++ {
// 		fmt.Print("\n valor ", i, ":", cadena[i])
// 	}
