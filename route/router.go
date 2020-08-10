package route

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go-echo-tutorial/api"
)

func Init() *echo.Echo {
	e := echo.New()
	
	// middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	
	// Routes
	v1 := e.Group("/api/v1")
	{
		v1.GET("/static", api.StaticMain())
	}
	
	e.Logger.Fatal(e.Start(":1323"))
	
	return e
}
