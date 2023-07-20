package modelos

type PerfilPostulante struct {
	ID         uint    `json:"id" gorm:"primary_key;auto_increment"`
	CarrerasId uint    `json:"carrera_id" gorm:"type:int REFERENCES carreras(id)"`
	Ptjmin     float32 `json:"ptjmin" gorm:"type:decimal(20,2)"`
	Ptjmax     float32 `json:"ptjmax" gorm:"type:decimal(20,2)"`
	Anio       int     `json:"anio" gorm:"type:int"`
	Vacantes   int     `json:"vacantes" gorm:"type:int"`
	Modalidad  string  `json:"modalidad" gorm:"type:varchar(50)"`
}
