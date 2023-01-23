package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetRouting(e *echo.Echo) {
	api := e.Group("/api")
	api.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	api.GET("/echo", func(c echo.Context) error {
		header := fmt.Sprintf("%#v", c.Request().Header)
		return c.String(http.StatusOK, header)
	})
}
