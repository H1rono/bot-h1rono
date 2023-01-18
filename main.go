package main

import (
	"log"
	"os"

	traqbot "github.com/traPtitech/traq-bot"
)

func main() {
	vt := os.Getenv("BOT_VERIFICATION_TOKEN")

	handlers := traqbot.EventHandlers{}
	handlers.SetJoinedHandler(func(payload *traqbot.JoinedPayload) {
		log.Println("=================================================")
		log.Println("チャンネルに参加しました。")
		log.Printf("チャンネル名: %s\n", payload.Channel.Name)
		log.Printf("チャンネルID: %s\n", payload.Channel.ID)
		log.Println("=================================================")
	})
	handlers.SetMessageCreatedHandler(func(payload *traqbot.MessageCreatedPayload) {
		log.Println("=================================================")
		log.Printf("%sさんがメッセージを投稿しました。\n", payload.Message.User.DisplayName)
		log.Println("メッセージ：")
		log.Println(payload.Message.Text)
		log.Println("=================================================")
	})

	server := traqbot.NewBotServer(vt, handlers)
	log.Fatal(server.ListenAndServe(":8080"))
}
