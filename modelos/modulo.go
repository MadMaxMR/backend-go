package modelos

type Modulos struct {
	ID           uint   `json:"id" gorm:"primary_key;auto_increment"`
	NombreModulo string `json:"nombre_modulo" gorm:"type:varchar(250);not null"`
	Descripcion  string `json:"descripcion" gorm:"type:varchar(250)"`
}
