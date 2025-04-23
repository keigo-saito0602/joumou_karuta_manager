package router

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/interface/handler"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/keigo-saito0602/joumou_karuta_manager/docs"
)

func RegisterRoutes(e *echo.Echo, userHandler *handler.UserHandler) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// User
	e.GET("/users", userHandler.ListUsers)
	e.GET("/users/:id", userHandler.GetUser)
	e.POST("/users", userHandler.CreateUser)
	e.PUT("/users/:id", userHandler.UpdateUser)
	e.DELETE("/users/:id", userHandler.DeleteUser)
}
