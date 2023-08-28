package controllers

import (
	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
)

func GetHistorialExamen(w http.ResponseWriter, req *http.Request) {
	type Result2 struct {
		Page      string
		Prev      bool
		Next      bool
		Total     int
		Historial []map[string]interface{}
	}

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
	var result []map[string]interface{}
	result2 := Result2{}

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	iduser, _ := strconv.Atoi(tk.Id_Usuario)

	resultQ := db.Table("mis_examenes mex").Where("usuario_id = ?", iduser).Find(&result)

	if resultQ.RowsAffected == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No hay registros")
		return
	}

	result2.Page = page
	result2.Next = true
	if pageInt == 1 {
		result2.Prev = false
	}
	if pageInt > 1 {
		result2.Prev = true
	}

	if int(resultQ.RowsAffected)%pageSize == 0 {
		result2.Total = int(resultQ.RowsAffected) / pageSize
	} else {
		result2.Total = (int(resultQ.RowsAffected) / pageSize) + 1
	}

	if pageInt == result2.Total {
		result2.Next = false
	}

	result = nil
	db.Table("mis_examenes mex").
		Select("mex.usuario_id,mex.examens_id,uni.nombre_uni as universidad,ar.nombre_area as area ,mex.nota,mex.condicion").
		Joins(" inner join universidads uni on uni.id = mex.universidads_id").
		Joins(" inner join areas ar on ar.id = mex.areas_id").
		Where("usuario_id = ?", iduser).
		Limit(pageSize).Offset((pageInt - 1) * pageSize).Order("fecha_examen DESC").Find(&result)

	if len(result) == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No hay registros")
		return
	}

	result2.Historial = result

	handler.SendSuccess(w, req, http.StatusOK, result2)
}

func GetHistorialExamenDetalle(w http.ResponseWriter, req *http.Request) {
	//var result []map[string]interface{}

}

