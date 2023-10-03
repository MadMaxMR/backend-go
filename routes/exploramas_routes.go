package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"
	"github.com/gorilla/mux"
)

func SetExploramasRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/v2").Subrouter()

	//INDICADORES GRAFICO
	subRoute.HandleFunc("/indicadores", controllers.GetIndicadores).Methods("GET")
}

//INDICADORES GRAFICO
