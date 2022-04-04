package modelos

type Estudiantes struct {
	ID           uint     `json:"id" gorm:"primary_key;auto_increment"`
	Id_Usuario   uint     `json:"id_usuario" gorm:"type:int REFERENCES usuarios(id) "`
	Uni_Pref     string   `json:"uni_pref" gorm:"type:varchar(250) REFERENCES universidads(codigo_uni) "`
	Carr_Pref    string   `json:"carr_pref" gorm:"type:varchar(250);not null"`
	Nick         string   `json:"nick" gorm:"type:varchar(250);not null"`
	Colegio_Proc string   `json:"colegio_proc" gorm:"type:varchar(250)"`
	Grad_Acad    string   `json:"grad_acad" gorm:"type:varchar(250)"`
	Lugar_Proc   string   `json:"lugar_proc" gorm:"type:varchar(250)"`
	Usuario      Usuarios `json:"usuario" gorm:"foreignkey:Id_Usuario"`
}
