package middlew

import (
	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/handler"
	"log"
	"net/http"
)

func ValidToken(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		_, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
		if err != nil {
			log.Println(err)
			handler.SendFail(w, req, http.StatusUnauthorized, "Error en el Token !"+err.Error())
			return
		}
		next.ServeHTTP(w, req)
	}
}
