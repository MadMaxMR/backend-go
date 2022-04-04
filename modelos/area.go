package modelos

type Areas struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Uni      string `json:"id_uni" sql:"type:varchar(250) REFERENCES universidads(codigo_uni) "`
	Nombre_Area string `json:"nombre_area" gorm:"type:varchar(250);not null"`
	Descripcion string `json:"descripcion" gorm:"type:varchar(250)"`
}
