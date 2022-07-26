package modelos

type RespuestaExs struct {
	ID                uint   `json:"id" gorm:"primary_key;auto_increment"`
	PreguntaExamensId uint   `json:"pregunta_ex_id" gorm:"type:int REFERENCES pregunta_examens(id) "`
	Valor             bool   `json:"valor" gorm:"type:bool"`
	Respuesta         string `json:"respuesta" gorm:"type:varchar(250)"`
	Image             bool   `json:"image" gorm:"type:bool"`
	ImgLink           string `json:"img_link" gorm:"type:varchar(250)"`
}
type Result struct {
	Correct   int     `json:"correct"`
	Incorrect int     `json:"incorrect"`
	Nota      float64 `json:"nota"`
}
