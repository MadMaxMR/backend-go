package modelos

type Carreras struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Uni      uint   `json:"id_uni" sql:"type:int REFERENCES universidads(id) "`
	Id_Area     uint   `json:"ud_area" sql:"type:int REFERENCES areas(id) "`
	Descripcion string `json:"descripcion" gorm:"type:varchar(250) "`
	Nombre_Carr string `json:"nombre_carr" gorm:"type:varchar(250) "`
}
