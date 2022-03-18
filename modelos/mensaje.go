package modelos

import "gorm.io/datatypes"

type Mensajes struct {
	ID            uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_Estudiante uint           `json:"id_estudiante" gorm:"type:int REFERENCES estudiantes(id)"`
	Id_Chat       uint           `json:"id_chat" gorm:"type:int REFERENCES chats(id)"`
	Fecha_Mensaje datatypes.Time `json:"fecha_mensaje" gorm:"type:time"`
	Hora_Mensaje  datatypes.Time `json:"hora_mensaje" gorm:"type:time"`
	Texto         string         `json:"texto" gorm:"type:varchar(250);not null"`
	Image         string         `json:"image" gorm:"type:varchar(250);not null"`
}
