package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"

	"github.com/gorilla/mux"
)

func SetUniRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/universidades", controllers.GetAllUniversidads).Methods("GET")
	//subRoute.HandleFunc("/universidads", SaveUniversidad).Methods("POST")
	subRoute.HandleFunc("/universidades/{id}", controllers.GetUniversidad).Methods("GET")
	//r.HandleFunc("/universidads/{id}", UpdateUniversidad).Methods("PUT")
	//r.HandleFunc("/universidads/{id}", DeleteUniversidad).Methods("DELETE")
}
