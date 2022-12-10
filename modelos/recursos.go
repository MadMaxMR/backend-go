package modelos

type Recursos struct {
	Id      int    `json:"id" gorm:"primary_key;auto_increment"`
	Titulo  string `json:"titulo" gorm:"type:varchar(250);not null"`
	Link    string `json:"Link" gorm:"type:varchar(250);not null"`
	TemasID uint   `json:"temas_id" gorm:"type:int REFERENCES temas(id)"`
}
