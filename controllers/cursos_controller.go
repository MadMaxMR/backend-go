package controllers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"

	"github.com/gorilla/mux"
)

func GetAllCursos(w http.ResponseWriter, req *http.Request) {
	curso := []modelos.Cursos{}
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&curso, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
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
	curso := modelos.Cursos{}
	id := mux.Vars(req)["id"]
	message, err := database.Delete(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccessMessage(w, req, http.StatusOK, message)
}

func UpdateCurso(w http.ResponseWriter, req *http.Request) {
	curso := modelos.Cursos{}
	id := mux.Vars(req)["id"]
	err := auth.ValidateBody(req, &curso)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	modelo, err := database.Update(&curso, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func UploadImage(w http.ResponseWriter, req *http.Request) {
	curso := modelos.Cursos{}

	db := database.GetConnection()
	defer db.Close()

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
	cursos := []modelos.CursosUniversidades{}

	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	result := db.Where("cod_area = ?", id).Find(&cursos)
	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusBadRequest, "No se encontró Cursos para el area : "+id)
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, cursos)
}

func GetCursosStudent(w http.ResponseWriter, req *http.Request) {
	cursos := modelos.Cursos{}
	student := modelos.Estudiante{}
	usuario := modelos.Usuarios{}
	cursoStudent := []modelos.CursosStudent{}

	db := database.GetConnection()
	defer db.Close()
	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	db.Where("usuarios_id = ?", tk.Id_Usuario).Find(&student)
	db.Where("id = ?", tk.Id_Usuario).Find(&usuario)

	//result := db.Where("cod_area = ?", student.Area_Pref).Find(&cursos)
	result := db.Model(&cursos).Select(" DISTINCT cursos_universidades.id,cursos.nombre_curso, cursos_universidades.cod_area, estudiantes.Carr_Pref as carrera,estudiantes.Uni_Pref as universidad").
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
