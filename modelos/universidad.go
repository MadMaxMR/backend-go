package modelos

type Universidads struct {
	ID            string `json:"id" gorm:"primary_key;type:varchar(50);not null"`
	Nombre_Uni    string `json:"nombre_uni" gorm:"type:varchar(250);not null"`
	Descripcion   string `json:"descripcion" gorm:"type:varchar(250)"`
	Sede_Princ    string `json:"sede_princ" gorm:"type:varchar(250)"`
	Sector        string `json:"sector" gorm:"type:varchar(250)"`
	Ecuacion_Pond string `json:"ecuacion_pond" gorm:"type:varchar(250)"`
	Ranking       int    `json:"ranking" gorm:"type:int"`
	Image         string `json:"image" gorm:"type:varchar(250)"`
}
