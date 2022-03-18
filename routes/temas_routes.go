package routes

import (
	"API/controllers"

	"github.com/gorilla/mux"
)

func SetTemasRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/temas/", controllers.GetAllTemas).Methods("GET")
	subRoute.HandleFunc("/temas/", controllers.SaveTema).Methods("POST")
	subRoute.HandleFunc("/temas/{id}", controllers.GetTema).Methods("GET")
	subRoute.HandleFunc("/temas/{id}", controllers.DeleteTema).Methods("DELETE")
	subRoute.HandleFunc("/temas/{id}", controllers.UpdateTema).Methods("PUT")
	subRoute.HandleFunc("/temas/curso/{id}", controllers.GetTemaByCurso).Methods("GET")
}
