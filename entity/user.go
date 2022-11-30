package entity

import (
	"context"
)

type User struct {
	Id       int    `gorm:"primaryKey" json:"userId"`
	Name     string `json:"name"`
	Username string `json:"username"  validate:"required"`
	Password string `json:"password"  validate:"required"`
	Email    string `json:"email"  validate:"required,email"`
	Role     string `json:"role"  validate:"required"`
}

type UserUCaseInterface interface {
	Index(ctx context.Context)
	InsertUser(ctx context.Context, user *User) error
	HashPassword(password string) (string, error)
}

type UserRepoInterface interface {
	StoreUser(user *User)
	GetByUsername(Username string) User
	GetByEmail(Username string) User
	GetByUsernameEmail(Username string, Password string) User
}
