package api

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"

	"git.trap.jp/H1rono_K/bot-h1rono/bot"
)

func SetRouting(e *echo.Echo, b *bot.Bot) {
	api := e.Group("/api")
	api.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "pong")
	})
	api.GET("/echo", func(c echo.Context) error {
		header := fmt.Sprintf("%#v", c.Request().Header)
		return c.String(http.StatusOK, header)
	})
	api.POST("/say", func(c echo.Context) error {
		return Say(c, b)
	})
}
