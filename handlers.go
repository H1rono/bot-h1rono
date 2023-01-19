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

/* --- ここからシステム系のイベントたち --- */

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
	msg := traq.NewPostMessageRequest(":leave:d")
	bot.client.MessageApi.
		PostMessage(bot.auth, payload.Channel.ID).
		PostMessageRequest(*msg).
		Execute()
}

/* --- ここまでシステム系のイベントたち --- */

/* --- ここからメッセージ系のイベントたち --- */

func (bot Bot) MessageCreatedHandler(payload *traqbot.MessageCreatedPayload) {
	log.Infof("%sさんがメッセージを投稿しました。\n", payload.Message.User.DisplayName)
	log.Infof("メッセージID: %s\n", payload.Message.ID)
	log.Infof("内容: %s\n", payload.Message.Text)
	log.Infof("埋め込み: %v\n", payload.Message.Embedded)
	// :kidoku:を押す
	bot.client.MessageApi.
		AddMessageStamp(bot.auth, payload.Message.ID, "aa9d4808-de1a-400d-9ab2-6db66fcd5bc7").
		PostMessageStampRequest(*traq.NewPostMessageStampRequest(1)).
		Execute()
}

/* --- ここまでメッセージ系のイベントたち --- */