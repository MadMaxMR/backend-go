package modelos

import "gorm.io/datatypes"

type Tareas struct {
	ID            uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_curso      uint           `json:"id_curso" gorm:"type:int REFERENCES cursos(id) "`
	Id_Profesor   uint           `json:"id_profesor" gorm:"type:int REFERENCES profesors(id) "`
	Titulo        string         `json:"titulo" gorm:"type:varchar(250);not null"`
	Instruciones  string         `json:"instruciones" gorm:"type:varchar(250);not null"`
	Valor         float64        `json:"valor" gorm:"type:float"`
	Calificacion  float64        `json:"calificacion" gorm:"type:float"`
	Feedback      string         `json:"feedback" gorm:"type:varchar(250)"`
	Fecha_Cierre  datatypes.Date `json:"fecha_cierre" gorm:"type:date"`
	Fechas_Subida datatypes.Date `json:"fechas_subida" gorm:"type:date"`
}
