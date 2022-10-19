package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"

	"github.com/gorilla/mux"
)

func SetImageRoute(router *mux.Router) {
	subRoute := router.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/upimage", controllers.UploadImages).Methods("POST")
	subRoute.HandleFunc("/upimage64", controllers.UpImage64).Methods("GET")

}
