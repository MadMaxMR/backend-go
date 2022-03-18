package models

type Temas struct {
	ID          uint   `json:"id" gorm:"primary_key;auto_increment"`
	Title       string `json:"title" gorm:"type:varchar(200);not null;unique"`
	Description string `json:"description" gorm:"type:varchar(200);not null"`
	CursosID    uint   `sql:"type:int REFERENCES cursos(id)"`
}
