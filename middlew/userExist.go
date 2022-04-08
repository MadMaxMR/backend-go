package middlew

import (
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"

	"github.com/gorilla/mux"
	"net/http"
)

func UserExist(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := modelos.Usuarios{}
		id := mux.Vars(req)["id"]
		_, err := database.Get(&user, id)
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, err.Error())
			return
		}
		next.ServeHTTP(w, req)
	}
}
