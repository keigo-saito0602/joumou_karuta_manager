package middleware

import (
	"fmt"
	"github.com/keigo-saito0602/joumou_karuta_manager/domain/entity"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserMiddleware struct{}

func (u *UserMiddleware) OnlyAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role := c.Get("jwtInfo").(*entity.JwtCustomClaims).Role

		if role == "administrator" {
			return next(c)
		}

		return echo.NewHTTPError(http.StatusForbidden,
			fmt.Sprintf("Role %s not allowed to access this endpoint!", role))
	}
}

func InitUserMiddleware() *UserMiddleware {
	return &UserMiddleware{}
}
