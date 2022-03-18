package modelos

type Preguntas struct {
	ID           uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Curso     uint   `json:"id_curso" gorm:"type:int REFERENCES cursos(id) "`
	Pregunta     string `json:"pregunta" gorm:"type:varchar(250);not null"`
	Valor_puntos int    `json:"valor_puntos" gorm:"type:int"`
	Nivel_Dif    string `json:"nivel_dif" gorm:"type:varchar(250);not null"`
}
