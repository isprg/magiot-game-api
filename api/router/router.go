package router

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	v1 "magiot-game-api/api/v1"
)

func dumpHandler(c echo.Context, reqBody, resBody []byte) {
	fmt.Fprintf(os.Stdout, "Request: %+v\n", string(reqBody))
}

func NewRouter() *echo.Echo {
	e := echo.New() // Echo Instance

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}))

	// Routes
	e.GET("/", hello)

	// API Version 1.0
	e1 := e.Group("/api/v1")
	e1.GET("", v1.Info())

	e1.GET("/devices", v1.GetAllDevices())
	e1.POST("/devices", v1.PostDevice())
	e1.GET("/devices/:id", v1.GetDeviceFromID())
	e1.DELETE("/devices/:id", v1.DeleteDeviceFromID())

	// Start Server
	return e
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hi!")
}
