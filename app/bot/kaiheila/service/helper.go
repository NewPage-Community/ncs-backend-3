package service

import (
	"backend/app/game/chat"
	"github.com/gunslinger23/kaiheila"
)

func IsAllChatEvent(event kaiheila.EventMsg) bool {
	return !event.Extra.Author.Bot &&
		event.TargetID == chat.KaiheilaAllChatChannelID &&
		event.Type == kaiheila.MsgTypeText
}
