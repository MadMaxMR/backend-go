package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"

	"github.com/gorilla/mux"
)

func SetCarrerasRoutes(route *mux.Router) {
	subRoute := route.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/carreras", controllers.GetAllCarreras).Methods("GET")
	//subRoute.HandleFunc("/carreras", SaveCarrera).Methods("POST")
	subRoute.HandleFunc("/carreras/{id}", controllers.GetCarrera).Methods("GET")
	//r.HandleFunc("/carreras/{id}", UpdateCarrera).Methods("PUT")
	//r.HandleFunc("/carreras/{id}", DeleteCarrera).Methods("DELETE")
	subRoute.HandleFunc("/carreras/uni/{id}", controllers.GetCarreraUni).Methods("GET")
	subRoute.HandleFunc("/carreras/area/{id}", controllers.GetCarreraByArea).Methods("GET")
}
