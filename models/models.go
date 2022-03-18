package models

import "time"

type Areas struct {
	Id_Area uint   `json:"id_area" gorm:"primary_key; auto_increment"`
	Area    string `json:"area" gorm:"type:varchar(250);not null"`
}

type Asignacion_Curso struct {
	Id_Asignacion uint `json:"id_asignación" gorm:"primary_key; auto_increment"`
	Usuario_Id    uint `json:"usuario_id" gorm:"not null;references:usuarios(id);index"`
	Area_Id       uint `json:"area_id" gorm:"not null;references:id_area;index"`
}

type Consultas struct {
	Id_Consulta uint   `json:"id_consulta" gorm:"primary_key; auto_increment"`
	Consulta    string `json:"consulta" gorm:"type:varchar(255);not null"`
	Respuesta   string `json:"respuesta" gorm:"type:varchar(255);not null"`
	Tema_Id     uint   `json:"tema_id" gorm:"not null;references:id_tema;index"`
}

type Cursos_Area struct {
	Id       uint `json:"id" gorm:"primary_key; auto_increment"`
	Curso_Id uint `json:"curso_id" gorm:"not null;references:id_curso;index"`
	Area_Id  uint `json:"area_id" gorm:"not null;references:id_area;index"`
}

type Detalles_Horario struct {
	Id_Detalle_Horario uint      `json:"id_detalle_horario" gorm:"primary_key; auto_increment"`
	Hora_inicio        time.Time `json:"hora_inicio" gorm:"type:timestamp"`
	Hora_fin           time.Time `json:"hora_fin" gorm:"type:timestamp"`
	Curso_Id           uint      `json:"curso_id" gorm:"not null;references:id_curso;index"`
	Horario_Id         uint      `json:"horario_id" gorm:"not null;references:id_horario;index"`
	Usuario_Id         uint      `json:"usuario_id" gorm:"not null;references:usuarios(id);index"`
}

type Files_Cursos struct {
	Id_File uint   `json:"id_file" gorm:"primary_key; auto_increment"`
	Title   string `json:"title" gorm:"type:varchar(255);not null"`
	File    string `json:"file_name" gorm:"type:varchar(255);not null"`
	Autor   uint   `json:"autor" gorm:"not null;references:usuarios(id);index"`
	Tema_Id uint   `json:"tema_id" gorm:"not null;references:id_tema;index"`
}

type Horario struct {
	Id_Horario uint      `json:"id_horario" gorm:"primary_key; auto_increment"`
	Num_horas  int       `json:"num_horas" gorm:"type:int;not null"`
	Fecha_hora time.Time `json:"fecha_hora" gorm:"type:timestamp"`
	Curso_Id   uint      `json:"curso_id" gorm:"not null;references:id_curso;index"`
	Usuario_Id uint      `json:"usuario_id" gorm:"not null;references:usuarios(id);index"`
}

type Matricula struct {
	Id_Matricula uint `json:"id_matricula" gorm:"primary_key; auto_increment"`
	Usuario_Id   uint `json:"usuario_id" gorm:"not null;references:usuarios(id);index"`
}

type Perfil struct {
	Id_Perfil     uint `json:"id_perfil" gorm:"primary_key; auto_increment"`
	Permission_Id uint `json:"permission_id" gorm:"not null;references:id_permission;index"`
	Matricula_Id  uint `json:"matricula_id" gorm:"not null;references:id_matricula;index"`
}

type User_Permission struct {
	Id_Permission      uint `json:"id_permission" gorm:"primary_key; auto_increment"`
	Teacher_permission bool `json:"teacher_permission" gorm:"type:boolean;default:false;not null"`
	Student_permission bool `json:"student_permission" gorm:"type:boolean;default:false;not null"`
	Usuario_Id         uint `json:"usuario_id" gorm:"not null;references:usuarios(id);index"`
}
