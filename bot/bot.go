package bot

import (
	"context"

	log "github.com/sirupsen/logrus"

	"github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
)

type Bot struct {
	client      *traq.APIClient
	auth        context.Context
	id          string
	userId      string
	verifToken  string
	accessToken string
}

func NewBot(botId string, userId string, accessToken string, verificationToken string) Bot {
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, accessToken)
	return Bot{client, auth, botId, userId, verificationToken, accessToken}
}

func (bot Bot) MakeHandlers() traqbot.EventHandlers {
	handlers := traqbot.EventHandlers{}
	handlers.SetPingHandler(bot.PingHandler)
	handlers.SetJoinedHandler(bot.JoinHandler)
	handlers.SetLeftHandler(bot.LeftHandler)
	handlers.SetMessageCreatedHandler(bot.MessageCreatedHandler)
	return handlers
}

/* --- ここからシステム系のイベントたち --- */

// PING
func (bot Bot) PingHandler(payload *traqbot.PingPayload) {
	log.Info("ping")
}

// JOIN
func (bot Bot) JoinHandler(payload *traqbot.JoinedPayload) {
	log.Info("チャンネルに参加しました。")
	log.Infof("チャンネル名: %s", payload.Channel.Name)
	log.Infof("チャンネルID: %s", payload.Channel.ID)
	m := ":oisu-1::oisu-2::oisu-3::oisu-4yoko:"
	bot.SendMessage(payload.Channel.ID, m)
}

// LEFT
func (bot Bot) LeftHandler(payload *traqbot.LeftPayload) {
	log.Info("チャンネルから退出しました。")
	log.Infof("チャンネル名: %s", payload.Channel.Name)
	log.Infof("チャンネルID: %s", payload.Channel.ID)
	m := ":leave:d"
	bot.SendMessage(payload.Channel.ID, m)
}

/* --- ここまでシステム系のイベントたち --- */

/* --- ここからメッセージ系のイベントたち --- */

// MESSAGE_CREATED
func (bot Bot) MessageCreatedHandler(payload *traqbot.MessageCreatedPayload) {
	m := &payload.Message
	log.Info("メッセージが投稿されました。")
	log.Infof("投稿者: 名前:%s, traQ ID:%s, UUID:%s", m.User.DisplayName, m.User.Name, m.User.ID)
	log.Infof("メッセージID: %s", m.ID)
	log.Infof("内容: %s", m.PlainText)
	log.Infof("埋め込み: %v", m.Embedded)
	t := bot.JudgeMessageType(m)
	switch t {
	case MESSAGE_NORMAL:
		bot.HandleNormalMessage(m)
	case MESSAGE_JOIN:
		bot.HandleJoinMessage(m)
	case MESSAGE_LEAVE:
		bot.HandleLeaveMessage(m)
	case MESSAGE_PING:
		bot.HandlePingMessage(m)
	case MESSAGE_MENTIONED:
		bot.HandleMentionMessage(m)
	case MESSAGE_FROM_BOT:
		// BOTの発言には反応しない
		return
	}
	// :kidoku:を押す
	bot.client.MessageApi.
		AddMessageStamp(bot.auth, payload.Message.ID, "aa9d4808-de1a-400d-9ab2-6db66fcd5bc7").
		PostMessageStampRequest(*traq.NewPostMessageStampRequest(1)).
		Execute()
}

func (bot Bot) HandleNormalMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleNormalMessageでメッセージを処理")
}

func (bot Bot) HandleJoinMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleJoinMessageでメッセージを処理")
	bot.client.BotApi.
		LetBotJoinChannel(bot.auth, bot.id).
		PostBotActionJoinRequest(*traq.NewPostBotActionJoinRequest(message.ChannelID)).
		Execute()
}

func (bot Bot) HandleLeaveMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleLeaveMessageでメッセージを処理")
	bot.client.BotApi.
		LetBotLeaveChannel(bot.auth, bot.id).
		PostBotActionLeaveRequest(*traq.NewPostBotActionLeaveRequest(message.ChannelID)).
		Execute()
}

func (bot Bot) HandlePingMessage(message *traqbot.MessagePayload) {
	log.Trace("HandlePingMessageでメッセージを処理")
	bot.SendMessage(message.ChannelID, "pong")
}

func (bot Bot) HandleMentionMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleMentionMessageでメッセージを処理")
}

/* --- ここまでメッセージ系のイベントたち --- */
