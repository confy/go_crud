package routes

import (
	"go_crud/database"
	"go_crud/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

var db = database.Connect()

func PostData(c echo.Context) error {
	var data models.Data
	err := c.Bind(&data)
	if err != nil {
		return err
	}

	db.Create(&data)
	return c.JSON(http.StatusCreated, data)
}

func GetData(c echo.Context) error {
	id := c.Param("id")
	var data models.Data
	db.First(&data, id)
	if data.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Data not found")
	}
	return c.JSON(http.StatusOK, data)
}

func PutData(c echo.Context) error {
	id := c.Param("id")
	id_uint, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid ID")
	}

	var data models.Data
	data.ID = uint(id_uint)
	db.First(&data, id)
	if err := c.Bind(&data); err != nil {
		return err
	}
	db.Save(&data)
	return c.JSON(http.StatusOK, data)
}

func DeleteData(c echo.Context) error {
	//with error handling
	id := c.Param("id")
	var data models.Data
	db.First(&data, id)
	if data.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Data not found")
	}
	db.Delete(&data)
	return c.NoContent(http.StatusNoContent)
}

func GetAverage(c echo.Context) error {
  items := []models.Data{}
  db.Find(&items)
  var sum uint
  for _, item := range items {
    sum += item.Value
  }
  length := uint(len(items))
  average := sum / length
  res := map[string]uint{"average": average, "length": length}
  return c.JSON(http.StatusOK, res)

}
