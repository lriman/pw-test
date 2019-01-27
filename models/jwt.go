package models

import "github.com/dgrijalva/jwt-go"

type Token struct {
	Email string
	jwt.StandardClaims
}
