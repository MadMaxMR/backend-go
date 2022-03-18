package modelos

type PerfilUsuarios struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Permiso  uint   `json:"id_permiso" gorm:"type:int REFERENCES permiso_accesos(id)"`
	Description string `json:"description" gorm:"type:varchar(250)"`
}
