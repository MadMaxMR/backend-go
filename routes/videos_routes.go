package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"

	"github.com/gorilla/mux"
)

func SetVideosRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/videos").Subrouter()

	subRoute.HandleFunc("/", controllers.GetAllVideos).Methods("GET")
	subRoute.HandleFunc("/", controllers.SaveVideo).Methods("POST")
	subRoute.HandleFunc("/{id}", controllers.GetVideo).Methods("GET")
	subRoute.HandleFunc("/{id}", controllers.UpdateVideo).Methods("PUT")
	subRoute.HandleFunc("/{id}", controllers.DeleteVideo).Methods("DELETE")
	subRoute.HandleFunc("/tema/{id}", controllers.GetVideoByTema).Methods("GET")
	subRoute.HandleFunc("/subtema/{id}", controllers.GetVideoBySubTema).Methods("GET")

}
