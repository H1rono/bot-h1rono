package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	traqbot "github.com/traPtitech/traq-bot"

	"git.trap.jp/H1rono_K/bot-h1rono/bot"
	"git.trap.jp/H1rono_K/bot-h1rono/util"
)

func main() {
	bid := os.Getenv("BOT_ID")
	uid := os.Getenv("BOT_USER_ID")
	at := os.Getenv("BOT_ACCESS_TOKEN")
	vt := os.Getenv("BOT_VERIFICATION_TOKEN")

	util.SetupLogging()

	b := bot.NewBot(bid, uid, at, vt)

	handlers := b.MakeHandlers()
	server := traqbot.NewBotServer(vt, handlers)
	log.Fatal(server.ListenAndServe(":8080"))
}
