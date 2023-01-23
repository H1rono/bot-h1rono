package bot

import (
	log "github.com/sirupsen/logrus"
	"github.com/traPtitech/go-traq"
	traqbot "github.com/traPtitech/traq-bot"
)

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
	m := ":leaves:"
	bot.SendMessage(payload.Channel.ID, m)
}

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
