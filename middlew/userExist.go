package middlew

import (
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/handler"
	"github.com/MadMaxMR/backend-go/modelos"

	"net/http"

	"github.com/gorilla/mux"
)

func UserExist(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		user := modelos.Usuarios{}
		_, err := database.Get(&user, mux.Vars(req)["id"])
		if err != nil {
			handler.SendFail(w, req, http.StatusBadRequest, err.Error())
			return
		}
		next.ServeHTTP(w, req)
	}
}
