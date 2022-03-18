package main

import (
	"API/database"
	"API/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	database.Migrate()

	router := mux.NewRouter()

	routes.SetCursosRoutes(router)
	routes.SetTemasRoutes(router)
	routes.SetUsuariosRoutes(router)

	server := http.Server{
		Addr:    ":8000",
		Handler: cors.AllowAll().Handler(router),
	}

	log.Println("Starting development server at http://localhost:8000/")
	log.Println("Listening....\n \n \n")

	log.Println(server.ListenAndServe())

}
