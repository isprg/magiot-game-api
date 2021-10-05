package v1

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Device struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

func Info() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.String(http.StatusOK, "Magiot Game API Version 1.0")
	}
}

func GetAllDevices() echo.HandlerFunc {
	return func(c echo.Context) error {
		// id := c.Param("id")
		var devices = []Device{}

		// db

		return c.JSON(http.StatusOK, devices)
	}
}

func GetDeviceFromID() echo.HandlerFunc {
	return func(c echo.Context) error {
		var devices = []Device{}

		// db

		return c.JSON(http.StatusOK, devices)
	}
}
