package modelos

type Examens struct {
	ID              uint              `json:"id" gorm:"primary_key;auto_increment"`
	Id_Uni          string            `json:"id_uni" gorm:"type:varchar(250) REFERENCES universidads(id) "`
	AreasId         string            `json:"id_area" sql:"type:varchar(250) REFERENCES areas(id) "`
	Nivel_Dif       string            `json:"nivel_dif" gorm:"type:varchar(250) "`
	Nota            float64           `json:"nota" gorm:"type:float"`
	Descripcion     string            `json:"descripcion" gorm:"type:varchar(250)"`
	Modalidad       string            `json:"modalidad" gorm:"type:varchar(250)"`
	Ciclo           string            `json:"ciclo" gorm:"type:varchar(250)"`
	PreguntaExamens []PreguntaExamens `json:"preguntas"`
}
