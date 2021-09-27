package service

import (
	pb "backend/app/bot/kaiheila/api/grpc/v1"
	"backend/app/game/chat"
	chatEvent "backend/app/game/chat/event"
	"backend/pkg/json"
	"backend/pkg/log"
	"context"
	"fmt"

	"github.com/gunslinger23/kaiheila"
)

const (
	KaiheilaMessage = "[{\"type\":\"card\",\"theme\":\"secondary\",\"size\":\"lg\",\"modules\":[{\"type\":\"section\",\"text\":{\"type\":\"kmarkdown\",\"content\":\":mega: [%s](%s) `%s` :\"}},{\"type\":\"section\",\"text\":{\"type\":\"plain-text\",\"content\":\"%s\"}}]}]"
	ServerURL       = "https://game.new-page.xyz"
)

func (s *Service) SendChannelMsg(ctx context.Context, req *pb.SendMessageReq) (resp *pb.SendMessageResp, err error) {
	resp = &pb.SendMessageResp{}
	res, err := s.kaiheilaClient.SendChannelMsg(kaiheila.SendMessageReq{
		Type:         int(req.Type),
		ChannelID:    req.ChannelId,
		Content:      req.Content,
		Quote:        req.Quote,
		Nonce:        req.Nonce,
		TempTargetID: req.TempTargetId,
	})
	if err == nil {
		resp.Nonce = res.Nonce
		resp.MsgId = res.MsgID
		resp.MsgTimestamp = res.MsgTimestamp
	}
	return
}

func (s *Service) EventHandler(event kaiheila.EventMsg) {
	raw, _ := json.Marshal(event)
	log.Info(string(raw))

	// All chat event
	if s.IsAllChatEvent(event) {
		log.CheckErr(s.dao.CreateAllChatEvent(context.Background(), &chatEvent.AllChatEventData{
			Name:       event.Extra.Author.Username + "#" + event.Extra.Author.IdentifyNum,
			Message:    event.Content,
			ServerId:   chat.KaiheilaID,
			ServerName: chat.KaiheilaName,
		}))
	}
}

func (s *Service) AllChatEvent(ctx context.Context, data *chatEvent.AllChatEventData) {
	_, err := s.SendChannelMsg(ctx, &pb.SendMessageReq{
		Type:      10,
		ChannelId: s.kaiheilaConfig.AllChatChannelID,
		Content:   fmt.Sprintf(KaiheilaMessage, data.ServerName, chat.GetUrl(int(data.ServerId)), removeColor(removeInvalidChar(data.Name)), data.Message),
	})
	log.CheckErr(err)
}
