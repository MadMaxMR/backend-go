package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"sync"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/gorilla/mux"
)

//SavePreguntasRespuestas guarda una pregunta con sus respuestas e imagenes, recibe el body en tipo Form-Data
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
	cI, _ := strconv.Atoi(req.Form.Get("CursosId"))
	pregunta.CursosId = uint(cI)
	ti, _ := strconv.Atoi(req.Form.Get("TemasId"))
	pregunta.TemasId = uint(ti)
	pregunta.Nivel = req.Form.Get("Nivel")
	pregunta.Tipo = req.Form.Get("Tipo")

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
	_, err = database.Update(&pregunta, preguntaID)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		index := strconv.Itoa(i + 1)
		idRes := strconv.FormatUint(uint64(pregunta.RespuestaExs[i].ID), 10)
		bol, _ := strconv.ParseBool(req.Form.Get("Valor" + index))
		fileR, _, _ := req.FormFile("image" + index)
		if fileR == nil {
			continue
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			urlRes, err := UpImages(fileR, idRes, "Respuesta")
			if err != nil {
				handler.SendFail(w, req, http.StatusBadRequest, err.Error())
				return
			}
			respuestas[i].ID = pregunta.RespuestaExs[i].ID
			respuestas[i].Valor = bol
			respuestas[i].Respuesta = req.Form.Get("Respuesta" + index)
			respuestas[i].ImgLink = urlRes
			_, err = database.Update(&respuestas[i], idRes)
			if err != nil {
				handler.SendFail(w, req, http.StatusBadRequest, err.Error())
				return
			}
		}(i)
	}

	wg.Wait()
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
		Nivel        string
		Tipo         string
	}
	result := []Result{}

	db := database.GetConnection()
	defer db.Close()

	resultQ := db.Model(&preguntas).Select("DISTINCT pregunta_examens.id,pregunta_examens.enunciado1,pregunta_examens.nivel,pregunta_examens.tipo, cursos.nombre_curso,temas.nombre_tema").
		Joins("LEFT JOIN temas on pregunta_examens.temas_id = temas.id").
		Joins("LEFT JOIN cursos on pregunta_examens.cursos_id = cursos.id").
		Limit(25).Offset((pageInt - 1) * 25).Order("id DESC").Scan(&result)
	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró preguntas")
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, result)
}

func GetPreguntasForExamen(w http.ResponseWriter, req *http.Request) {
	preguntas := modelos.PreguntaExamens{}
	idExamen := mux.Vars(req)["idExamen"]
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
		Nivel        string
	}
	result := []Result{}

	db := database.GetConnection()
	defer db.Close()

	resultQ := db.Model(&preguntas).Select("DISTINCT pregunta_examens.id,pregunta_examens.enunciado1,pregunta_examens.nivel, cursos.nombre_curso,temas.nombre_tema").
		Joins("INNER JOIN temas on pregunta_examens.temas_id = temas.id").
		Joins("INNER JOIN cursos on pregunta_examens.cursos_id = cursos.id").
		Where("pregunta_examens.id <> ALL (select pregunta_examens_id from examen_preguntas where examens_id =?)", idExamen).
		Limit(25).Offset((pageInt - 1) * 25).Order("id DESC").Scan(&result)
	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró preguntas")
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, result)
}

