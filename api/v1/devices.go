package v1

import (
	"magiot-game-api/api/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Device struct {
	ID   int    `params:"id" json:"id"`
	Name string `params:"name" json:"name"`
}

func Info() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Magiot Game API Version 1.0")
	}
}

func GetAllDevices() echo.HandlerFunc {
	return func(c echo.Context) error {
		// db
		database := db.GetConnection()
		var deviceList []Device
		database.Find(&deviceList)

		return c.JSON(http.StatusOK, deviceList)
	}
}

func PostDevice() echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")

		// db
		database := db.GetConnection()
		database.Create(&Device{Name: name})

		var device Device
		database.Where("name = ?", name).Find(&device)
		return c.JSON(http.StatusCreated, device)
	}
}

func GetDeviceFromID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		// db
		database := db.GetConnection()
		var device Device
		database.Where("id = ?", id).Find(&device)

		return c.JSON(http.StatusOK, device)
	}
}

func DeleteDeviceFromID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		// db
		database := db.GetConnection()
		var device Device
		database.Where("id = ?", id).Delete(&device)

		return c.NoContent(http.StatusNoContent)
	}
}
