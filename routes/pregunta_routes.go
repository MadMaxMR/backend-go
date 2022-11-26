package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"

	"github.com/gorilla/mux"
)

func SetPreguntasRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/guardarPreguntas", controllers.SavePreguntasRespuestas).Methods("POST")
	subRoute.HandleFunc("/allPreguntas", controllers.GetAllPreguntas).Methods("GET")
	subRoute.HandleFunc("/preguntaCursoTema/{id}", controllers.GetPreguntasCursoTema).Methods("GET")
}
