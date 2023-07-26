package controllers

import (
	"fmt"
	"math"
	"time"

	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

// SaveExamens controller para crear y guardar un nuevo examen con preguntas y respuestas
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
	examen.LimitePreguntas = 50
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
	pageSizes := req.URL.Query().Get("pageSize")

	if page == "" {
		page = "1"
	}
	if pageSizes == "" {
		pageSizes = "10"
	}

	pageInt, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(pageSizes)

	db := database.GetConnection()
	defer db.Close()

	type Result struct {
		Page     string
		Prev     bool
		Next     bool
		Total    int
		Examenes []modelos.Examens
	}
	result2 := Result{}

	result2.Page = page
	result2.Next = true
	if pageInt == 1 {
		result2.Prev = false
	}
	if pageInt > 1 {
		result2.Prev = true
	}

	_, _ = database.GetAll(&examen, "")
	if len(examen)%pageSize == 0 {
		result2.Total = len(examen) / pageSize
	} else {
		result2.Total = (len(examen) / pageSize) + 1
	}

	if pageInt == result2.Total {
		result2.Next = false
	}

	result := db.Model(&examen).Select("*").Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("id DESC").Find(&examen)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No se encontró exámenes")
		return
	}

	result2.Examenes = examen
	handler.SendSuccess(w, req, http.StatusOK, result2)
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

	err := auth.ValidateBody(req, &examen)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	_, err = database.Update(&examen, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, "Examen actualizado correctamente")
}

func DeleteExamen(w http.ResponseWriter, req *http.Request) {
	examen := modelos.Examens{}
	examenPreguntas := []modelos.ExamenPreguntas{}

	id := mux.Vars(req)["id"]

	db := database.GetConnection()

	defer db.Close()

	db.Raw("SELECT * FROM examen_preguntas where examens_id = ?", id).Scan(&examenPreguntas)
	if len(examenPreguntas) != 0 {
		handler.SendFail(w, req, http.StatusNotAcceptable, "Invalid-El examen tiene preguntas")
		return
	}
	message, err := database.Delete(&examen, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, message)
}

func GetPreguntasExamenByArea(w http.ResponseWriter, req *http.Request) {
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

	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	historial := modelos.MisExamenes{}
	points := modelos.Result{Resultado: make(map[string]string), Solucion: make(map[string]uint)}
	result := map[string]interface{}{}
	var solution, answers string
	idExamen := req.URL.Query().Get("examen")
	if idExamen == "" {
		idExamen = "2"
	}
	idExamenInt, _ := strconv.Atoi(idExamen)

	db := database.GetConnection()
	defer db.Close()
	correct, incorrect, note := 0, 0, 0.0

	err = auth.ValidateBody(req, &result)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
	}

	examen := modelos.Examens{}
	db.Model(&examen).Where("id = ? ", idExamenInt).Find(&examen)

	for i := 1; i < (len(result)/2 + 1); i++ {
		respuesta := modelos.RespuestaExs{}
		ponderado := modelos.Ponderacion{}
		pregunta := modelos.PreguntaExamens{}

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
		} else if result["id_respuesta"+val] == float64(0) {
			points.Resultado["pregunta"+val] = "No contestada"
			points.Solucion["pregunta"+val] = respuesta.ID
			incorrect++
		}
		// respuesta.ID =

		historial.AreasId = examen.AreasId
		historial.UniversidadsId = examen.Id_Uni

	}
	points.Correct = correct
	points.Incorrect = incorrect
	points.Nota = math.Round(((note*20)/float64(examen.LimitePreguntas))*100) / 100

	historial.ExamensId = uint(idExamenInt)
	historial.Fecha_Examen = time.Now()
	historial.Nota = points.Nota
	if points.Nota < 10.5 {
		historial.Condicion = "Desaprobado"
	}
	historial.Condicion = "Aprobado"

	iduser, _ := strconv.Atoi(tk.Id_Usuario)
	historial.UsuarioId = uint(iduser)

	db.Create(&historial)

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

// GetExamensPregByArea retorna todos los examenes de un area con sus preguntas y alternativas
func GetExamensPregByArea(w http.ResponseWriter, req *http.Request) {
	preguntas := []modelos.PreguntaExamens{}
	examen := modelos.Examens{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	resultQ := db.Model(&examen).Where("id = ?", id).Find(&examen)
	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No existe el examen - "+id)
		return
	}
	if examen.LimitePreguntas != examen.CantidadPreguntas {
		handler.SendFail(w, req, http.StatusBadRequest, "El examen no tiene preguntas completas - "+fmt.Sprint(examen.LimitePreguntas)+"/"+fmt.Sprint(examen.CantidadPreguntas))
		return
	}
	resultQ = db.Preload("RespuestaExs").
		Where("ex.examens_id = ?", examen.ID).
		Select("pregunta_examens.id,ex.examens_id,pregunta_examens.enunciado1,pregunta_examens.grafico," +
			"pregunta_examens.enunciado2,pregunta_examens.enunciado3,row_number() OVER () AS num_question," +
			"pregunta_examens.cursos_id,pregunta_examens.temas_id,pregunta_examens.nivel").
		Joins("INNER JOIN examen_preguntas ex on ex.pregunta_examens_id = pregunta_examens.id").Order("pregunta_examens.id DESC").
		Find(&preguntas)
	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se agregó preguntas al examen")
		return
	}
	examen.PreguntaExamens = preguntas

	handler.SendSuccess(w, req, http.StatusOK, examen)
}

/*
SELECT pre.id,ex.examens_id,pre.enunciado1,pre.grafico,pre.enunciado2,pre.enunciado3,ex.num_question,pre.cursos_id,
pre.temas_id,pre.nivel
FROM examen_preguntas ex
INNER JOIN pregunta_examens  pre on ex.pregunta_examens_id = pre.id
WHERE ex.examens_id = 1
*/
func GetModalidad(w http.ResponseWriter, req *http.Request) {
	type Modalidades struct {
		Name string
		Code string
	}
	modalidad := []Modalidades{}

	db := database.GetConnection()
	defer db.Close()

	db.Raw("SELECT name,code FROM modalidades").Scan(&modalidad)
	if len(modalidad) == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No hay modalidades registradas")
		return
	}

	handler.SendSuccess(w, req, http.StatusOK, modalidad)
}

func GetExamensbyAnio(w http.ResponseWriter, req *http.Request) {
	type Año struct {
		Anio string
	}
	años := []Año{}
	id := mux.Vars(req)["id"]
	examenes := make(map[string]interface{})

	db := database.GetConnection()
	defer db.Close()

	db.Raw("SELECT anio FROM examens WHERE areas_id = $1 GROUP BY anio ", id).Scan(&años)
	if len(años) == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No se encontró examenes para el area seleccionada")
		return
	}

	for _, v := range años {
		var Examens []modelos.Examens
		db.Where("anio= $1 and areas_id = $2", v.Anio, id).Find(&Examens)
		examenes[v.Anio] = Examens
	}

	handler.SendSuccess(w, req, http.StatusOK, examenes)
}
