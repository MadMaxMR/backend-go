package modelos

type Chats struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Curso    uint   `json:"id_curso" gorm:"type:int REFERENCES cursos(id) "`
	Nombre_Chat string `json:"nombre_chat" gorm:"type:varchar(100);not null"`
}
