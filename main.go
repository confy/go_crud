package main

import (
	"go_crud/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

  e.POST("/data", routes.Post)
  e.GET("/data/:id", routes.Get)
  e.PUT("/data/:id", routes.Put)
  e.DELETE("/data/:id", routes.Delete)

  e.Logger.Fatal(e.Start(":8080"))
}
