package main

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/os-hun/go-echo-tutorial/route"
)

func main() {
	router := route.Init()
	router.Run(fasthttp.New(":1323"))
}
