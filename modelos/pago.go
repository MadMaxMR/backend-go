package modelos

import "gorm.io/datatypes"

type Pagos struct {
	ID             uint           `json:"id" gorm:"primary_key;auto_increment"`
	Id_Estudiante  uint           `json:"id_estudiante" gorm:"type:int REFERENCES estudiantes(id)"`
	Id_Plan        uint           `json:"id_plan" gorm:"type:int REFERENCES plans(id)"`
	Fecha_Pago     datatypes.Date `json:"fecha_pago" gorm:"type:date"`
	Modalidad      string         `json:"modalidad" gorm:"type:varchar(250);not null"`
	Importe_Total  float64        `json:"importe_total" gorm:"type:decimal(20,8);not null"`
	Nombre_Cliente string         `json:"nombre_cliente" gorm:"type:varchar(250);not null"`
	Estado_Pago    string         `json:"estado_pago" gorm:"type:varchar(250);not null"`
	Metodo_Pago    string         `json:"metodo_pago" gorm:"type:varchar(250);not null"`
}
