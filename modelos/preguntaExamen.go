package modelos

type PreguntaExamens struct {
	ID           uint           `json:"id" gorm:"primary_key;auto_increment"`
	ExamensId    uint           `json:"examen_id" gorm:"type:int REFERENCES examens(id) "`
	Enunciado1   string         `json:"enunciado1" gorm:"type:varchar(250)"`
	Grafico      string         `json:"grafico" gorm:"type:varchar(250)"`
	Enunciado2   string         `json:"enunciado2" gorm:"type:varchar(250)"`
	Enunciado3   string         `json:"enunciado3" gorm:"type:varchar(250)"`
	NumQuestion  uint           `json:"num_question" gorm:"type:int"`
	CursosId     uint           `json:"curso_id" gorm:"type:int REFERENCES cursos(id)"`
	TemasId      uint           `json:"tema_id" gorm:"type:int REFERENCES temas(id)"`
	RespuestaExs []RespuestaExs `json:"respuesta" form:"respuesta"`
}
