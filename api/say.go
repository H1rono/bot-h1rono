package api

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/H1rono/bot-h1rono/bot"
)

type SayRequest struct {
	Content   string `json:"content"`
	ChannelId string `json:"channelId,omitempty"`
}

func Say(c echo.Context, b *bot.Bot) error {
	req := c.Request()
	if req.Header.Get("X-TRAQ-BOT-TOKEN") != b.VerificationToken {
		return c.NoContent(http.StatusForbidden)
	}
	if !strings.HasPrefix(req.Header.Get("Content-Type"), "application/json") {
		return c.NoContent(http.StatusBadRequest)
	}
	payload := &SayRequest{}
	if err := c.Bind(payload); err != nil {
		log.Error(err)
		return c.String(http.StatusInternalServerError, err.Error())
	}
	log.Infof("Send message `%#v`", payload)
	cid := payload.ChannelId
	if len(cid) == 0 {
		cid = b.LogChannelId
	}
	b.SendMessage(cid, payload.Content)
	return c.NoContent(http.StatusCreated)
}
