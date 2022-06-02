package auth

import (
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Name     string
}

type JwtClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}
