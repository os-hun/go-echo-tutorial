package route

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func Init() *echo.Echo {
	e := echo.New()
	
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/static", func(c echo.Context) error {
			return c.String(http.StatusOK, "Hello, World")
		})
	}
	
	e.Logger.Fatal(e.Start(":1323"))
	
	return e
}
