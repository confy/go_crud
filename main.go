package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)



type Data struct {
	gorm.Model
	Value string `json:"value"`
  ID        uint           `gorm:"primaryKey"`
  CreatedAt time.Time
  UpdatedAt time.Time
  DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

  db.AutoMigrate(&Data{})

	// C
	e.POST("/data", func(c echo.Context) error {
    var data Data
    err := c.Bind(&data); if err != nil {
      return err
    }

		db.Create(&data)
    e.Logger.Info("Created ID: %s, value: %s", data.ID, data.Value)
		return c.JSON(http.StatusCreated, data)
	})
	// R
	e.GET("/data/:id", func(c echo.Context) error {
		id := c.Param("id")
		var data Data
		db.First(&data, id)
		e.Logger.Info("GET /data/%s", id)
		return c.JSON(http.StatusOK, data)
	})
	// U
	e.PUT("/data/:id", func(c echo.Context) error {
		id := c.Param("id")
    //convert id string to uint
    id_uint, err := strconv.ParseUint(id, 10, 64)
    if err != nil {
      return c.JSON(http.StatusBadRequest, "Invalid ID")
    }

		var data Data
    data.ID = uint(id_uint)
		db.First(&data, id)
		if err := c.Bind(&data); err != nil {
			return err
		}
		db.Save(&data)
		e.Logger.Info("Updated data: ", data)
		return c.JSON(http.StatusOK, data)
	})

	// D
	e.DELETE("/data/:id", func(c echo.Context) error {
		//with error handling
    		id := c.Param("id")
		var data Data
		db.First(&data, id)
		if data.ID == 0 {
			return echo.NewHTTPError(http.StatusNotFound, "Data not found")
		}
		db.Delete(&data)
		e.Logger.Info("Deleted data: ", data)
		return c.NoContent(http.StatusNoContent)
	})

	e.Logger.Print(e.Start(":1323"))
}
