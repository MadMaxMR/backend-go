package modelos

type CursosUniversidades struct {
	ID           uint   `json:"id" gorm:"primary_key;auto_increment"`
	IdCurso      uint   `json:"id_curso" gorm:"type:int REFERENCES cursos(id) "`
	CodArea      string `json:"cod_area" gorm:"type:varchar(250) REFERENCES areas(id) "`
	Nombre_Curso Cursos `json:"nombre_curso" gorm:"foreingKey:idcurso"`
}
