package middlewares

import (
	"ApiSup/pkg/hashing"
	"ApiSup/pkg/mapear/response"
	"net/http"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func VerifyTokenHandler() echo.MiddlewareFunc {
	handler := echojwt.JWT([]byte(hashing.Key))

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if err := handler(next)(c); err != nil {
				return c.JSON(http.StatusUnauthorized, response.Error{Message: "Acesso não autorizado!", Description: err})
			}
			return next(c)
		}
	}
}
