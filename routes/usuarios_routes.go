package routes

import (
	"API/controllers"
	"API/middlew"

	"github.com/gorilla/mux"
)

func SetUsuariosRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/usuarios/", controllers.GetAllUsuarios).Methods("GET")
	subRoute.HandleFunc("/usuarios/", controllers.SaveUsuario).Methods("POST")
	subRoute.HandleFunc("/usuarios/{id}", controllers.GetUsuario).Methods("GET")
	subRoute.HandleFunc("/usuarios/{id}", controllers.UpdateUsuario).Methods("PUT")
	subRoute.HandleFunc("/usuarios/{id}", controllers.DeleteUsuario).Methods("DELETE")
	subRoute.HandleFunc("/login/", controllers.Login).Methods("POST")

	subRoute.HandleFunc("/verperfil/{id}", middlew.ValidToken(controllers.VerPerfil)).Methods("GET")
	subRoute.HandleFunc("/avatar/{id}", controllers.UpdateAvatar).Methods("PUT")
	subRoute.HandleFunc("/avatar/{id}", controllers.GetAvatar).Methods("GET")
}
