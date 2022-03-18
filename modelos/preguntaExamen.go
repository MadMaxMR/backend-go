package modelos

type PreguntaExamens struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Examen   uint   `json:"id_examen" gorm:"type:int REFERENCES examens(id) "`
	Pregunta    string `json:"pregunta" gorm:"type:varchar(250);not null"`
	Ponderacion int    `json:"ponderacion" gorm:"type:int"`
	Curso_Preg  string `json:"curso_preg" gorm:"type:varchar(250)"`
}
