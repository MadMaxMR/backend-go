package modelos

type Universidads struct {
	ID            uint   `json:"id" gorm:"primary_key;auto_increment"`
	Nombre_Uni    string `json:"nombre_uni" gorm:"type:varchar(250);not null"`
	Descripcion   string `json:"descripcion" gorm:"type:varchar(250)"`
	Sede_Princ    string `json:"sede_princ" gorm:"type:varchar(250)"`
	Sector        string `json:"sector" gorm:"type:varchar(250)"`
	Ecuacion_Pond string `json:"ecuacion_pond" gorm:"type:varchar(250)"`
}
