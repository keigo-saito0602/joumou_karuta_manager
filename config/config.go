package util

import (
	"fmt"
	"github.com/keigo-saito0602/JoumouKarutaTyping/docs"
	"io/ioutil"
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	AppPort      string `mapstructure:"PORT_API"`
	AppHost      string `mapstructure:"HOST"`
	HostDB       string `mapstructure:"HOST_DB"`
	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`
}

func LoadConfig() (config Config, err error) {
	// set path config
	viper.AddConfigPath(".")

	// set config file
	viper.SetConfigFile(".env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)

	// initialize swag config
	initSwag(&config)

	return
}

func initSwag(config *Config) {
	docs.SwaggerInfo.Host = fmt.Sprintf("%s%s", config.AppHost, config.AppPort)
	docs.SwaggerInfo.Version = loadVersion()
}

func loadVersion() string {
	content, err := ioutil.ReadFile("version.txt")

	if err != nil {
		log.Fatal(err)
	}

	return strings.Trim(string(content), "")
}
