package modelos

import "gorm.io/datatypes"

type Resolucions struct {
	ID               uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_Tarea         uint           `json:"id_tarea" gorm:"type:int REFERENCES tareas(id)"`
	Est_Entrega      string         `json:"est_entrega" gorm:"type:varchar(255)"`
	Est_Calificacion string         `json:"est_calificacion" gorm:"type:varchar(255)"`
	Fecha_Entrega    datatypes.Date `json:"fecha_entrega" gorm:"type:date"`
	Tiempo_Rest      datatypes.Date `json:"tiempo_rest" gorm:"type:date"`
	Comentario       string         `json:"comentario" gorm:"type:varchar(255)"`
}
