package router

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/interface/handler"
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"

	_ "github.com/keigo-saito0602/joumou_karuta_manager/docs"
)

func RegisterRoutes(e *echo.Echo, h *handler.Handlers) {
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// User
	e.GET("/users", h.User.ListUsers)
	e.GET("/users/:id", h.User.GetUser)
	e.POST("/users", h.User.CreateUser)
	e.PUT("/users/:id", h.User.UpdateUser)
	e.DELETE("/users/:id", h.User.DeleteUser)

	// Memo
	e.GET("/memos", h.Memo.ListMemos)
	e.GET("/memos/:id", h.Memo.GetMemo)
	e.POST("/memos", h.Memo.CreateMemo)
	e.PUT("/memos/:id", h.Memo.UpdateMemo)
	e.DELETE("/memos/:id", h.Memo.DeleteMemo)
}
