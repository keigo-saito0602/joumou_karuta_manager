package entity

import "github.com/golang-jwt/jwt/v4"

type Login struct {
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
}

type ResponseLogin struct {
	Bearer string `json:"bearer"  validate:"required"`
}

type JwtCustomClaims struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.RegisteredClaims
}

type AuthUCaseInterface interface {
	Login(login *Login, bearer *string) error
	CheckHashPassword(password, hash string) bool
	ValidateLogin(login *Login) error
	GenerateToken(user *User) string
}
