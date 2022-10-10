package middlew

import (
	"net/http"

	"github.com/MadMaxMR/backend-go/auth"
	"github.com/MadMaxMR/backend-go/handler"
)

func ValidAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		tk, _, _, err := auth.ValidateToken(req.Header.Get("Authorization"))
		if err != nil {
			handler.SendFail(w, req, http.StatusUnauthorized, "Error en el Token !"+err.Error())
			return
		}
		if tk.UserTipe != "admin" {
			handler.SendFail(w, req, http.StatusUnauthorized, "Unauthorized !! "+err.Error())
		}
		next.ServeHTTP(w, req)
	}
}
