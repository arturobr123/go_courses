package utils

import (
	"github.com/dgrijalva/jwt-go"
)

var JwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}
