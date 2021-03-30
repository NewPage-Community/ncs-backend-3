package service

import (
	kaiheilaBot "backend/app/bot/kaiheila/api/grpc/v1"
	"backend/app/game/chat"
	pb "backend/app/game/chat/api/grpc"
	server "backend/app/game/server/api/grpc"
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
	serverShortName := ""

	// Non game notify function
	sendToKaiheila := func() {
		_, err = s.kaiheila.SendChannelMsg(context.Background(), &kaiheilaBot.SendMessageReq{
			Type:      10,
			ChannelId: chat.KaiheilaAllChatChannelID,
			Content:   fmt.Sprintf(chat.KaiheilaMessage, serverShortName, chat.ServerURL, removeColor(removeInvalidChar(req.Name)), req.Message),
		})
		if err != nil {
			log.Error(err)
		}
	}
	sendToDiscord := func() {
	}
	sendToQQ := func() {
		/*_, err = s.qq.SendGroupMessage(context.Background(), &qqBot.SendGroupMessageReq{
			Message:    fmt.Sprintf(chat.QQMessage, serverShortName, removeColor(removeInvalidChar(req.Name)), req.Message),
			AutoEscape: true,
		})
		if err != nil {
			log.Error(err)
		}*/
	}

	// Set non game server short name
	switch req.ServerId {
	case chat.KaiheilaID:
		serverShortName = chat.KaiheilaName
	case chat.DiscordID:
		serverShortName = chat.DiscordName
	case chat.QQID:
		serverShortName = chat.QQName
	default:
		chatServer, err := s.server.Info(ctx, &server.InfoReq{ServerId: req.ServerId})
		if err == nil {
			serverShortName = chatServer.Info.ShortName
		}
	}

	// Broadcast to non game server
	if req.ServerId > 0 {
		go sendToKaiheila()
		go sendToDiscord()
		go sendToQQ()
	}

	// Always send to game server
	var prefix string
	if len(serverShortName) > 0 {
		prefix = " · " + serverShortName
	}
	_, err = s.ChatNotify(ctx, &pb.ChatNotifyReq{
		Uid:     0,
		Prefix:  fmt.Sprintf(AllChatPrefix, prefix),
		Message: req.Name + " : " + req.Message,
	})
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
