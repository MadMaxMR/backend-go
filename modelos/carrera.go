package modelos

type Carreras struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Uni      string `json:"id_uni" sql:"type:varchar(250) REFERENCES universidads(id) "`
	Cod_Area    string `json:"id_area" sql:"type:varchar(250) REFERENCES areas(id) "`
	Nombre_Carr string `json:"nombre_carr" gorm:"type:varchar(250) "`
	Image       string `json:"image" gorm:"type:varchar(250) "`
}
