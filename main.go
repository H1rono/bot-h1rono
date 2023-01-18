package main

import (
	"context"
	"fmt"
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

	v, _, _ := client.ChannelApi.
		GetChannels(auth).
		IncludeDm(true).
		Execute()
	fmt.Printf("%#v", v)

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
