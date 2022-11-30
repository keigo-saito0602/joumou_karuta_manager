package util

import (
	"go-clean-api/entity"
	"gorm.io/gorm"
)

func MigrateDbMysql(db *gorm.DB) {
	_ = db.AutoMigrate(&entity.User{})
}
