package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"

	"github.com/gorilla/mux"
)

func SetAreasRoutes(route *mux.Router) {
	subRoute := route.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/areas", controllers.GetAllAreas).Methods("GET")
	//subRoute.HandleFunc("/areas", SaveArea).Methods("POST")
	subRoute.HandleFunc("/areas/{id}", controllers.GetArea).Methods("GET")
	//r.HandleFunc("/areas/{id}", UpdateArea).Methods("PUT")
	//r.HandleFunc("/areas/{id}", DeleteArea).Methods("DELETE")

	subRoute.HandleFunc("/areas/uni/{id}", controllers.GetAreaByUni).Methods("GET")
	subRoute.HandleFunc("/areas/uni/{id}/carreras", controllers.GetAreaCarrerasByUni).Methods("GET")
}
