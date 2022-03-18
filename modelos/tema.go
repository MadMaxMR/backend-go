package modelos

type Temas struct {
	ID          uint    `json:"id" gorm:"primary_key;auto_increment"`
	Id_Curso    uint    `json:"id_curso" gorm:"type:int REFERENCES cursos(id)"`
	Nombre_Tema string  `json:"nombre_tema" gorm:"type:varchar(250);not null"`
	Recurrencia float64 `json:"recurrencia" gorm:"type:float"`
}
