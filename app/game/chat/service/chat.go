package service

import (
	pb "backend/app/game/chat/api/grpc"
	server "backend/app/game/server/api/grpc"
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
	serverShortName := ""

	chatServer, err := s.server.Info(ctx, &server.InfoReq{ServerId: req.ServerId})
	if err == nil {
		serverShortName = chatServer.Info.ShortName
		serverShortName = " · " + serverShortName
	}

	_, err = s.ChatNotify(ctx, &pb.ChatNotifyReq{
		Uid:     0,
		Prefix:  fmt.Sprintf(AllChatPrefix, serverShortName),
		Message: req.Name + " : " + req.Message,
	})
	return
}

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
