package service

import (
	pb "backend/app/game/chat/api/grpc/v1"
	event "backend/app/game/chat/event"
	server "backend/app/game/server/api/grpc/v1"
	donateEvent "backend/app/service/donate/event"
	"backend/pkg/log"
	"context"
	"fmt"
)

const (
	ChatNotifyCMD = "ncs_chat_notify %d \"%s\" \"%s\""
	AllChatPrefix = "[全服%s]"
	DefaultPrefix = "[系统提示]"
)

func (s *Service) AllChat(ctx context.Context, req *pb.AllChatReq) (resp *pb.AllChatResp, err error) {
	resp = &pb.AllChatResp{}

	chatServer, err := s.server.Info(ctx, &server.InfoReq{ServerId: req.ServerId})
	if err == nil {
		req.ServerName = chatServer.Info.ShortName
	}

	// Create event to message queue
	err = s.dao.CreateAllChatEvent(ctx, (*event.AllChatEventData)(req))

	return
}

// ChatNotify game server notify
func (s *Service) ChatNotify(ctx context.Context, req *pb.ChatNotifyReq) (resp *pb.ChatNotifyResp, err error) {
	resp = &pb.ChatNotifyResp{}

	if len(req.Prefix) == 0 {
		req.Prefix = DefaultPrefix
	}

	_, err = s.server.RconAll(ctx, &server.RconAllReq{
		Cmd: fmt.Sprintf(ChatNotifyCMD, req.Uid, req.Prefix, req.Message),
	})
	return
}

func (s *Service) AllChatEvent(ctx context.Context, data *event.AllChatEventData) {
	var prefix string
	if len(data.ServerName) > 0 {
		prefix = " · " + data.ServerName
	}
	_, err := s.ChatNotify(ctx, &pb.ChatNotifyReq{
		Uid:     0,
		Prefix:  fmt.Sprintf(AllChatPrefix, prefix),
		Message: data.Name + " : " + data.Message,
	})
	log.CheckErr(err)
}

func (s *Service) DonateEvent(ctx context.Context, data *donateEvent.DonateEventData) {
	_, err := s.ChatNotify(ctx, &pb.ChatNotifyReq{
		Uid:     0,
		Message: fmt.Sprintf("感谢 {green}%s{default} 捐助了 {green}%d元{default} !", data.Username, data.Amount),
	})
	log.CheckErr(err)
}
