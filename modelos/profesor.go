package modelos

type Profesors struct {
	ID         uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Usuario uint   `json:"id_usuario" gorm:"type:int REFERENCES usuarios(id) "`
	Area_Prof  string `json:"area_prof" gorm:"type:varchar(250);not null"`
	Grad_Acad  string `json:"grad_acad" gorm:"type:varchar(250)"`
}
