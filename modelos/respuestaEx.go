package modelos

type RespuestaExs struct {
	ID                uint   `json:"id" gorm:"primary_key;auto_increment"`
	PreguntaExamensId uint   `json:"pregunta_ex_id" gorm:"type:int REFERENCES pregunta_examens(id) "`
	Valor             bool   `json:"valor" gorm:"type:bool"`
	Respuesta         string `json:"respuesta" gorm:"type:varchar(250)"`
}
