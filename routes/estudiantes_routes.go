package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"
	"github.com/MadMaxMR/backend-go/middlew"

	"github.com/gorilla/mux"
)

func SetStudentRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/student/", controllers.GetAllStudent).Methods("GET")
	subRoute.HandleFunc("/student/", controllers.SaveStudent).Methods("POST")
	subRoute.HandleFunc("/student/{id}", middlew.UserExist(controllers.GetStudent)).Methods("GET")
	//subRoute.HandleFunc("/student/{id}", middlew.UserExist(controllers.UpdateUsuario)).Methods("PUT")
	//subRoute.HandleFunc("/student/{id}", controllers.DeleteUsuario).Methods("DELETE")
}
