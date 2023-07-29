package controllers

import (
	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
)

func GetHistorialExamen(w http.ResponseWriter, req *http.Request) {
	var result []map[string]interface{}
	//result2 := []modelos.ResMisExamen{}

	db := database.GetConnection()
	dbc, _ := db.DB()
	defer dbc.Close()

	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	iduser, _ := strconv.Atoi(tk.Id_Usuario)

	db.Table("mis_examenes mex").
		Select("mex.usuario_id,mex.examens_id,uni.nombre_uni as universidad,ar.nombre_area as area ,mex.nota,mex.condicion").
		Joins(" inner join universidads uni on uni.id = mex.universidads_id").
		Joins(" inner join areas ar on ar.id = mex.areas_id").
		Where("usuario_id = ?", iduser).Find(&result)

	if len(result) == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No hay registradas")
		return
	}

	handler.SendSuccess(w, req, http.StatusOK, result)
}

func GetHistorialExamenDetalle(w http.ResponseWriter, req *http.Request) {

}
