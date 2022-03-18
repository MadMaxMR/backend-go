package modelos

import "gorm.io/datatypes"

type Horarios struct {
	ID       uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_Clase uint           `json:"id_clase" gorm:"type:int REFERENCES clases(id) "`
	Periodo  string         `json:"periodo" gorm:"type:varchar(250);not null"`
	Fecha    datatypes.Date `json:"fecha" gorm:"type:date"`
	Hora     datatypes.Time `json:"hora" gorm:"type:time"`
}