func GetPreguntasOfExamen(w http.ResponseWriter, req *http.Request) {
	preguntas := modelos.PreguntaExamens{}
	idExamen := mux.Vars(req)["idExamen"]
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
		Nivel        string
	}
	result := []Result{}

	db := database.GetConnection()
	defer db.Close()

	resultQ := db.Model(&preguntas).Select("DISTINCT examen_preguntas.id,pregunta_examens.enunciado1,pregunta_examens.nivel, cursos.nombre_curso,temas.nombre_tema").
		Joins("LEFT JOIN temas on pregunta_examens.temas_id = temas.id").
		Joins("LEFT JOIN cursos on pregunta_examens.cursos_id = cursos.id").
		Joins("INNER JOIN examen_preguntas on pregunta_examens.id = examen_preguntas.pregunta_examens_id").
		Where("examen_preguntas.examens_id = ?", idExamen).
		Limit(25).Offset((pageInt - 1) * 25).Order("id DESC").Scan(&result)
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
		Id           uint
		Enunciado1   string
		Nombre_curso string
		Nombre_tema  string
		Nivel        string
	}
	result := []Result{}

	db := database.GetConnection()
	defer db.Close()

	resultQ := db.Model(&preguntas).Select("DISTINCT pregunta_examens.id,pregunta_examens.enunciado1,pregunta_examens.nivel, cursos.nombre_curso,temas.nombre_tema").
		Joins("LEFT JOIN temas on pregunta_examens.temas_id = temas.id").
		Joins("LEFT JOIN cursos on pregunta_examens.cursos_id = cursos.id").
		Where("temas.id  = ?", id).
		Limit(25).Offset((pageInt - 1) * 25).Order("id DESC").Scan(&result)
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

func UpdatePreguntaRespuestas(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	idPregunta, _ := strconv.Atoi(id)
	pregunta := modelos.PreguntaExamens{}
	respuestas := []modelos.RespuestaExs{}

	err := req.ParseMultipartForm(32 << 20)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "error en parse "+err.Error())
		return
	}

	db := database.GetConnection()
	defer db.Close()

	db.Model(&respuestas).Where("pregunta_examens_id  = ?", id).Find(&respuestas)

	pregunta.ID = uint(idPregunta)
	pregunta.ExamensId = 2
	pregunta.Enunciado1 = req.Form.Get("enunciado1")
	pregunta.Enunciado2 = req.Form.Get("enunciado2")
	pregunta.Enunciado3 = req.Form.Get("enunciado3")
	nq, _ := strconv.Atoi(req.Form.Get("NumQuestion"))
	pregunta.NumQuestion = uint(nq)
	cI, _ := strconv.Atoi(req.Form.Get("CursosId"))
	pregunta.CursosId = uint(cI)
	ti, _ := strconv.Atoi(req.Form.Get("TemasId"))
	pregunta.TemasId = uint(ti)
	pregunta.Nivel = req.Form.Get("Nivel")
	pregunta.Tipo = req.Form.Get("Tipo")

	pregunta.RespuestaExs = respuestas

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
		bol, _ := strconv.ParseBool(req.Form.Get("Valor" + index))
		respuestas[i].Valor = bol
		respuestas[i].Respuesta = req.Form.Get("Respuesta" + index)
		respuestas[i].PreguntaExamensId = uint(idPregunta)
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
	}
	pregunta.RespuestaExs = respuestas

	err = db.Save(&pregunta).Error
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusCreated, pregunta)
}

func DeletePreguntaRespuestas(w http.ResponseWriter, req *http.Request) {
	pregunta := modelos.PreguntaExamens{}
	respuestas := []modelos.RespuestaExs{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	result := modelos.ExamenPreguntas{}

	response := db.Table("examen_preguntas").Where("pregunta_examens_id = ?", id).Take(&result)
	if response.RowsAffected >= 1 {
		handler.SendSuccessMessage(w, req, http.StatusBadRequest, "No se puede eliminar una pregunta que ya se agregó a un examen")
		return
	}

	db.Model(&respuestas).Where("pregunta_examens_id  = ?", id).Find(&respuestas)

	err := db.Where("pregunta_examens_id  = ?", id).Delete(&respuestas).Error
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "error eliminando respuestas "+err.Error())
		return
	}

	_, err = database.Delete(&pregunta, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "error eliminando pregunta "+err.Error())
		return
	}

	if pregunta.Grafico != "" {
		_, err := DeleteImage(id, "Pregunta")
		if err != nil {
			fmt.Println("Error al borrar imagen de servidor" + err.Error())
		}
	}

	for i := 0; i < len(respuestas); i++ {
		if respuestas[i].ImgLink != "" {
			val := strconv.Itoa(int(respuestas[i].ID))
			_, err := DeleteImage(val, "Respuesta")
			if err != nil {
				fmt.Println("Error al borrar imagen de servidor" + err.Error())
			}
		}
	}

	handler.SendSuccessMessage(w, req, http.StatusOK, "Pregunta eliminada correctamente")
}

