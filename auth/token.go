package auth

import (
	"errors"
	"fmt"
	"github.com/MadMaxMR/backend-go/models"
	"strconv"
	"strings"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

/*Token es la estructura para devolver el token y el Id de usuario*/
type Token struct {
	Id_Usuario uint   `json:"id" `
	Token      string `json:"token"`
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

func ValidateToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("SECRET")
	claims := &models.Claim{}

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
		return claims, true, string(IDUsuario), nil
	}
	return claims, false, string(""), errors.New("token invalido")

}
