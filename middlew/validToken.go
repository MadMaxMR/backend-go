package middlew

import (
	"API/auth"
	"API/handler"
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
