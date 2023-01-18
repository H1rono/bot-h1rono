package main

import (
	"context"
	"log"
	"os"

	traq "github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
)

func main() {
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
