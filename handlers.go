package main

import (
	"context"
	"log"

	"github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
)

type Bot struct {
	client *traq.APIClient
	auth   context.Context
}

func NewBot(client *traq.APIClient, auth context.Context) Bot {
	return Bot{client, auth}
}

func (bot Bot) PingHandler(payload *traqbot.PingPayload) {
	log.Println("ping")
}

func (bot Bot) JoinHandler(payload *traqbot.JoinedPayload) {
	log.Println("=================================================")
	log.Println("チャンネルに参加しました。")
	log.Printf("チャンネル名: %s\n", payload.Channel.Name)
	log.Printf("チャンネルID: %s\n", payload.Channel.ID)
	log.Println("=================================================")
	msg := traq.NewPostMessageRequest(":oisu-1::oisu-2::oisu-3::oisu-4yoko:")
	bot.client.MessageApi.
		PostMessage(bot.auth, payload.Channel.ID).
		PostMessageRequest(*msg).
		Execute()
}

func (bot Bot) LeftHandler(payload *traqbot.LeftPayload) {
	log.Println("=================================================")
	log.Println("チャンネルから退出しました。")
	log.Printf("チャンネル名: %s\n", payload.Channel.Name)
	log.Printf("チャンネルID: %s\n", payload.Channel.ID)
	log.Println("=================================================")
	msg := traq.NewPostMessageRequest("byebye:8bit_sunglasses:")
	bot.client.MessageApi.
		PostMessage(bot.auth, payload.Channel.ID).
		PostMessageRequest(*msg).
		Execute()
}

func (bot Bot) MessageCreatedHandler(payload *traqbot.MessageCreatedPayload) {
	log.Println("=================================================")
	log.Printf("%sさんがメッセージを投稿しました。\n", payload.Message.User.DisplayName)
	log.Printf("メッセージID: %s\n", payload.Message.ID)
	log.Printf("内容: %s\n", payload.Message.Text)
	log.Println("=================================================")
	bot.client.MessageApi.
		AddMessageStamp(bot.auth, payload.Message.ID, "05f9b2b0-b2c7-4c48-8fd1-c68daaa23103").
		PostMessageStampRequest(*traq.NewPostMessageStampRequest(1)).
		Execute()
}
