package modelos

import (
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/datatypes"
)

type Usuarios struct {
	ID uint `json:"id" gorm:"primary_key;auto_increment"`
	//Id_Perfil        uint           `json:"id_perfil" gorm:"type:int REFERENCES perfil_usuarios(id)"`
	Password         string         `json:"password" gorm:"type:varchar(250);not null"`
	Nombres          string         `json:"nombres" gorm:"type:varchar(250);not null"`
	Apellidos        string         `json:"apellidos" gorm:"type:varchar(250);not null"`
	Dni              int            `json:"dni" gorm:"type:int;not null;unique"`
	Fecha_Nacimiento datatypes.Date `json:"fecha_nacimiento" gorm:"type:date"`
	Genero           string         `json:"genero" gorm:"type:varchar(200)"`
	Direccion        string         `json:"direccion" gorm:"type:varchar(250);not null"`
	Image            string         `json:"image" gorm:"type:varchar(250);default:'https://res.cloudinary.com/umachayfiles/image/upload/v1651092722/user/default_fhvgoc.jpg'"`
	Email            string         `json:"email" gorm:"type:varchar(250);not null;unique"`
	Celular          int            `json:"celular" gorm:"type:int"`
	Fecha_Registro   time.Time      `json:"fecha_registro" gorm:"type:timestamp;default:current_timestamp"`
	Last_Login       time.Time      `json:"last_login" gorm:"type:timestamp"`
}
type ChangePassword struct {
	Currentpassword string
	Newpassword     string
}
type ImageUpdate struct {
	Image string `json:"image"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func BeforeSave(password string) string {
	hasedPassword, err := Hash(password)
	if err != nil {
		return err.Error()
	}
	password = string(hasedPassword)
	return password
}
