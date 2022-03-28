package models

import jwt "github.com/dgrijalva/jwt-go"

type Claim struct {
	Id_Usuario string `json:"id"`
	Authorized bool   `json:"authorized"`
	jwt.StandardClaims
}
