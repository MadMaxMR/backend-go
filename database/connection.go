package database

import (
	"log"

	"API/modelos"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func GetConnection() *gorm.DB {
	connStr := "postgres://postgres:123456@localhost:5432/umachay1?sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Printf("Migrando base de datos")
	/*db.AutoMigrate(&models.Cursos{}, &models.Temas{}, &models.Videos{}, &models.Usuarios{}, &models.Areas{}, &models.Asignacion_Curso{},
	&models.Consultas{}, &models.Cursos_Area{}, &models.Detalles_Horario{}, &models.Files_Cursos{}, &models.Horario{},
	&models.Matricula{}, &models.Perfil{}, &models.User_Permission{})*/

	db.AutoMigrate(&modelos.Modulos{}, &modelos.PermisoAccesos{}, &modelos.PerfilUsuarios{},
		&modelos.Usuarios{}, &modelos.Plans{}, &modelos.Estudiantes{}, &modelos.Pagos{}, &modelos.Administradors{},
		&modelos.ConsultaInvitados{}, &modelos.Profesors{}, &modelos.Cursos{}, &modelos.Tareas{}, &modelos.Chats{},
		&modelos.Mensajes{}, &modelos.Publicacions{}, &modelos.Temas{}, &modelos.Videos{}, &modelos.Preguntas{},
		&modelos.Respuestas{}, &modelos.Universidads{}, &modelos.Areas{}, &modelos.Carreras{}, &modelos.Examens{},
		&modelos.HistorialExamens{}, &modelos.PreguntaExamens{}, &modelos.RespuestaExs{}, &modelos.Ebooks{}, &modelos.Clases{},
		&modelos.Horarios{}, &modelos.Resolucions{}, &modelos.Archivos{})
}
