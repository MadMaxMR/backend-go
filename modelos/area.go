package modelos

type Areas struct {
	Codigo_Area string `json:"codigo_area" gorm:"primary_key;type:varchar(250);not null"`
	Id_Uni      string `json:"id_uni" sql:"type:varchar(250) REFERENCES universidads(codigo_uni) "`
	Nombre_Area string `json:"nombre_area" gorm:"type:varchar(250);not null"`
	Descripcion string `json:"descripcion" gorm:"type:varchar(250)"`
}
