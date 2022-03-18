package modelos

type Respuestas struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Pregunta uint   `json:"id_pregunta" gorm:"type:int REFERENCES preguntas(id)"`
	Respuesta   string `json:"respuesta" gorm:"type:varchar(250);not null"`
	Valor       bool   `json:"valor" gorm:"type:int"`
}
