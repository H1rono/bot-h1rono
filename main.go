package main

import (
	"os"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"git.trap.jp/H1rono_K/bot-h1rono/bot"
	"git.trap.jp/H1rono_K/bot-h1rono/util"
)

func main() {
	bid := os.Getenv("BOT_ID")
	uid := os.Getenv("BOT_USER_ID")
	at := os.Getenv("BOT_ACCESS_TOKEN")
	vt := os.Getenv("BOT_VERIFICATION_TOKEN")
	l := os.Getenv("LOG_LEVEL")

	util.SetupLogging(l)

	b := bot.NewBot(bid, uid, at, vt)

	handlers := b.MakeHandlers()
	e := echo.New()
	e.POST("/bot", util.MakeBotHandler(vt, handlers))
	log.Fatal(e.Start(":1323"))
}
