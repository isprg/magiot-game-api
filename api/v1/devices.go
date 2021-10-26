package v1

import (
	"fmt"
	"magiot-game-api/api/db"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type Device struct {
	ID   int    `gorm:"primary_key" param:"id" json:"id"`
	Name string `param:"name" json:"name"`
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
		name := c.Request().Header.Get("name")
		if len(name) == 0 {
			apierr := APIError{
				Code:    100,
				Message: "invalid request - header is empty",
			}
			return c.JSON(http.StatusBadRequest, apierr)
		}

		device := Device{
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

func GetDeviceFromID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		// db
		database := db.GetConnection()
		var device Device
		database.Where("id = ?", id).Find(&device)

		return c.JSON(http.StatusOK, device)
	}
}

func DeleteDeviceFromID() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, _ := strconv.Atoi(c.Param("id"))

		// db
		database := db.GetConnection()
		var device Device
		database.Where("id = ?", id).Delete(&device)

		return c.NoContent(http.StatusNoContent)
	}
}
