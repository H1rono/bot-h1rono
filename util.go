package main

import (
	"regexp"
	"strings"

	traqbot "github.com/traPtitech/traq-bot"
)

type MessageType uint8

const (
	MESSAGE_NORMAL    MessageType = iota // 参加しているチャンネルでの普通のメッセージ
	MESSAGE_JOIN                         // "@BOT_H1rono :oisu-:"みたいな
	MESSAGE_LEAVE                        // "@BOT_H1rono :wave:"みたいな
	MESSAGE_PING                         // `(@BOT_H1rono )?ping`
	MESSAGE_MENTIONED                    // "@BOT_H1rono"を含む何か
)

var (
	JOIN_REGEXP  = regexp.MustCompile(`^\s*@bot_h1rono\s+:oisu-(1::oisu-2::oisu-3::oisu-4yoko)?:\s*$`)
	LEAVE_REGEXP = regexp.MustCompile(`^\s*@bot_h1rono\s+:wave:\s*$`)
	PING_REGEXP  = regexp.MustCompile(`^\s*(@bot_h1rono\s)?\s*:[A-Za-z_]*ping[A-Za-z_]*:\s*$`)
)

func (bot Bot) JudgeMessageType(message *traqbot.MessagePayload) MessageType {
	b := []byte(strings.ToLower(message.PlainText))
	if JOIN_REGEXP.Match(b) {
		return MESSAGE_JOIN
	}
	if LEAVE_REGEXP.Match(b) {
		return MESSAGE_LEAVE
	}
	if PING_REGEXP.Match(b) {
		return MESSAGE_PING
	}
	for _, e := range message.Embedded {
		if e.ID == bot.userId {
			return MESSAGE_MENTIONED
		}
	}
	return MESSAGE_NORMAL
}
