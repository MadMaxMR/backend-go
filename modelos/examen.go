package modelos

type Examens struct {
	ID          uint    `json:"id" gorm:"primary_key;auto_increment"`
	Id_Uni      uint    `json:"id_uni" gorm:"type:int REFERENCES universidads(id) "`
	Id_Carrera  uint    `json:"id_carrera" gorm:"type:int REFERENCES carreras(id) "`
	Nivel_Dif   string  `json:"nivel_dif" gorm:"type:varchar(250);not null"`
	Nota        float64 `json:"nota" gorm:"type:float"`
	Descripcion string  `json:"descripcion" gorm:"type:varchar(250)"`
	Modalidad   string  `json:"modalidad" gorm:"type:varchar(250)"`
	Ciclo       string  `json:"ciclo" gorm:"type:varchar(250)"`
}
