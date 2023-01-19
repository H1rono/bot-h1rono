package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	traqbot "github.com/traPtitech/traq-bot"
)

func main() {
	// TRACE, DEBUG, INFO, WARNING, ERROR, FATAL, PANIC
	l := os.Getenv("LOG_LEVEL")
	switch l {
	case "TRACE":
		log.SetLevel(log.TraceLevel)
	case "DEBUG":
		log.SetLevel(log.DebugLevel)
	case "INFO", "":
		// l = "" の場合に後の出力が残念になるため
		l = "INFO"
		log.SetLevel(log.InfoLevel)
	case "WARNING":
		log.SetLevel(log.WarnLevel)
	case "ERROR":
		log.SetLevel(log.ErrorLevel)
	case "FATAL":
		log.SetLevel(log.FatalLevel)
	case "PANIC":
		log.SetLevel(log.PanicLevel)
	default:
		log.Fatalf("Unexpected environment variable LOG_LEVEL=%s", l)
	}
	log.Infof("log level is set at %s", l)

	bid := os.Getenv("BOT_ID")
	uid := os.Getenv("BOT_USER_ID")
	at := os.Getenv("BOT_ACCESS_TOKEN")
	vt := os.Getenv("BOT_VERIFICATION_TOKEN")
	bot := NewBot(bid, uid, at, vt)

	handlers := traqbot.EventHandlers{}
	handlers.SetPingHandler(bot.PingHandler)
	handlers.SetJoinedHandler(bot.JoinHandler)
	handlers.SetLeftHandler(bot.LeftHandler)
	handlers.SetMessageCreatedHandler(bot.MessageCreatedHandler)

	server := traqbot.NewBotServer(vt, handlers)
	log.Fatal(server.ListenAndServe(":8080"))
}
