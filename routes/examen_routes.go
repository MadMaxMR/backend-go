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
	subRoute.HandleFunc("/examen/puntos/", controllers.GetPoints).Methods("POST")

	//Ruta para ver las preguntas de un examen en el administrador

	//Ruta para ver las preguntas de un examen en el estudiante
	subRoute.HandleFunc("/examen/preguntas/area/{id}", controllers.GetExamensPregByArea).Methods("GET")

	//Ruta para ver los examenes por año
	subRoute.HandleFunc("/examen/anio/area/{id}", controllers.GetExamensbyAnio).Methods("GET")
	
	//Ruta para ver las modalidades de los examenes
	subRoute.HandleFunc("/examen/modalidades/",controllers.GetModalidad).Methods("GET")
	
	//No carga las preguntas ni las respuestas - Antiguo
	subRoute.HandleFunc("/examen/preguntasRespuestas/{id}", controllers.GetPreguntasExamenByArea).Methods("GET")

}
