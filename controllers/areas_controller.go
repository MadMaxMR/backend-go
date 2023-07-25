package controllers

import (
	"strconv"

	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
	"github.com/jinzhu/gorm"

	"net/http"

	"github.com/gorilla/mux"
)

func GetAllAreas(w http.ResponseWriter, req *http.Request) {
	areas := []modelos.Areas{}
	db := database.GetConnection()
	defer db.Close()
	page := req.URL.Query().Get("page")
	modelo, err := database.GetAll(&areas, page)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)
}

func GetArea(w http.ResponseWriter, req *http.Request) {
	area := modelos.Areas{}
	id := mux.Vars(req)["id"]

	modelo, err := database.Get(&area, id)
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, modelo)

}

func GetAreaByUni(w http.ResponseWriter, req *http.Request) {
	areas := []modelos.Areas{}
	id := mux.Vars(req)["id"]

	db := database.GetConnection()
	defer db.Close()

	result := db.Where("id_uni = ?", id).Find(&areas)
	if result.Error != nil {
		handler.SendFail(w, req, http.StatusBadRequest, result.Error.Error())
		return
	}
	handler.SendSuccess(w, req, http.StatusOK, areas)
}

// GetAreaCarrerasByUni retorna todas las areas de una "universidad" incluido las carreras de sus areas
func GetAreaCarrerasByUni(w http.ResponseWriter, req *http.Request) {
	type Result2 struct {
		Page  string
		Prev  bool
		Next  bool
		Total int
		Data  []modelos.Universidads
	}
	result2 := Result2{}

	universidads := []modelos.Universidads{}
	idCarrera := req.URL.Query().Get("idCarrera")
	idArea := req.URL.Query().Get("idArea")
	idUniversidad := req.URL.Query().Get("idUniversidad")
	page := req.URL.Query().Get("page")
	pageSizes := req.URL.Query().Get("pageSize")
	if page == "" {
		page = "1"
	}
	if pageSizes == "" {
		pageSizes = "20"
	}
	pageInt, _ := strconv.Atoi(page)
	pageSize, _ := strconv.Atoi(pageSizes)

	if idUniversidad == "" {
		idUniversidad = "%"
	}
	if idArea == "" {
		idArea = "%"
	}
	if idCarrera == "" {
		idCarrera = "%"
	}

	db := database.GetConnection()
	defer db.Close()

	result := db.Where("id LIKE ?", idUniversidad).Preload("Area", "id LIKE ?", idArea).Preload("Area.Carreras", "id::text LIKE ?", idCarrera, func(db *gorm.DB) *gorm.DB {
		return db.Order("Carreras.nombre_carr ASC")
	}).Preload("Area.Carreras.PerfilPostulante").Find(&universidads)

	if result.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusInternalServerError, "No se encontró areas para la universidad solicitada")
		return
	}
	var totpep int
	for _, universidad := range universidads {
		for _, area := range universidad.Area {
			for _, carrera := range area.Carreras {
				totpep = totpep + len(carrera.PerfilPostulante)
			}
		}
	}

	result2.Page = page
	result2.Next = true
	if pageInt == 1 {
		result2.Prev = false
	}
	if pageInt > 1 {
		result2.Prev = true
	}

	if int(totpep)%pageSize == 0 {
		result2.Total = int(totpep) / pageSize
	} else {
		result2.Total = (int(totpep) / pageSize) + 1
	}

	if pageInt == result2.Total {
		result2.Next = false
	}

	db.Debug().Where("id LIKE ?", idUniversidad).Preload("Area", "id LIKE ?", idArea).Preload("Area.Carreras", "id::text LIKE ?", idCarrera, func(db *gorm.DB) *gorm.DB {
		return db.Order("id ASC")
	}).Preload("Area.Carreras.PerfilPostulante", func(db *gorm.DB) *gorm.DB {
		return db.Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("carreras_id ASC")
	}).Find(&universidads)

	result2.Data = universidads

	handler.SendSuccess(w, req, http.StatusOK, result2)
}
