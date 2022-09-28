package modelos

type Ponderacion struct {
	ID          uint    `json:"id" gorm:"primary_key;auto_increment"`
	CursosId    uint    `json:"id_curso" gorm:"type:int REFERENCES cursos(id) "`
	CodArea     string  `json:"cod_area" gorm:"type:varchar(250) REFERENCES areas(id) "`
	Ponderacion float64 `json:"ponderacion" gorm:"type:float"`
	Preguntas   int     `json:"reguntas" gorm:"type:int"`
}
