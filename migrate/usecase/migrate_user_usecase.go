package usecase

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/entity"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type migrateUserUseCase struct {
	db        *gorm.DB
	userRepo  entity.UserRepoInterface
	userUCase entity.UserUCaseInterface
}

func InitMigrateUserUCase(db *gorm.DB, uRepo entity.UserRepoInterface, uCase entity.UserUCaseInterface) entity.MigrateUserUCaseInterface {
	return &migrateUserUseCase{
		db:        db,
		userRepo:  uRepo,
		userUCase: uCase,
	}
}

func (mu *migrateUserUseCase) MigrateUserTable() {
	res := mu.db.Migrator().HasTable(&entity.User{})
	if res == false {
		log.Info("Create table users")
		_ = mu.db.Migrator().CreateTable(&entity.User{})

		hashPassword, _ := mu.userUCase.HashPassword("123123")
		initUser := entity.User{
			Id:       1,
			Name:     "Administrator",
			Username: "admin",
			Role:     "administrator",
			Email:    "admin@mail.com",
			Password: hashPassword,
		}

		mu.userRepo.StoreUser(&initUser)
		log.Info("User with role `administrator` has been inserted [admin:123123]. Password is so weak! please update/change password immediately")

	}
}
