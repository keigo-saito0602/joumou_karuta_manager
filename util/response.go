package util

import "github.com/labstack/echo/v4"

func ErrorJSON(c echo.Context, code int, message string) error {
	return c.JSON(code, map[string]string{"error": message})
}
