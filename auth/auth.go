package auth

import (
	"oopLab1/utils"

	"github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.JWT(utils.GetJWTSecret())
}
