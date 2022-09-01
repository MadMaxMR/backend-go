package modelos

type SubTemas struct {
	ID          uint     `json:"id" gorm:"primary_key;auto_increment"`
	TemasId     uint     `json:"temas_id" gorm:"type:int REFERENCES temas(id)"`
	NameSubtema string   `json:"name_subtema" gorm:"type:varchar(250);not null"`
	Nivel       int      `json:"nivel" gorm:"type:int"`
	Videos      []Videos `json:"videos"`
}
