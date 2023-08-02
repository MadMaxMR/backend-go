package modelos

import "time"

type Examens struct {
	ID                uint              `json:"id" gorm:"primary_key;auto_increment"`
	Id_Uni            string            `json:"id_uni" gorm:"type:varchar(250) REFERENCES universidads(id) "`
	AreasId           string            `json:"id_area" sql:"type:varchar(250) REFERENCES areas(id) "`
	Descripcion       string            `json:"descripcion" gorm:"type:varchar(250)"`
	Modalidad         string            `json:"modalidad" gorm:"type:string "`
	LimitePreguntas   int               `json:"limite_preguntas" gorm:"type:int "`
	CantidadPreguntas int               `json:"cantidad_preguntas" gorm:"type:int "`
	Anio              string            `json:"anio" gorm:"type:string "`
	TipoExamen        string            `json:"tipo_examen" gorm:"type:string "`
	FechaCreacion     time.Time         `json:"fecha_creacion" gorm:"type:timestamp"`
	PreguntaExamens   []PreguntaExamens `json:"preguntas"`
}
