package modelos

type RespuestaExs struct {
	ID                uint   `json:"id" gorm:"primary_key;auto_increment"`
	PreguntaExamensId uint   `json:"pregunta_ex_id" gorm:"type:int REFERENCES pregunta_examens(id) "`
	Valor             bool   `json:"valor" gorm:"type:bool"`
	Respuesta         string `json:"respuesta" gorm:"type:varchar(250)"`
	ImgLink           string `json:"img_link" gorm:"type:varchar(250)"`
}
type Result struct {
	Correct   int               `json:"correct"`
	Incorrect int               `json:"incorrect"`
	Nota      float64           `json:"nota"`
	Resultado map[string]string `json:"Resultado"`
	Solucion  map[string]uint   `json:"Solucion"`
}

type ExamenEstudiante struct {
	ID         uint    `json:"id" gorm:"primary_key;auto_increment"`
	UsuariosId uint    `json:"estudiante_id" gorm:"type:int REFERENCES usuarios(id) "`
	ExamenId   uint    `json:"examen_id" gorm:"type:int REFERENCES examens(id) "`
	Fecha      string  `json:"fecha" gorm:"type:varchar(250)"`
	Nota       float64 `json:"nota" gorm:"type:float"`
	Resultado  string  `json:"resultado" gorm:"type:varchar(250)"`
	Solucion   string  `json:"solucion" gorm:"type:varchar(250)"`
}
