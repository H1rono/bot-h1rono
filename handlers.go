package main

import (
	"context"

	log "github.com/sirupsen/logrus"

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
	log.Info("ping")
}

func (bot Bot) JoinHandler(payload *traqbot.JoinedPayload) {
	log.Info("チャンネルに参加しました。")
	log.Infof("チャンネル名: %s\n", payload.Channel.Name)
	log.Infof("チャンネルID: %s\n", payload.Channel.ID)
	msg := traq.NewPostMessageRequest(":oisu-1::oisu-2::oisu-3::oisu-4yoko:")
	bot.client.MessageApi.
		PostMessage(bot.auth, payload.Channel.ID).
		PostMessageRequest(*msg).
		Execute()
}

func (bot Bot) LeftHandler(payload *traqbot.LeftPayload) {
	log.Info("チャンネルから退出しました。")
	log.Infof("チャンネル名: %s\n", payload.Channel.Name)
	log.Infof("チャンネルID: %s\n", payload.Channel.ID)
	msg := traq.NewPostMessageRequest("byebye:8bit_sunglasses:")
	bot.client.MessageApi.
		PostMessage(bot.auth, payload.Channel.ID).
		PostMessageRequest(*msg).
		Execute()
}

func (bot Bot) MessageCreatedHandler(payload *traqbot.MessageCreatedPayload) {
	log.Infof("%sさんがメッセージを投稿しました。\n", payload.Message.User.DisplayName)
	log.Infof("メッセージID: %s\n", payload.Message.ID)
	log.Infof("内容: %s\n", payload.Message.Text)
	// :eyes_chuukunn:を押す
	bot.client.MessageApi.
		AddMessageStamp(bot.auth, payload.Message.ID, "ca76e807-ca02-463a-bf97-4339bc5f305b").
		PostMessageStampRequest(*traq.NewPostMessageStampRequest(1)).
		Execute()
}
