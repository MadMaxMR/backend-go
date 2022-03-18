package main

import (
	"github.com/MadMaxMR/backend-go/database"
	"github.com/MadMaxMR/backend-go/routes"
	"github.com/gorilla/mux"
	//"github.com/rs/cors"
	"log"
	"net/http"
	"os"
)

const defaultPort = "8080"

func main() {
	serverPort := os.Getenv("PORT")
	database.Migrate()

	router := mux.NewRouter()

	routes.SetCursosRoutes(router)
	routes.SetTemasRoutes(router)
	routes.SetUsuariosRoutes(router)
	if serverPort == "" {
		serverPort = defaultPort
	}
	server := http.Server{
		Addr:    ":" + serverPort,
		Handler: router,
	}
	log.Fatal(server.ListenAndServe())

}
