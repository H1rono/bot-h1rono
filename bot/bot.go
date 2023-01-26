package bot

import (
	"context"

	"git.trap.jp/H1rono_K/bot-h1rono/util"

	"github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
)

type Bot struct {
	client            *traq.APIClient
	auth              context.Context
	stamps            util.Stamps
	Id                string
	UserId            string
	VerificationToken string
	AccessToken       string
	LogChannelId      string
}

func New(botId string, userId string, accessToken string, verificationToken string, logChannelId string) Bot {
	client := traq.NewAPIClient(traq.NewConfiguration())
	auth := context.WithValue(context.Background(), traq.ContextAccessToken, accessToken)
	stamps := util.FetchStamps(client, auth)
	return Bot{
		client,
		auth,
		stamps,
		botId,
		userId,
		verificationToken,
		accessToken,
		logChannelId,
	}
}

func (bot Bot) MakeHandlers() traqbot.EventHandlers {
	handlers := traqbot.EventHandlers{}
	handlers.SetPingHandler(bot.PingHandler)
	handlers.SetJoinedHandler(bot.JoinHandler)
	handlers.SetLeftHandler(bot.LeftHandler)
	handlers.SetMessageCreatedHandler(bot.MessageCreatedHandler)
	return handlers
}