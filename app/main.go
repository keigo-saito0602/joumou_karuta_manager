package main

import (
	"flag"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	echoSwagger "github.com/swaggo/echo-swagger"
	authHttpDelivery "go-clean-api/auth/delivery/http"
	_authMiddleware "go-clean-api/auth/middleware"
	_authUsecase "go-clean-api/auth/usecase"
	_ "go-clean-api/docs"
	userHttpDelivery "go-clean-api/user/delivery/http"
	_userMiddleware "go-clean-api/user/middleware"
	_userRepository "go-clean-api/user/repository"
	_userUsecase "go-clean-api/user/usecase"
	"go-clean-api/util"
	"time"
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
func main() {
	e := echo.New()

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	timeoutContext := time.Duration(30) * time.Second

	//connect database
	db, err := util.ConnectDb(config)
	if err != nil {
		panic("failed to connect database")
	}

	// run migrate
	runMigrate := flag.Bool("migrate", false, "Migrate True or false")
	flag.Parse()

	if *runMigrate {
		log.Info("Migrate Run")
		util.MigrateDbMysql(db)
	}

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
