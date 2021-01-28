package service

import (
	pb "backend/app/bot/kaiheila/api/grpc/v1"
	chatService "backend/app/game/chat/api/grpc"
	"backend/pkg/log"
	"context"
	"github.com/gunslinger23/kaiheila"
)

const (
	KaiheilaServerID = -1
	AllChatChannelID = "4349043952859849"
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
	if !event.Extra.Author.Bot && event.TargetID == AllChatChannelID && event.Type == kaiheila.MsgTypeText {
		_, err := s.chat.AllChat(context.Background(), &chatService.AllChatReq{
			Name:     event.Extra.Author.Username + "#" + event.Extra.Author.IdentifyNum,
			Message:  event.Content,
			ServerId: KaiheilaServerID,
		})
		if err != nil {
			log.Error(err)
		}
	}
}
