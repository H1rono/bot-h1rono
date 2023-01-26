package bot

import (
	"strings"

	"git.trap.jp/H1rono_K/bot-h1rono/util"
	log "github.com/sirupsen/logrus"

	"github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
)

func (bot Bot) StampPatternMatch(message *traqbot.MessagePayload) {
	msg := message.PlainText
	patterns := util.ExtractStampPatterns(msg)
	if len(patterns) == 0 {
		return
	}
	pat_s := strings.Join(patterns, "")
	result := make([]string, 0, len(patterns))
	for _, pattern := range patterns {
		stamps := util.FindAllStamps(pattern, bot.stamps)
		result = append(result, strings.Join(stamps, ""))
	}
	res := strings.Join(result, "")
	if pat_s == res {
		return
	}
	log.Infof("パターン: %#v", patterns)
	log.Infof("結果: %#v", result)
	bot.SendMessage(message.ChannelID, res)
}

func (bot Bot) HandleNormalMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleNormalMessageでメッセージを処理")
}

func (bot Bot) HandleJoinMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleJoinMessageでメッセージを処理")
	r, err := bot.client.BotApi.
		LetBotJoinChannel(bot.auth, bot.Id).
		PostBotActionJoinRequest(*traq.NewPostBotActionJoinRequest(message.ChannelID)).
		Execute()
	if err != nil {
		log.Error(err)
	}
	util.LogResponse(r)
}

func (bot Bot) HandleLeaveMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleLeaveMessageでメッセージを処理")
	r, err := bot.client.BotApi.
		LetBotLeaveChannel(bot.auth, bot.Id).
		PostBotActionLeaveRequest(*traq.NewPostBotActionLeaveRequest(message.ChannelID)).
		Execute()
	if err != nil {
		log.Error(err)
	}
	util.LogResponse(r)
}

func (bot Bot) HandlePingMessage(message *traqbot.MessagePayload) {
	log.Trace("HandlePingMessageでメッセージを処理")
	bot.SendMessage(message.ChannelID, "pong")
}

func (bot Bot) HandleMentionMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleMentionMessageでメッセージを処理")
}
