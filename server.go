package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// echo instance
	e := echo.New()
	
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Routes
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	
	// Server start
	e.Logger.Fatal(e.Start(":1323"))
}
