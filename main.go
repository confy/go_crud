package main

import (
	"go_crud/routes"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/data", routes.PostData)
	e.GET("/data/:id", routes.GetData)
	e.PUT("/data/:id", routes.PutData)
	e.DELETE("/data/:id", routes.DeleteData)

	e.GET("/average", routes.GetAverage)

	e.GET("/hello", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello 🌎")
    })
	
	e.Logger.Fatal(e.Start(":8080"))
}
