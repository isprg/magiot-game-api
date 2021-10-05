package router

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	v1 "magiot-game-api/api/v1"
)

func NewRouter() *echo.Echo {
	e := echo.New() // Echo Instance

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// API Version 1.0
	e1 := e.Group("/api/v1")
	e1.GET("/info", v1.Info())
	e1.GET("/devices/:id", v1.GetDeviceFromID())
	e1.GET("/devices", v1.GetAllDevices())

	// Start Server
	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hi!")
}
