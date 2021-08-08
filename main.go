package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"example.com/niceidea/Controllers"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/home", Controllers.IdeaListPage)

	e.Start(":1323")
}

