package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/MadMaxMR/backend-go/routes"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

const defaultPort = "8000"

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error al cargar variables de entorno")
	}
}

func main() {
	serverPort := os.Getenv("PORT")
	//database.Migrate()

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
	routes.SetEvalsRoutes(router)
	routes.SetExamenRoutes(router)
	routes.SetImageRoute(router)
	routes.SetPreguntasRoutes(router)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Servicios rest ACADEMIA-UMACHAY version 1.2"))
	}).Methods("GET")

	if serverPort == "" {
		serverPort = defaultPort
	}
	server := http.Server{
		Addr:    ":" + serverPort,
		Handler: cors.AllowAll().Handler(router),
	}
	log.Printf("Starting on PORT: " + serverPort)
	log.Fatal(server.ListenAndServe())

}
