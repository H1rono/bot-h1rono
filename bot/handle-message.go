package bot

import (
	"math/rand"
	"os"
	"reflect"
	"strings"

	"github.com/H1rono/bot-h1rono/util"
	log "github.com/sirupsen/logrus"

	"github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
)

const (
	MAX_STAMPS         = 500
	MAX_MESSAGE_LENGTH = 10000 // https://github.com/traPtitech/traQ_S-UI/blob/1daa36143945172c641943edfbd1412570e5b26b/src/lib/validate.ts#L10
)

func (bot *Bot) StampPatternMatch(message *traqbot.MessagePayload) {
	msg := message.PlainText
	patterns := util.ExtractStampPatterns(msg)
	if len(patterns) == 0 {
		return
	}
	rand.Shuffle(len(bot.stamps), func(i, j int) {
		bot.stamps[i], bot.stamps[j] = bot.stamps[j], bot.stamps[i]
	})
	result := make([]string, 0, MAX_STAMPS)
	lenStamps := 0
	lenMessage := 0
	for _, pattern := range patterns {
		stamps := util.PickStamps(pattern, bot.stamps)
		ls := 0
		for _, s := range stamps {
			ln := len(s)
			if lenMessage+ln > MAX_MESSAGE_LENGTH {
				break
			}
			ls++
			lenMessage += ln
		}
		// ls <= len(stamps)
		lenStamps += ls
		stamps = stamps[:ls]
		if lenStamps >= MAX_STAMPS {
			over := lenStamps - MAX_STAMPS
			stamps = stamps[:ls-over]
			result = append(result, stamps...)
			break
		}
		result = append(result, stamps...)
	}
	if reflect.DeepEqual(result, patterns) {
		return
	}
	res := strings.Join(result, "")
	log.Infof("パターン: %#v", patterns)
	log.Infof("結果: %#v", result)
	log.Infof("メッセージ長: %d", len(res))
	bot.SendMessage(message.ChannelID, res, false)
}

func (bot *Bot) HandleNormalMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleNormalMessageでメッセージを処理")
}

func (bot *Bot) HandleJoinMessage(message *traqbot.MessagePayload) {
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

func (bot *Bot) HandleLeaveMessage(message *traqbot.MessagePayload) {
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

func (bot *Bot) HandlePingMessage(message *traqbot.MessagePayload) {
	log.Trace("HandlePingMessageでメッセージを処理")
	bot.SendMessage(message.ChannelID, ":ping_pong:", false)
}

func (bot *Bot) HandleHelpMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleHelpMessageでメッセージを処理")
	readme, err := os.ReadFile("./README.md")
	if err != nil {
		log.Fatalf("[bot.HandleHelpMessage] %v", err)
	}
	bot.SendMessage(message.ChannelID, string(readme), false)
}

func (bot *Bot) HandleMentionMessage(message *traqbot.MessagePayload) {
	log.Trace("HandleMentionMessageでメッセージを処理")
}
