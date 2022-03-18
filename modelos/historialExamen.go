package modelos

import "gorm.io/datatypes"

type HistorialExamens struct {
	ID             uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_Examen      uint           `json:"id_examen" gorm:"type:int REFERENCES examens(id) "`
	Fecha_Exame    datatypes.Date `json:"fecha_exame" gorm:"type:date"`
	Nota_Max       float64        `json:"nota_max" gorm:"type:float"`
	Nota_Min       float64        `json:"nota_min" gorm:"type:float"`
	Nota_Tentativa float64        `json:"nota_tentativa" gorm:"type:float"`
}
