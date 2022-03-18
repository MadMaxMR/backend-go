package models

type Videos struct {
	ID        uint   `json:"id_video" gorm:"primary_key;auto_increment"`
	Url_video string `json:"url_video" gorm:"type:varchar(200);not null;unique"`
	Nivel     int    `json:"nivel" `
	Tema_Id   uint   `json:"tema_id" gorm:"not null,type:int REFERENCES temas(id)"`
}
