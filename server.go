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
		return c.String(http.StatusOK, "Hello, World")
	})
	e.GET("/show", show)
	e.GET("/users/:id", getUser)
	/*
	e.POST("/users", saveUser)
	e.PUT("/users/:id", updateUser)
	e.DELETE("/users/:id", deleteUser)
	*/
	
	// Server start
	e.Logger.Fatal(e.Start(":1323"))
}

// path parameters
// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// query parameters
// e.GET("/show", show)
func show(c echo.Context) error {
	// GEt team and member from the query string
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:" + team + ", member:" + member)
}
