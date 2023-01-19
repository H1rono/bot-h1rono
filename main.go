package main

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"

	traq "github.com/traPtitech/go-traq"
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

	at := os.Getenv("BOT_ACCESS_TOKEN")
	vt := os.Getenv("BOT_VERIFICATION_TOKEN")

	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, at)

	handlers := traqbot.EventHandlers{}
	bot := NewBot(client, auth)
	handlers.SetPingHandler(bot.PingHandler)
	handlers.SetJoinedHandler(bot.JoinHandler)
	handlers.SetLeftHandler(bot.LeftHandler)
	handlers.SetMessageCreatedHandler(bot.MessageCreatedHandler)

	server := traqbot.NewBotServer(vt, handlers)
	log.Fatal(server.ListenAndServe(":8080"))
}
