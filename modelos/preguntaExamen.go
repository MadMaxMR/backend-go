package modelos

type PreguntaExamens struct {
	ID           uint           `json:"id" gorm:"primary_key;auto_increment"`
	ExamensId    uint           `json:"examen_id" gorm:"type:int REFERENCES examens(id) "`
	NumQuestion  uint           `json:"num_question" gorm:"type:int"`
	Pregunta     string         `json:"pregunta" gorm:"type:varchar(250);not null"`
	Ponderacion  int            `json:"ponderacion" gorm:"type:int"`
	Curso_Preg   string         `json:"curso_preg" gorm:"type:varchar(250)"`
	RespuestaExs []RespuestaExs `json:"respuesta"`
}
