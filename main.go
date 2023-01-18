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
		fmt.Println("=================================================")
		fmt.Println("チャンネルに参加しました。")
		fmt.Printf("チャンネル名: %s\n", payload.Channel.Name)
		fmt.Printf("チャンネルID: %s\n", payload.Channel.ID)
		fmt.Println("=================================================")
	})
	handlers.SetMessageCreatedHandler(func(payload *traqbot.MessageCreatedPayload) {
		fmt.Println("=================================================")
		fmt.Printf("%sさんがメッセージを投稿しました。\n", payload.Message.User.DisplayName)
		fmt.Println("メッセージ：")
		fmt.Println(payload.Message.Text)
		fmt.Println("=================================================")
	})

	server := traqbot.NewBotServer(vt, handlers)
	log.Fatal(server.ListenAndServe(":8080"))
}
