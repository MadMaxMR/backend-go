package auth

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
)

/*Token es la estructura para devolver el token y el Id de usuario*/
type Token struct {
	Id_Usuario uint   `json:"id" `
	Token      string `json:"token"`
	UserTipe   string `json:"user_type"`
}
type Claim struct {
	Id_Usuario string `json:"id"`
	Authorized bool   `json:"authorized"`
	UserTipe   string `json:"userTipe"`
	jwt.StandardClaims
}

var IDUsuario uint
var Authorized bool

func CreateToken(id_usuario uint, tipeUser string) (string, error) {
	claims := jwt.MapClaims{}
	claims["userTipe"] = tipeUser
	claims["authorized"] = true
	claims["user_id"] = id_usuario
	claims["exp"] = time.Now().Add(time.Hour * (24)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString([]byte("SECRET"))
	return tokenStr, err
}

func CreateTokenReset(email string, id uint) (string, error) {
	claims := jwt.MapClaims{}
	claims["id"] = id
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Minute * 5).Unix()
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

	tkn, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err != nil {
		return claims, false, string(""), err
	}
	claim, ok := tkn.Claims.(jwt.MapClaims)
	if ok && tkn.Valid {
		claims.Id_Usuario = strconv.FormatUint(uint64(uint(claim["user_id"].(float64))), 10)
		claims.Authorized = claim["authorized"].(bool)
		claims.UserTipe = claim["userTipe"].(string)
		return claims, true, claims.Id_Usuario, nil
	}
	return claims, false, string(""), errors.New("token invalido")

}

func ValidateTokenReset(tk string) (string, error) {
	miClave := []byte("SECRET")
	tkn, err := jwt.Parse(tk, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err != nil {
		return "", err
	}
	claim, ok := tkn.Claims.(jwt.MapClaims)
	if ok && tkn.Valid {
		return string(claim["email"].(string)), nil
	}
	return "", errors.New("Token invalido")
}
