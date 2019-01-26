package models

import jwt "github.com/dgrijalva/jwt-go"

//Claim es una estructura para los token
type Claim struct {
	User `json:"user"`
	jwt.StandardClaims
}
