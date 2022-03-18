package modelos

type Publicacions struct {
	ID         uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Mensaje uint   `json:"id_mensaje" gorm:"type:int REFERENCES mensajes(id) "`
	Reaccion   string `json:"reaccion" gorm:"type:varchar(250)"`
}
