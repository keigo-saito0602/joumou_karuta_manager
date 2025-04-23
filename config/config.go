package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort string
	DBUser  string
	DBPass  string
	DBHost  string
	DBPort  string
	DBName  string
}

var AppConfig Config

func LoadConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	AppConfig = Config{
		AppPort: viper.GetString("app.port"),
		DBUser:  viper.GetString("db.user"),
		DBPass:  viper.GetString("db.pass"),
		DBHost:  viper.GetString("db.host"),
		DBPort:  viper.GetString("db.port"),
		DBName:  viper.GetString("db.name"),
	}
}
