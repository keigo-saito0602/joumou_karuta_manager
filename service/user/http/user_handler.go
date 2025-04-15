package http

import (
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/entity"
	_userMiddleware "github.com/keigo-saito0602/joumou_karuta_manager/user/middleware"
	"github.com/keigo-saito0602/joumou_karuta_manager/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type UserHandler struct {
	UserUc entity.UserUCaseInterface
}

func RouteUserHandler(
	e *echo.Echo,
	uscase entity.UserUCaseInterface,
	jwtConfig middleware.JWTConfig,
	userMiddle _userMiddleware.UserMiddleware) {

	handler := &UserHandler{
		UserUc: uscase,
	}
	userGroup := e.Group("/user/", middleware.JWTWithConfig(jwtConfig))

	userGroup.POST("add", handler.AddUser, userMiddle.OnlyAdmin)

	userGroup.GET("auth", handler.TestAuth, userMiddle.OnlyAdmin)
}

func (u *UserHandler) Index(c echo.Context) error {
	ctx := c.Request().Context()
	u.UserUc.Index(ctx)
	return c.String(http.StatusOK, "Hello, World!")
}

// AddUser godoc
// @ID user-insert
// @Accept json
// @description This endpoint use for add user
// @Produce json
// @Security BearerToken
// @Tags user
// @Param request body entity.User true "query add user"
// @Router /user/add [post]
func (u *UserHandler) AddUser(c echo.Context) (err error) {
	var user entity.User
	ctx := c.Request().Context()

	err = c.Bind(&user)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = u.UserUc.InsertUser(ctx, &user)
	if err != nil {
		resp := util.BuildResponseError(err)
		return c.JSON(http.StatusBadRequest, resp)
	}

	return c.JSON(http.StatusOK, user)
}

// TestAuth godoc
// @ID user-test-auth
// @description Testing bearer token is valid or not
// @Router /user/auth [get]
// @Security BearerToken
// @tags user
// @Success 200 {string} string "Authenticated
func (u *UserHandler) TestAuth(c echo.Context) error {
	return c.String(200, "Authenticated!")
}
