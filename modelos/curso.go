package modelos

type Cursos struct {
	ID           uint   `json:"id" gorm:"primary_key;auto_increment"`
	Nombre_Curso string `json:"nombre_curso" gorm:"type:varchar(250);not null"`
	Image        string `json:"imagen" gorm:"type:varchar(250);default:'default.jpg'"`
	//Id_Profesor  uint   `json:"id_profesor" gorm:"type:int REFERENCES profesors(id) "`
}

type CursosStudent struct {
	ID           uint   `json:"id"`
	Nombre_Curso string `json:"nombre_curso"`
	Cod_Area     string `json:"cod_area"`
	Image        string `json:"imagen"`
	Carrera      string `json:"carrera"`
	Universidad  string `json:"universidad"`
}

type CursosArea struct {
	ID           uint   `json:"id"`
	Nombre_Curso string `json:"nombre_curso"`
	Cod_Area     string `json:"cod_area"`
}
