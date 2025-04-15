package middleware

import (
	"github.com/keigo-saito0602/JoumouKarutaTyping/domain/entity"
	"github.com/keigo-saito0602/JoumouKarutaTyping/util"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type AuthMiddleware struct {
	Config *util.Config
}

func (a *AuthMiddleware) JwtConfigCustom() middleware.JWTConfig {
	jwtConfig := middleware.JWTConfig{
		SigningKey:     []byte(a.Config.JwtSecretKey),
		Claims:         &entity.JwtCustomClaims{},
		SuccessHandler: a.JWTHandlerSuccess,
		ContextKey:     "userJwtInfo",
	}
	return jwtConfig
}

func (a *AuthMiddleware) JWTHandlerSuccess(c echo.Context) {
	c.Set("jwtInfo", c.Get("userJwtInfo").(*jwt.Token).Claims.(*entity.JwtCustomClaims))
}

func InitAuthMiddleware(config *util.Config) *AuthMiddleware {
	return &AuthMiddleware{
		Config: config,
	}
}
