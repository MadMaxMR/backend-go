package modelos

type Estudiantes struct {
	ID           uint   `json:"id" gorm:"primary_key;auto_increment"`
	Id_Usuario   uint   `json:"id_usuario" gorm:"type:int REFERENCES usuarios(id) "`
	Colegio_Proc string `json:"colegio_proc" gorm:"type:varchar(250);not null"`
	Uni_Pref     string `json:"uni_pref" gorm:"type:varchar(250);not null"`
	Carr_Pref    string `json:"carr_pref" gorm:"type:varchar(250);not null"`
	Grad_Acad    string `json:"grad_acad" gorm:"type:varchar(250);not null"`
	Nick         string `json:"nick" gorm:"type:varchar(250);not null"`
	Lugar_Proc   string `json:"lugar_proc" gorm:"type:varchar(250);not null"`
}
