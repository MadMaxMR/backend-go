package modelos

type PermisoAccesos struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Modulo   uint   `json:"id_modulo" gorm:"type:int REFERENCES modulos(id) "`
	Nombre      string `json:"nombre" gorm:"type:varchar(250);not null"`
	Description string `json:"description" gorm:"type:varchar(250)"`
}
