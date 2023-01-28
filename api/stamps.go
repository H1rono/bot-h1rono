package api

import (
	"net/http"

	"github.com/H1rono/bot-h1rono/bot"
	"github.com/labstack/echo/v4"
)

func UpdateStamps(c echo.Context, b *bot.Bot) error {
	b.UpdateStamps()
	return c.NoContent(http.StatusNoContent)
}