func SavePreguntasExamen(w http.ResponseWriter, req *http.Request) {
	preguntaEx := modelos.ExamenPreguntas{}
	data := map[string]interface{}{}

	db := database.GetConnection()
	defer db.Close()

	err := auth.ValidateBody(req, &data)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	idExamen := uint(data["id_examen"].(float64))
	preguntaEx.ExamensId = idExamen
	preguntas := data["preguntas"].([]interface{})

	result := db.Model(&preguntaEx).Where("examens_id = ?", idExamen).Find(&preguntaEx)

	if (len(preguntas) + int(result.RowsAffected)) > 50 {
		handler.SendFail(w, req, http.StatusBadRequest, "Las cantidad de preguntas superan el limite de preguntas del examen")
		return
	}

	var wg sync.WaitGroup
	wg.Add(len(preguntas))

	for _, pregunta := range preguntas {
		go func(pregunta interface{}) {
			defer wg.Done()
			preguntaEx := modelos.ExamenPreguntas{}
			preguntaEx.ExamensId = idExamen
			preguntaEx.PreguntaExamensId = uint(pregunta.(map[string]interface{})["id_pregunta"].(float64))
			err = db.Create(&preguntaEx).Error
			if err != nil {
				handler.SendFail(w, req, http.StatusBadRequest, "Error al guardar pregunta - "+err.Error())
				return
			}
		}(pregunta)
	}

	wg.Wait()
	db.Table("examens").Where("id = ?", idExamen).UpdateColumn("cantidad_preguntas", result.RowsAffected+int64(len(preguntas)))

	handler.SendSuccessMessage(w, req, http.StatusOK, "Preguntas agregadas exitosamente")
}

func ChangePreguntaExamen(w http.ResponseWriter, req *http.Request) {
	preguntaEx := modelos.ExamenPreguntas{}
	id := mux.Vars(req)["id"]

	err := auth.ValidateBody(req, &preguntaEx)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	_, err = database.Update(&preguntaEx, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	handler.SendSuccessMessage(w, req, http.StatusOK, "Pregunta actualizada correctamente")
}

func DeletePreguntaExamen(w http.ResponseWriter, req *http.Request) {
	preguntaEx := modelos.ExamenPreguntas{}
	id := mux.Vars(req)["idPregunta"]

	db := database.GetConnection()
	defer db.Close()

	_, err := database.Delete(&preguntaEx, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	examenPregunta := modelos.ExamenPreguntas{}
	result := db.Model(&examenPregunta).Where("examens_id = ?", preguntaEx.ExamensId).Find(&examenPregunta)
	if result.RowsAffected != 0 {
		db.Table("examens").Where("id = ?", preguntaEx.ExamensId).UpdateColumn("cantidad_preguntas", result.RowsAffected)
	}

	handler.SendSuccessMessage(w, req, http.StatusOK, "Pregunta eliminada correctamente")
}

//GetPreguntasforETA devuelve 10 preguntas para FAS TEST 7 de tipo ETA y 3 de tipo Admision
func GetPreguntasforETA(w http.ResponseWriter, req *http.Request) {
	preguntas := []modelos.PreguntaExamens{}
	id := mux.Vars(req)["id"]
	db := database.GetConnection()
	defer db.Close()

	err := db.Raw("select * from fn_preguntas_eta($1)", id).Scan(&preguntas).Error

	if err != nil {
		handler.SendFail(w, req, http.StatusInternalServerError, err.Error())
		return
	}

	for i, pregunta := range preguntas {
		db.Raw("Select * from respuesta_exs where pregunta_examens_id = $1", pregunta.ID).Scan(&preguntas[i].RespuestaExs)
	}
	handler.SendSuccess(w, req, http.StatusOK, preguntas)
}
