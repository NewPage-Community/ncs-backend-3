package service

import (
	"github.com/gunslinger23/kaiheila"
)

func (s *Service) IsAllChatEvent(event kaiheila.EventMsg) bool {
	return !event.Extra.Author.Bot &&
		event.TargetID == s.kaiheilaConfig.AllChatChannelID &&
		event.Type == kaiheila.MsgTypeText
}
