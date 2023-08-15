package modelos

import (
	"time"
)

type HistorialExamens struct {
	ID             uint      `json:"id" gorm:"primary_key;auto_increment"`
	UsuarioId      uint      `json:"id_usuario" gorm:"REFERENCES usuarios(id) "`
	Fecha_Examen   time.Time `json:"last_login" gorm:"type:timestamp"`
	Nota_Max       float64   `json:"nota_max" gorm:"type:float"`
	Nota_Min       float64   `json:"nota_min" gorm:"type:float"`
	Nota_Tentativa float64   `json:"nota_tentativa" gorm:"type:float"`
	Respuestas     string    `json:"respuestas"`
	Solucion       string    `json:"solucion"`
	Id_Examen      uint      `json:"id_examen" gorm:"REFERENCES examens(id) "`
}

type MisExamenes struct {
	UsuarioId      uint      `json:"id_usuario"`
	UniversidadsId string    `json:"id_universidad"`
	ExamensId      uint      `json:"id_examen"`
	AreasId        string    `json:"id_area"`
	Nota           float64   `json:"nota"`
	Condicion      string    `json:"condicion"`
	Fecha_Examen   time.Time `json:"last_login" gorm:"type:timestamp"`
}

type ResMisExamen struct {
	ExamensId   string  `json:"id_examen"`
	UsuarioId   uint    `json:"id_usuario"`
	Universidad string  `json:"universidad"`
	Area        string  `json:"area"`
	Nota        float64 `json:"nota"`
	Condicion   string  `json:"condicion"`
}
