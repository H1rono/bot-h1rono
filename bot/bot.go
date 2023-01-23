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

func New(botId string, userId string, accessToken string, verificationToken string) Bot {
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
