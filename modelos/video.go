package modelos

type Videos struct {
	ID           uint   `json:"id" gorm:"primary_key;auto_increment"`
	TemasId      uint   `json:"TemasId" gorm:"type:int REFERENCES temas(id)"`
	Titulo       string `json:"titulo" gorm:"type:varchar(250);not null"`
	Duracion     string `json:"duracion" gorm:"type:varchar(250)"`
	Valor_Puntos int    `json:"valor_puntos" gorm:"type:int"`
	Link         string `json:"link" gorm:"type:varchar(250);not null"`
	Finished     bool   `json:"finished" gorm:"type:boolean;default:false"`
	Nivel        int    `json:"nivel" gorm:"type:int"`
}
