package database

import (
	"fmt"
	"log"
	"os"

	"github.com/MadMaxMR/backend-go/modelos"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func initConnection() *gorm.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_DATABASE")
	/*Coneccion con Railway*/
	//dbURI := "postgresql://postgres:pTUXyukv5bNNsB8caBOZ@containers-us-west-18.railway.app:7907/railway"
	//connStr := "postgresql://postgres:Z6csX3syUbpUwp5b5GUc@containers-us-west-18.railway.app:6849/railway"
	/*Conecciòn por .env*/
	dbURI := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s", host, port, user, dbname, password)
	//connStr := "host='" + PGHOST + "' port=5432 user=postgres dbname='" + PGDATABASE + "' password='" + PGPASSWORD + "' sslmode=disable"
	db, err := gorm.Open("postgres", dbURI)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Printf("Migrando base de datos")

	db.AutoMigrate(&modelos.MisExamenes{})

	// db.AutoMigrate(&modelos.Modulos{}, &modelos.Universidads{}, &modelos.Areas{}, &modelos.PermisoAccesos{}, &modelos.PerfilUsuarios{},
	// 	&modelos.Usuarios{}, &modelos.Plans{}, &modelos.Estudiantes{}, &modelos.Pagos{}, &modelos.Administradors{},
	// 	&modelos.ConsultaInvitados{}, &modelos.Profesors{}, &modelos.Cursos{}, &modelos.CursosUniversidades{}, &modelos.Tareas{}, &modelos.Chats{},
	// 	&modelos.Mensajes{}, &modelos.Publicacions{}, &modelos.Temas{}, &modelos.Recursos{}, &modelos.SubTemas{}, &modelos.Videos{}, &modelos.Evaluaciones{}, &modelos.Preguntas{},
	// 	&modelos.Respuestas{}, &modelos.Carreras{}, &modelos.Examens{}, &modelos.PerfilPostulante{},
	// 	&modelos.HistorialExamens{}, &modelos.PreguntaExamens{}, &modelos.RespuestaExs{}, &modelos.Ebooks{}, &modelos.Clases{},
	// 	&modelos.Horarios{}, &modelos.Resolucions{}, &modelos.Archivos{}, &modelos.Ponderacion{}, &modelos.UserTipe{}, &modelos.ExamenPreguntas{})
}

func GetConnection() *gorm.DB {
	db := initConnection()
	return db
}
