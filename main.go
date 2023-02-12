package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	log "github.com/sirupsen/logrus"

	"github.com/H1rono/bot-h1rono/api"
	"github.com/H1rono/bot-h1rono/bot"
	"github.com/H1rono/bot-h1rono/util"
)

func main() {
	bid := os.Getenv("BOT_ID")
	uid := os.Getenv("BOT_USER_ID")
	at := os.Getenv("BOT_ACCESS_TOKEN")
	vt := os.Getenv("BOT_VERIFICATION_TOKEN")
	cid := os.Getenv("BOT_LOG_CHANNEL_ID")
	l := os.Getenv("LOG_LEVEL")

	rand.Seed(time.Now().Unix())
	util.SetupLogging(l)
	b := bot.New(bid, uid, at, vt, cid)
	e := echo.New()
	api.SetRouting(e, &b)
	log.Fatal(e.Start(":1323"))
}
