package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"

	"github.com/gorilla/mux"
)

func SetEvalsRoutes(route *mux.Router) {
	subRoute := route.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/evals/", controllers.GetEvaluaciones).Methods("GET")
	subRoute.HandleFunc("/evals/", controllers.SaveEvaluaciones).Methods("POST")
}
