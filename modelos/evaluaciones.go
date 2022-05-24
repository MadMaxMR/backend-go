package modelos

type Evaluaciones struct {
	ID      uint   `json:"id" gorm:"primary_key;auto_increment"`
	TemasID uint   `json:"temas_id" gorm:"type:int REFERENCES temas(id)"`
	Nivel   string `json:"nivel" gorm:"type:varchar(250);not null"`
	PdfLink string `json:"pdf_link" gorm:"type:varchar(250);not null"`
}
