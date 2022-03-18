package modelos

type ConsultaInvitados struct {
	ID        uint   `json:"id" gorm:"primary_key;auto_increment"`
	Nombre    string `json:"nombre" gorm:"type:varchar(250);not null"`
	Apellidos string `json:"apellidos" gorm:"type:varchar(250);not null"`
	Correo    string `json:"correo" gorm:"type:varchar(250);not null"`
	Celular   int    `json:"celular" gorm:"type:int"`
	Contenido string `json:"contenido" gorm:"type:varchar(250);not null"`
}
