package v1

import (
	"fmt"
	"magiot-game-api/api/db"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Device struct {
	ID     string `gorm:"primaryKey" json:"id"`
	Name   string `json:"name"`
	Target string `json:"target"`
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
		id := c.Request().Header.Get("id")
		name := c.Request().Header.Get("name")
		if len(name) == 0 || len(id) == 0 {
			apierr := APIError{
				Code:    100,
				Message: "invalid request",
			}
			return c.JSON(http.StatusBadRequest, apierr)
		}

		device := Device{
			ID:   id,
			Name: name,
		}

		// debug
		for key, values := range c.Request().Header {
			fmt.Println(key)
			for _, value := range values {
				fmt.Println(value)
			}
		}

		// db
		database := db.GetConnection()
		database.Create(&device)
		return c.JSON(http.StatusCreated, device)
	}
}

func UpdateDeviceStatus() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		newDevice := new(Device)
		newDevice.Target = c.Request().Header.Get("target")

		// db
		database := db.GetConnection()
		var device Device
		database.First(&device, id).Updates(newDevice)

		return c.JSON(http.StatusOK, device)
	}
}

func GetDeviceFromID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		// db
		database := db.GetConnection()
		var device Device
		database.First(&device, id)

		return c.JSON(http.StatusOK, device)
	}
}

func DeleteDeviceFromID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")

		// db
		database := db.GetConnection()
		var device Device
		database.First(&device, id)
		database.Delete(device)

		return c.NoContent(http.StatusNoContent)
	}
}
