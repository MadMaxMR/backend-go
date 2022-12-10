package controllers

import (
	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/gorilla/mux"
)

func SavePreguntasRespuestas(w http.ResponseWriter, req *http.Request) {
	pregunta := modelos.PreguntaExamens{}
	respuestas := []modelos.RespuestaExs{
		{Valor: true, Respuesta: ""}, {Valor: false, Respuesta: ""}, {Valor: false, Respuesta: ""},
		{Valor: false, Respuesta: ""}, {Valor: false, Respuesta: ""}}

	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "error parse"+err.Error())
		return
	}

	pregunta.ExamensId = 2
	pregunta.Enunciado1 = req.Form.Get("enunciado1")
	pregunta.Enunciado2 = req.Form.Get("enunciado2")
	pregunta.Enunciado3 = req.Form.Get("enunciado3")
	nq, _ := strconv.Atoi(req.Form.Get("NumQuestion"))
	pregunta.NumQuestion = uint(nq)
	cI, _ := strconv.Atoi(req.Form.Get("CursosId"))
	pregunta.CursosId = uint(cI)

	for i := 0; i < 5; i++ {
		index := strconv.Itoa(i + 1)
		bol, _ := strconv.ParseBool(req.Form.Get("Valor" + index))
		respuestas[i].Valor = bol
		respuestas[i].Respuesta = req.Form.Get("Respuesta" + index)
	}
	pregunta.RespuestaExs = respuestas

	_, err = database.Create(&pregunta)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "error al crear pregunta "+err.Error())
		return
	}

	preguntaID := strconv.FormatUint(uint64(pregunta.ID), 10)
	file, _, _ := req.FormFile("grafico")
	if file != nil {
		urlPreg, err := UpImages(file, preguntaID, "Pregunta")
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "error al subir imagen a servidor "+err.Error())
			return
		}
		pregunta.Grafico = urlPreg
	}

	for i := 0; i < 5; i++ {
		index := strconv.Itoa(i + 1)
		idRes := strconv.FormatUint(uint64(pregunta.RespuestaExs[i].ID), 10)
		fileR, _, _ := req.FormFile("image" + index)
		if fileR == nil {
			continue
		}
		urlRes, err := UpImages(fileR, idRes, "Respuesta")
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, err.Error())
			return
		}
		respuestas[i].ID = pregunta.RespuestaExs[i].ID
		respuestas[i].ImgLink = urlRes
		_, err = database.Update(&respuestas[i], idRes)
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, err.Error())
			return
		}
	}
	_, err = database.Update(&pregunta, preguntaID)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusCreated, pregunta)
}

func GetAllPreguntas(w http.ResponseWriter, req *http.Request) {
	preguntas := modelos.PreguntaExamens{}
	page := req.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageInt, _ := strconv.Atoi(page)
	type Result struct {
		Id           uint
		Enunciado1   string
		Nombre_curso string
		Nombre_tema  string
	}
	result := []Result{}

	db := database.GetConnection()
	defer db.Close()

	resultQ := db.Model(&preguntas).Select("DISTINCT pregunta_examens.id,pregunta_examens.enunciado1, cursos.nombre_curso,temas.nombre_tema").
		Joins("LEFT JOIN temas on pregunta_examens.temas_id = temas.id").
		Joins("LEFT JOIN cursos on pregunta_examens.cursos_id = cursos.id").
		Limit(25).Offset((pageInt - 1) * 25).Order("id ASC").Scan(&result)
	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró preguntas")
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, result)
}

func GetPreguntasCursoTema(w http.ResponseWriter, req *http.Request) {
	preguntas := modelos.PreguntaExamens{}
	id := mux.Vars(req)["id"]
	page := req.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageInt, _ := strconv.Atoi(page)
	type Result struct {
		Id           uint   `json:"id"`
		Enunciado1   string `json:"enunciado"`
		Nombre_curso string `json:"curso"`
		Nombre_tema  string `json:"tema"`
	}
	result := []Result{}

	db := database.GetConnection()
	defer db.Close()

	resultQ := db.Model(&preguntas).Select("DISTINCT pregunta_examens.id,pregunta_examens.enunciado1, cursos.nombre_curso,temas.nombre_tema").
		Joins("LEFT JOIN temas on pregunta_examens.temas_id = temas.id").
		Joins("LEFT JOIN cursos on pregunta_examens.cursos_id = cursos.id").
		Where("temas.id  = ?", id).
		Limit(25).Offset((pageInt - 1) * 25).Order("id ASC").Scan(&result)
	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró preguntas")
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, result)
}

func GetPregunta(w http.ResponseWriter, req *http.Request) {
	preguntas := modelos.PreguntaExamens{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	result := db.Model(&preguntas).Where("id = ?", id).Preload("RespuestaExs").Find(&preguntas)

	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró pregunta con el id: "+id)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, preguntas)
}
