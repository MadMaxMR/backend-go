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

const defaultPort = "8000"

func main() {
	serverPort := os.Getenv("PORT")
	database.Migrate()

	router := mux.NewRouter()

	routes.ResetPasswordRoutes(router)
	routes.SetUniRoutes(router)
	routes.SetAreasRoutes(router)
	routes.SetCarrerasRoutes(router)
	routes.SetStudentRoutes(router)
	routes.SetCursosRoutes(router)
	routes.SetTemasRoutes(router)
	routes.SetUsuariosRoutes(router)
	routes.SetVideosRoutes(router)
	if serverPort == "" {
		serverPort = defaultPort
	}
	server := http.Server{
		Addr:    ":" + serverPort,
		Handler: cors.AllowAll().Handler(router),
	}
	log.Printf("Starting")
	log.Fatal(server.ListenAndServe())

}
