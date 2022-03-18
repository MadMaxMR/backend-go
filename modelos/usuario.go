package modelos

import "gorm.io/datatypes"

type Usuarios struct {
	ID               uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_Perfil        uint           `json:"id_perfil" sql:"type:int REFERENCES perfil_usuarios(id) "`
	Password         string         `json:"password" gorm:"type:varchar(250);not null"`
	Nombres          string         `json:"nombres" gorm:"type:varchar(250);not null"`
	Apellidos        string         `json:"apellidos" gorm:"type:varchar(250);not null"`
	Dni              int            `json:"dni" gorm:"type:int;not null;unique"`
	Fecha_Nacimiento datatypes.Date `json:"fecha_nacimiento" gorm:"type:date"`
	Genero           string         `json:"genero" gorm:"type:varchar(200)"`
	Direccion        string         `json:"direccion" gorm:"type:varchar(250);not null"`
	Image            string         `json:"image" gorm:"type:varchar(250);not null;default:'avatar.png'"`
	Email            string         `json:"email" gorm:"type:varchar(250);not null;unique"`
	Celular          int            `json:"celular" gorm:"type:int;not null;unique"`
	Fecha_Registro   datatypes.Time `json:"fecha_registro" gorm:"type:timestamp"`
}
