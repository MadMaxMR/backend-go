package routes

import (
	"github.com/MadMaxMR/backend-go/controllers"
	"github.com/MadMaxMR/backend-go/middlew"

	"github.com/gorilla/mux"
)

func SetUsuariosRoutes(router *mux.Router) {
	subRoute := router.PathPrefix("/").Subrouter()

	subRoute.HandleFunc("/usuarios/", controllers.GetAllUsuarios).Methods("GET")
	subRoute.HandleFunc("/usuarios/", controllers.SaveUsuario).Methods("POST")
	subRoute.HandleFunc("/usuarios/{id}", controllers.GetUsuario).Methods("GET")
	subRoute.HandleFunc("/usuarios/{id}", middlew.UserExist(controllers.UpdateUsuario)).Methods("PUT")
	subRoute.HandleFunc("/usuarios/{id}", controllers.DeleteUsuario).Methods("DELETE")
	subRoute.HandleFunc("/login/", controllers.Login).Methods("POST")

	subRoute.HandleFunc("/verperfil/{id}", middlew.ValidToken(controllers.VerPerfil)).Methods("GET")
	subRoute.HandleFunc("/avatar/{id}", controllers.UpdateAvatar1).Methods("PUT")
	subRoute.HandleFunc("/avatar/{id}", controllers.GetAvatar1).Methods("GET")

	subRoute.HandleFunc("/updateAvatar/", controllers.SaveAvatar).Methods("POST")
	subRoute.HandleFunc("/deleteAvatar/{id}", middlew.UserExist(controllers.UpdateAvatar)).Methods("PUT")
	subRoute.HandleFunc("/changepassword/", controllers.ChangePassword).Methods("PUT")
}
