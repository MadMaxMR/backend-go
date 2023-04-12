package modelos

type Videos struct {
	ID           uint   `json:"id" gorm:"primary_key;auto_increment"`
	TemasId      uint   `json:"Temas_id" gorm:"type:int REFERENCES temas(id)"`
	Titulo       string `json:"titulo" gorm:"type:varchar(250);not null"`
	Duracion     string `json:"duracion" gorm:"type:varchar(250)"`
	Valor_Puntos int    `json:"valor_puntos" gorm:"type:int"`
	Link         string `json:"link" gorm:"type:varchar(250);not null"`
	Finished     bool   `json:"finished" gorm:"type:boolean;default:false"`
	Nivel        int    `json:"nivel" gorm:"type:int"`
	ImgLink      string `json:"img_link" gorm:"type:varchar(250);default:'https://res.cloudinary.com/umachayfiles/image/upload/v1651179392/user/user-1.jpg'"`
	//SubTemasId   uint   `json:"subtemas_id" gorm:"type:int REFERENCES sub_temas(id)"`
}
