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

  data_route := e.Group("/data")
  data_route.POST("", routes.Post)
  data_route.GET("/:id", routes.Get)
  data_route.PUT("/:id", routes.Put)
  data_route.DELETE("/:id", routes.Delete)

  e.Logger.Fatal(e.Start(":8080"))
}
