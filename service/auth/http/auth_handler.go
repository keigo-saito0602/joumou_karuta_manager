package http

import (
	"github.com/keigo-saito0602/JoumouKarutaTyping/domain/entity"
	"github.com/keigo-saito0602/JoumouKarutaTyping/util"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

type AuthHandler struct {
	AuthUc entity.AuthUCaseInterface
}

func RouteAuthHandler(c *echo.Echo, uscase entity.AuthUCaseInterface) {
	handler := &AuthHandler{
		AuthUc: uscase,
	}
	c.POST("/login", handler.Login)
}

// Login godoc
// @ID process-login
// @Accept json
// @description Endpoint for input username and password
// @Produce json
// @Tags auth
// @Param request body entity.Login true "query login"
// @Router /login [post]
// @Success 200 {object} entity.ResponseLogin
// @failure 400 {object} entity.ResponseError
// @failure 401 {object} entity.ResponseError
func (a *AuthHandler) Login(c echo.Context) (err error) {
	var login entity.Login
	var bearer string

	err = c.Bind(&login)
	if err != nil {
		return c.String(http.StatusBadRequest, err.Error())
	}

	err = a.AuthUc.Login(&login, &bearer)
	if err != nil {
		resp := util.BuildResponseError(err)
		return c.JSON(http.StatusBadRequest, resp)
	}

	if bearer == "" {
		msg := "Unauthorized"
		return c.JSON(http.StatusUnauthorized, entity.ResponseError{Message: &msg})
	}
	log.Print("User `", login.Username, "` has login")

	return c.JSON(http.StatusOK, entity.ResponseLogin{Bearer: bearer})
}
