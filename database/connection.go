package database

import (
	"log"

	"github.com/MadMaxMR/backend-go/modelos"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func initConnection() *gorm.DB {
	/*Coneccion con ElephantSQL*/
	//connStr := "postgres://arwpboxu:qP449bZjdC9jEpih47th8Hn21yi2Aj6h@motty.db.elephantsql.com/arwpboxu"
	/*Coneccion con Heroku*/
	connStr := "postgres://umachay:PoEGxikOL6atTxJrJGvkihikgRCmsZEJ@dpg-cfa8h9pgp3jsh6eefoo0-a.frankfurt-postgres.render.com/umachay"
	//connStr := "postgresql://postgres:Z6csX3syUbpUwp5b5GUc@containers-us-west-18.railway.app:6849/railway"
	//postgresql://${{ PGUSER }}:${{ PGPASSWORD }}@${{ PGHOST }}:${{ PGPORT }}/${{ PGDATABASE }}
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

	db.AutoMigrate(&modelos.Modulos{}, &modelos.Universidads{}, &modelos.Areas{}, &modelos.PermisoAccesos{}, &modelos.PerfilUsuarios{},
		&modelos.Usuarios{}, &modelos.Plans{}, &modelos.Estudiantes{}, &modelos.Pagos{}, &modelos.Administradors{},
		&modelos.ConsultaInvitados{}, &modelos.Profesors{}, &modelos.Cursos{}, &modelos.CursosUniversidades{}, &modelos.Tareas{}, &modelos.Chats{},
		&modelos.Mensajes{}, &modelos.Publicacions{}, &modelos.Temas{}, &modelos.Recursos{}, &modelos.SubTemas{}, &modelos.Videos{}, &modelos.Evaluaciones{}, &modelos.Preguntas{},
		&modelos.Respuestas{}, &modelos.Carreras{}, &modelos.Examens{},
		&modelos.HistorialExamens{}, &modelos.PreguntaExamens{}, &modelos.RespuestaExs{}, &modelos.Ebooks{}, &modelos.Clases{},
		&modelos.Horarios{}, &modelos.Resolucions{}, &modelos.Archivos{}, &modelos.Ponderacion{}, &modelos.UserTipe{}, &modelos.ExamenPreguntas{})
}

func GetConnection() *gorm.DB {
	db := initConnection()
	return db
}
