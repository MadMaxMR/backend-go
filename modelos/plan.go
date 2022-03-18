package modelos

type Plans struct {
	ID         uint    `json:"id" gorm:"primary_key;auto_increment"`
	Id_Permiso uint    `json:"id_permiso" gorm:"type:int REFERENCES permiso_accesos(id)"`
	Tipo_Plan  string  `json:"tipo_plan" gorm:"type:varchar(250);not null"`
	Costo      float32 `json:"costo" gorm:"type:float;not null"`
	Beneficios string  `json:"benefiticios" gorm:"type:varchar(250);not null"`
}
