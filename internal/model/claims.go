package entity

import (
	"github.com/golang-jwt/jwt"
)

type Claims struct {
	ID       string `json:"id"`
	Account  string `json:"account"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}
