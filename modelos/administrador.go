package modelos

type Administradors struct {
	ID         uint `json:"id" gorm:"primary_key;auto_increment"`
	Id_Usuario uint `json:"id_usuario" gorm:"type:int REFERENCES usuarios(id) "`
}
