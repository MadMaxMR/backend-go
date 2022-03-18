package modelos

import "gorm.io/datatypes"

type Clases struct {
	ID           uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_Curso     uint           `json:"id_curso" sql:"type:int REFERENCES cursos(id) "`
	Nombre_Clase string         `json:"nombre_clase" gorm:"type:varchar(250);not null"`
	Fecha_Clase  datatypes.Date `json:"fecha_clase" gorm:"type:date"`
	Hora_Inicio  datatypes.Time `json:"hora_inicio" gorm:"type:time"`
	Hora_Fin     datatypes.Time `json:"hora_fin" gorm:"type:time"`
	Link_Clase   string         `json:"link_clase" gorm:"type:varchar(250);not null"`
}
