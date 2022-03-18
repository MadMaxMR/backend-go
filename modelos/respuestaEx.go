package modelos

type RespuestaExs struct {
	ID             uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Pregunta_Ex uint   `json:"id_pregunta_ex" gorm:"type:int REFERENCES pregunta_examens(id) "`
	Valor          bool   `json:"valor" gorm:"type:bool"`
	Respuesta      string `json:"respuesta" gorm:"type:varchar(250)"`
}
