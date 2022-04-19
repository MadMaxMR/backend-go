package modelos

type Cursos struct {
	ID          uint `json:"id" gorm:"primary_key;auto_increment"`
	//Id_Profesor uint `json:"id_profesor" gorm:"type:int REFERENCES profesors(id) "`
	//Descripcion  string `json:"descripcion" gorm:"type:varchar(250)"`
	//Estado       string `json:"estado" gorm:"type:varchar(250)"`
	Nombre_Curso string `json:"nombre_curso" gorm:"type:varchar(250);not null"`
	Cod_Area     string `json:"ud_area" sql:"type:varchar(250) REFERENCES areas(id) "`
	Image        string `json:"imagen" gorm:"type:varchar(250);default:'default.jpg'"`
}

type CursosArea struct {
	ID     uint   `json:"id" gorm:"primary_key;auto_increment"`
	Cursos Cursos `json:"cursos" gorm:"foreignkey:Cod_Area"`
}
