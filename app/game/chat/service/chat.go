package service

import (
	pb "backend/app/game/chat/api/grpc"
	server "backend/app/game/server/api/grpc"
	"backend/pkg/log"
	"context"
	"fmt"
)

const (
	ChatNotifyCMD = "ncs_chat_notify %d \"%s\" \"%s\""
	AllChatPrefix = "[全服聊天]"
	DefaultPrefix = "[系统提示]"
)

func (s *Service) AllChat(ctx context.Context, req *pb.AllChatReq) (resp *pb.AllChatResp, err error) {
	resp = &pb.AllChatResp{}

	_, err = s.ChatNotify(ctx, &pb.ChatNotifyReq{
		Uid:     0,
		Prefix:  AllChatPrefix,
		Message: req.Name + " : " + req.Message,
	})
	return
}

func (s *Service) ChatNotify(ctx context.Context, req *pb.ChatNotifyReq) (resp *pb.ChatNotifyResp, err error) {
	resp = &pb.ChatNotifyResp{}

	if len(req.Prefix) == 0 {
		req.Prefix = DefaultPrefix
	}

	res, err := s.server.RconAll(ctx, &server.RconAllReq{
		Cmd: fmt.Sprintf(ChatNotifyCMD, req.Uid, req.Prefix, req.Message),
	})
	if err == nil && res.Success == 0 {
		log.Warn("ChatNotify: no server sent successfully")
	}
	return
}
