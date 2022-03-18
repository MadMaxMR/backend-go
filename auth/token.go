package auth

import (
	"errors"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/*Token es la estructura para devolver el token y el Id de usuario*/
type Token struct {
	Id_Usuario uint   `json:"id" `
	Token      string `json:"token"`
}

/*Claim es la estructura para procesar el JWT */
type Claim struct {
	Id_Usuario uint `json:"id"`
	authorized bool `json:"authorized"`
	jwt.StandardClaims
}

var IDUsuario uint
var Authorized bool

func CreateToken(id_usuario uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = id_usuario
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte("SECRET"))
	return tokenStr, err
}

func ValidateToken(tk string) (*Claim, bool, string, error) {
	miClave := []byte("SECRET")
	claims := &Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}
	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		IDUsuario = claims.Id_Usuario
		Authorized = claims.authorized
		return claims, true, string(IDUsuario), nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}
	return claims, false, string(""), err
}
