package modelos

import "gorm.io/datatypes"

type Archivos struct {
	ID            uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_Resolucion uint           `json:"id" gorm:"type:int REFERENCES resolucions(id) "`
	Nombre        string         `json:"nombre" gorm:"type:varchar(250);not null"`
	Fecha_Creac   datatypes.Date `json:"fecha_creac" gorm:"type:date"`
}
