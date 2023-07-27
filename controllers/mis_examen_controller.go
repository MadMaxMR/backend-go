package controllers

import (
	"net/http"
	"strconv"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"
)

func GetMisExamens(w http.ResponseWriter, req *http.Request) {
	//var results []map[string]interface{}

	result := []modelos.ResMisExamen{}

	db := database.GetConnection()
	defer db.Close()

	tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
	if err != nil {
		handler.SendFail(w, req, http.StatusBadRequest, err.Error())
		return
	}
	iduser, _ := strconv.Atoi(tk.Id_Usuario)

	db.Raw("SELECT mex.usuario_id,mex.examens_id,uni.nombre_uni as universidad,ar.nombre_area as area ,mex.nota,mex.condicion"+
		" from mis_examenes mex inner join universidads uni on uni.id = mex.universidads_id inner"+
		" join areas ar on ar.id = mex.areas_id where usuario_id = ?", iduser).Scan(&result)

	if len(result) == 0 {
		handler.SendFail(w, req, http.StatusNotFound, "No hay registradas")
		return
	}

	handler.SendSuccess(w, req, http.StatusOK, result)
}
