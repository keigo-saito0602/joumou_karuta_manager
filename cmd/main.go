package main

import (
	"flag"
	authHttpDelivery "github.com/keigo-saito0602/JoumouKarutaTyping/auth/delivery/http"
	_authMiddleware "github.com/keigo-saito0602/JoumouKarutaTyping/auth/middleware"
	_authUsecase "github.com/keigo-saito0602/JoumouKarutaTyping/auth/usecase"
	_ "github.com/keigo-saito0602/JoumouKarutaTyping/docs"
	_migrateUsecase "github.com/keigo-saito0602/JoumouKarutaTyping/migrate/usecase"
	userHttpDelivery "github.com/keigo-saito0602/JoumouKarutaTyping/user/delivery/http"
	_userMiddleware "github.com/keigo-saito0602/JoumouKarutaTyping/user/middleware"
	_userRepository "github.com/keigo-saito0602/JoumouKarutaTyping/user/repository"
	_userUsecase "github.com/keigo-saito0602/JoumouKarutaTyping/user/usecase"
	"github.com/keigo-saito0602/JoumouKarutaTyping/util"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	"gorm.io/gorm"
)

var config *util.Config

func init() {
	c, _ := util.LoadConfig()
	config = &c
}

// @title Echo Simple Clean Api
// @version 0.0.1
// @description Simple echo rest api with clean architecture.

// @contact.name API Support
// @contact.url https://github.com/asdiyanarisha

// @BasePath /

// @Security BearerToken
// @securityDefinitions.apikey BearerToken
// @in header
// @name Authorization
func initEcho(db *gorm.DB) {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	timeoutContext := time.Duration(30) * time.Second

	authMiddle := _authMiddleware.InitAuthMiddleware(config)
	userMiddle := _userMiddleware.InitUserMiddleware()

	//initialize user repository and user usecase
	userRepo := _userRepository.NewUserRepository(db)
	userUc := _userUsecase.NewUserUseCase(userRepo, timeoutContext)

	//start handler user
	userHttpDelivery.RouteUserHandler(e, userUc, authMiddle.JwtConfigCustom(), *userMiddle)

	//initialize user auth use case
	authUc := _authUsecase.NewLoginUseCase(userRepo, config)

	//start handler user
	authHttpDelivery.RouteAuthHandler(e, authUc)

	e.Logger.Fatal(e.Start(config.AppPort))
}

func initMigrate(db *gorm.DB) {
	timeoutContext := time.Duration(30) * time.Second
	userRepo := _userRepository.NewUserRepository(db)
	userUc := _userUsecase.NewUserUseCase(userRepo, timeoutContext)

	migrateUserUc := _migrateUsecase.InitMigrateUserUCase(db, userRepo, userUc)
	migrateUserUc.MigrateUserTable()
}

func main() {
	// run migrate
	runMigrate := flag.Bool("migrate", false, "Migrate True or false")
	flag.Parse()

	//connect database
	db, err := util.ConnectDb(config)
	if err != nil {
		panic("failed to connect database")
	}

	switch {
	case *runMigrate:
		log.Info("Start Migrate Run")
		initMigrate(db)
	default:
		initEcho(db)
	}
}
