package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ConnectDb(config *Config) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(config.HostDB), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	return db, err
}
