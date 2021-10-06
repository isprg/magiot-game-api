package v1

import (
	"magiot-game-api/api/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Device struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
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

// jsonに正しく変換されない問題
func PostDevice() echo.HandlerFunc {
	return func(c echo.Context) error {
		device := new(Device)
		if err := c.Bind(device); err != nil {
			return err
		}

		// db
		database := db.GetConnection()
		database.Create(&device)

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
