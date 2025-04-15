package repository

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/entity"

	"gorm.io/gorm"
)

type userRepository struct {
	dbConn *gorm.DB
}

func NewUserRepository(db *gorm.DB) entity.UserRepoInterface {
	return &userRepository{
		dbConn: db,
	}
}

func (u *userRepository) StoreUser(user *entity.User) {
	db := u.dbConn
	db.Create(&user)
}

func (u *userRepository) GetByUsername(Username string) entity.User {
	var user entity.User
	db := u.dbConn
	db.Where("username = ?", Username).First(&user)
	return user
}

func (u *userRepository) GetByEmail(Email string) entity.User {
	var user entity.User
	db := u.dbConn
	db.Where("email = ?", Email).First(&user)
	return user
}

func (u *userRepository) GetByUsernameEmail(Username string, Email string) entity.User {
	var user entity.User
	db := u.dbConn
	db.Where("username = ? OR email = ?", Username, Email).First(&user)
	return user
}
