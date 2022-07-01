package modelos

type Areas struct {
	ID          string `json:"id" gorm:"primary_key;type:varchar(250);not null"`
	Id_Uni      string `json:"id_uni" sql:"type:varchar(250) REFERENCES universidads(id) "`
	Nombre_Area string `json:"nombre_area" gorm:"type:varchar(250);not null"`
}
