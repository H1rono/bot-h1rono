package bot

import (
	"io/ioutil"
	"math/rand"
	"reflect"
	"strings"

	"git.trap.jp/H1rono_K/bot-h1rono/util"
	log "github.com/sirupsen/logrus"

	"github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
)

const MAX_STAMPS = 1000

func (bot Bot) StampPatternMatch(message *traqbot.MessagePayload) {
	msg := message.PlainText
	patterns := util.ExtractStampPatterns(msg)
	if len(patterns) == 0 {
		return
	}
	result := make([]string, 0, MAX_STAMPS)
	lenStamps := 0
	for _, pattern := range patterns {
		stamps := util.FindAllStamps(pattern, bot.stamps)
		lenStamps += len(stamps)
		if lenStamps >= MAX_STAMPS {
			over := lenStamps - MAX_STAMPS
			stamps = stamps[:len(stamps)-over]
			result = append(result, stamps...)
			break
		}
		result = append(result, stamps...)
	}
	if reflect.DeepEqual(result, patterns) {
		return
	}
	rand.Shuffle(len(result), func(i, j int) {
		result[i], result[j] = result[j], result[i]
	})
	res := strings.Join(result, "")
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
	bot.SendMessage(message.ChannelID, ":ping_pong:")
}

func (bot Bot) HandleHelpMessage(message *traqbot.MessagePayload) {
	readme, err := ioutil.ReadFile("./README.md")
	if err != nil {
		log.Fatalf("[bot.HandleHelpMessage] %v", err)
	}
	bot.SendMessage(message.ChannelID, string(readme))
}

func (bot Bot) HandleMentionMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleMentionMessageでメッセージを処理")
}
