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
	subRoute.HandleFunc("/verpreguntaRespuestas/{id}", controllers.GetPregunta).Methods("GET")
	subRoute.HandleFunc("/actualizarPreguntas/{id}", controllers.UpdatePreguntaRespuestas).Methods("PUT")
	subRoute.HandleFunc("/eliminarPreguntaRespuestas/{id}", controllers.DeletePreguntaRespuestas).Methods("DELETE")

	subRoute.HandleFunc("/insertarExamenPregunta/", controllers.SavePreguntasExamen).Methods("POST")
}
