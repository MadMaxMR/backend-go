package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Usuarios struct {
	ID                    uint      `json:"id" gorm:"primary_key;auto_increment"`
	Email                 string    `json:"email" gorm:"type:varchar(200);not null;unique"`
	Password              string    `json:"password" gorm:"type:varchar(250);not null"`
	Nombres               string    `json:"nombres" gorm:"type:varchar(250);not null"`
	Apellidos             string    `json:"apellidos" gorm:"type:varchar(250);not null"`
	Dni                   int       `json:"dni" gorm:"type:int;not null;unique"`
	Universidad           string    `json:"universidad" gorm:"type:varchar(250);not null"`
	Celular               int       `json:"celular" gorm:"type:int;not null;unique"`
	Imagen                string    `json:"imagen" gorm:"type:varchar(250);default:'avatar.png'"`
	Area                  int       `json:"area" gorm:"type:int"`
	Is_Active             bool      `json:"is_active" gorm:"type:boolean;default:true"`
	Usuario_Administrador bool      `json:"usuario_administrador" gorm:"type:boolean;default:false"`
	Last_Login            time.Time `json:"last_login" gorm:"type:timestamp"`
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
