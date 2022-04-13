package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"
	"github.com/gorilla/mux"
)

func ResetPasswordRoutes(r *mux.Router) {
	router := r.PathPrefix("/").Subrouter()

	router.HandleFunc("/reset", controllers.ResetPassword).Methods("POST")
	router.HandleFunc("/recovery/password", controllers.RecoveryPassword).Methods("POST")
}
