package modelos

type Carreras struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Uni      string `json:"id_uni" sql:"type:varchar(250) REFERENCES universidads(codigo_uni) "`
	Cod_Area    string `json:"ud_area" sql:"type:varchar(250) REFERENCES areas(codigo_area) "`
	Descripcion string `json:"descripcion" gorm:"type:varchar(250) "`
	Nombre_Carr string `json:"nombre_carr" gorm:"type:varchar(250) "`
}
