package main

import (
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	database.Migrate()

	router := mux.NewRouter()

	routes.SetCursosRoutes(router)
	routes.SetTemasRoutes(router)
	routes.SetUsuariosRoutes(router)
	serverPort := os.Getenv("PORT")
	if serverPort == "" {
		serverPort = ":" + defaultPort
	}
	server := http.Server{
		Addr:    serverPort,
		Handler: cors.AllowAll().Handler(router),
	}

	log.Println("Starting development server at http://localhost:8000/")
	log.Println("Listening....\n \n \n")

	log.Println(server.ListenAndServe())

}
