package api

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func StaticMain() echo.HandlerFunc {
	return func(c echo.Context) (err error) {
		return c.JSON(http.StatusOK, "Hello, World!")
	}
}
