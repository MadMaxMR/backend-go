package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"
	"github.com/MadMaxMR/backend-go/middlew"
	"github.com/gorilla/mux"
)

func SetCursosRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/cursos/", controllers.GetAllCursos).Methods("GET")
	subRoute.HandleFunc("/cursos/", controllers.SaveCurso).Methods("POST")
	subRoute.HandleFunc("/cursos/{id}", controllers.GetCurso).Methods("GET")
	subRoute.HandleFunc("/cursos/{id}", controllers.DeleteCurso).Methods("DELETE")
	subRoute.HandleFunc("/cursos/{id}", controllers.UpdateCurso).Methods("PUT")

	subRoute.HandleFunc("/cursos/image/{id}", controllers.UploadImage).Methods("POST")
	subRoute.HandleFunc("/cursos/image/{id}", controllers.GetImage).Methods("GET")

	subRoute.HandleFunc("/cursos/area/{id}", controllers.GetCursoByArea).Methods("GET")
	subRoute.HandleFunc("/cursos/student/", middlew.ValidToken(controllers.GetCursosStudent)).Methods("GET")
}
