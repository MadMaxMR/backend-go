package modelos

type ExamenPreguntas struct {
	ID                uint `json:"id" gorm:"primary_key;auto_increment"`
	ExamensId         uint `json:"id_examen" gorm:"type:int REFERENCES examens(id) "`
	PreguntaExamensId uint `json:"id_pregunta" gorm:"type:int REFERENCES pregunta_examens(id) "`
	// NumQuestion       int  `json:"num_questions" gorm:"type:int" `
}
