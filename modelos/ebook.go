package modelos

import "gorm.io/datatypes"

type Ebooks struct {
	ID           uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_Curso     uint           `json:"id_curso" sql:"type:int REFERENCES cursos(id) "`
	Titulo       string         `json:"titulo" gorm:"type:varchar(250);not null"`
	Tema         string         `json:"tema" gorm:"type:varchar(250)"`
	Autor        string         `json:"autor" gorm:"type:varchar(250)"`
	Fecha_Public datatypes.Date `json:"fecha_public" gorm:"type:date"`
	Link         string         `json:"link" gorm:"type:varchar(250);not null"`
}
