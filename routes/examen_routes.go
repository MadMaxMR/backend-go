package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"

	"github.com/gorilla/mux"
)

func SetExamenRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/examen/", controllers.GetAllExamens).Methods("GET")
	subRoute.HandleFunc("/examen/{id}", controllers.GetExamenById).Methods("GET")
	subRoute.HandleFunc("/examen/", controllers.SaveExamens).Methods("POST")
	subRoute.HandleFunc("/examen/{id}", controllers.UpdateExamen).Methods("PUT")
	subRoute.HandleFunc("/examen/{id}", controllers.DeleteExamen).Methods("DELETE")
	subRoute.HandleFunc("/examen/preguntas/area/{id}", controllers.GetExamensPregByArea).Methods("GET")
	subRoute.HandleFunc("/examen/puntos/", controllers.GetPoints).Methods("POST")

	//Ruta para ver las preguntas de un examen en el administrador
	subRoute.HandleFunc("/examen/preguntas/{id}", controllers.GetPreguntasByExamen).Methods("GET")
}
