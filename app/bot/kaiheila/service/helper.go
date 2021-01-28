package service

import "github.com/gunslinger23/kaiheila"

func IsAllChatEvent(event kaiheila.EventMsg) bool {
	return !event.Extra.Author.Bot && event.TargetID == AllChatChannelID && event.Type == kaiheila.MsgTypeText
}
