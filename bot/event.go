package bot

import (
	"encoding/json"

	log "github.com/sirupsen/logrus"
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
	j, err := json.Marshal(payload)
	if err != nil {
		log.Error(err)
	} else {
		log.Info(string(j))
	}
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
	case MESSAGE_HELP:
		bot.HandleHelpMessage(m)
	case MESSAGE_FROM_BOT:
		// BOTの発言には反応しない
		return
	}
	bot.StampPatternMatch(&payload.Message)
}

// STAMP_CREATED
func (bot Bot) StampCreated(payload *traqbot.StampCreatedPayload) {
	log.Info("スタンプが作成されました。")
	j, err := json.Marshal(payload)
	if err != nil {
		log.Errorf("Error while json.Marshal: %v", err)
	} else {
		log.Infof("イベントペイロード: %s", j)
	}
}
