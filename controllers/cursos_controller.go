package controllers

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"

	"github.com/gorilla/mux"
)

func GetAllCursos(w http.ResponseWriter, req *http.Request) {
	cursos := []modelos.Cursos{}

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()
	result := db.Preload("Temas").Find(&cursos)

	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No se encontraron registros")
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, cursos)
}

func GetCurso(w http.ResponseWriter, req *http.Request) {
	curso := modelos.Cursos{}
	id := mux.Vars(req)["id"]
	modelo, err := database.Get(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusNotFound, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func SaveCurso(w http.ResponseWriter, req *http.Request) {
	curso := modelos.Cursos{}
	err := auth.ValidateBody(req, &curso)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	err = auth.ValidateCurso(&curso)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	model, err := database.Create(&curso)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusCreated, model)
}

func DeleteCurso(w http.ResponseWriter, req *http.Request) {
	// curso := modelos.Cursos{}
	db := database.GetConnection()
	id := mux.Vars(req)["id"]

	dbc, _ := db.DB()
	defer dbc.Close()

	type Result struct {
		Response string
		Status   int
	}
	var res []Result
	db.Raw("CALL delete_cursos($1)", id).Scan(&res)

	if res[0].Status == 400 {
		handler.SendFail(w, req, http.StatusBadRequest, res[0].Response)
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, res[0].Response)
}

func UpdateCurso(w http.ResponseWriter, req *http.Request) {
	curso := modelos.Cursos{}
	id := mux.Vars(req)["id"]
	err := auth.ValidateBody(req, &curso)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	_, err = database.Update(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	idI, _ := strconv.Atoi(id)
	curso.ID = uint(idI)
	handler.SendSuccess(w, req, http.StatusOK, curso)
}

func UploadImage(w http.ResponseWriter, req *http.Request) {
	curso := modelos.Cursos{}

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	id := mux.Vars(req)["id"]

	db.Find(&curso, id)

	if curso.ID > 0 {
		file, hand, err := req.FormFile("image")
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al leer el archivo - "+err.Error())
			return
		}
		var extension = strings.Split(hand.Filename, ".")[1]
		var archivo string = "media/cursos/" + id + "." + extension

		f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al crear el archivo - "+err.Error())
			return
		}
		_, err = io.Copy(f, file)
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, "Error al copiar el archivo - "+err.Error())
			return
		}
		defer f.Close()

		curso.Image = string(id) + "." + extension
		db.Save(&curso)
		handler.SendSuccess(w, req, http.StatusOK, curso)
	} else {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontro el curso")
	}
}

func GetImage(w http.ResponseWriter, req *http.Request) {
	curso := modelos.Cursos{}
	id := mux.Vars(req)["id"]
	_, err := database.Get(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}

	OpenFile, err := os.Open("media/cursos/" + curso.Image)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encuentra la imagen -"+err.Error())
		return
	}
	_, err = io.Copy(w, OpenFile)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, "Error al copiar el archivo - "+err.Error())
		return
	}
}

func GetCursoByArea(w http.ResponseWriter, req *http.Request) {
	cursos := modelos.Cursos{}
	CursosArea := []modelos.CursosArea{}

	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	result := db.Model(&cursos).Select("DISTINCT cursos.id as id,cursos_universidades.cod_area, cursos.nombre_curso").
		Joins("inner join cursos_universidades ON cursos.id = cursos_universidades.id_curso").
		Where("cursos_universidades.cod_area = ?", id).Scan(&CursosArea)

	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró Cursos para el area : "+id)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, CursosArea)
}

func GetCursosStudent(w http.ResponseWriter, req *http.Request) {
	cursos := modelos.Cursos{}
	student := modelos.Estudiante{}
	usuario := modelos.Usuarios{}
	cursoStudent := []modelos.CursosStudent{}

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()
	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	db.Where("usuarios_id = ?", tk.Id_Usuario).Find(&student)
	db.Where("id = ?", tk.Id_Usuario).Find(&usuario)

	result := db.Model(&cursos).Select(" DISTINCT cursos.id,cursos.nombre_curso, cursos_universidades.cod_area, estudiantes.Carr_Pref as carrera,estudiantes.Uni_Pref as universidad").
		Joins("inner join cursos_universidades on cursos.id = cursos_universidades.id_curso ").
		Joins("inner join estudiantes on cursos_universidades.cod_area = estudiantes.area_pref ").
		Where("cursos_universidades.cod_area = ? and estudiantes.Carr_Pref = ?", student.Area_Pref, student.Carr_Pref).Scan(&cursoStudent)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró Cursos para el estudiante : "+usuario.Nombres)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, cursoStudent)
}

/*
select DISTINCT cursos_universidades.id,cursos.nombre_curso, cursos_universidades.cod_area, estudiantes.Carr_Pref as carrera,estudiantes.Uni_Pref as universidad
from cursos
inner join cursos_universidades on cursos.id = cursos_universidades.id_curso
inner join estudiantes on cursos_universidades.cod_area = estudiantes.area_pref
where cursos_universidades.cod_area ='uncp2' and estudiantes.carr_pref ='INg. Sistemas'
*/
