package service

import (
	kaiheilaBot "backend/app/bot/kaiheila/api/grpc/v1"
	kaiheilaService "backend/app/bot/kaiheila/service"
	pb "backend/app/game/chat/api/grpc"
	server "backend/app/game/server/api/grpc"
	"backend/pkg/log"
	"context"
	"fmt"
)

const (
	ChatNotifyCMD   = "ncs_chat_notify %d \"%s\" \"%s\""
	AllChatPrefix   = "[全服%s]"
	DefaultPrefix   = "[系统提示]"
	KaiheilaName    = "开黑啦"
	KaiheilaURL     = "https://kaihei.co/p4Bl4i"
	KaiheilaMessage = "[{\"type\":\"card\",\"theme\":\"secondary\",\"size\":\"lg\",\"modules\":[{\"type\":\"section\",\"text\":{\"type\":\"kmarkdown\",\"content\":\":mega: [%s](%s) `%s` :\"}},{\"type\":\"section\",\"text\":{\"type\":\"plain-text\",\"content\":\"%s\"}}]}]"
	DiscordName     = "Discord"
	DiscordURL      = "https://discord.gg/SNxCf4Gv"
	DiscordMessage  = ""
	DiscordServerID = -100
	ServerURL       = "https://game.new-page.xyz/dashboard/home"
)

func (s *Service) AllChat(ctx context.Context, req *pb.AllChatReq) (resp *pb.AllChatResp, err error) {
	resp = &pb.AllChatResp{}
	serverShortName := ""

	sendToKaiheila := func(serverName, url string) {
		_, err = s.kaiheila.SendChannelMsg(context.Background(), &kaiheilaBot.SendMessageReq{
			Type:      10,
			ChannelId: kaiheilaService.AllChatChannelID,
			Content:   fmt.Sprintf(KaiheilaMessage, serverName, url, removeColor(removeInvalidChar(req.Name)), req.Message),
		})
		if err != nil {
			log.Error(err)
		}
	}

	sendToDiscord := func(serverName, url string) {
	}

	switch req.ServerId {
	// From Kaiheila
	case kaiheilaService.KaiheilaServerID:
		serverShortName = KaiheilaName
		go sendToDiscord(serverShortName, KaiheilaURL)
	// From Discord
	case DiscordServerID:
		serverShortName = DiscordName
		go sendToKaiheila(serverShortName, DiscordURL)
	// From game server
	default:
		chatServer, err := s.server.Info(ctx, &server.InfoReq{ServerId: req.ServerId})
		if err == nil {
			serverShortName = chatServer.Info.ShortName
		}
		go sendToKaiheila(serverShortName, ServerURL)
		go sendToDiscord(serverShortName, ServerURL)
	}

	// Always send to game server
	if len(serverShortName) > 0 {
		serverShortName = " · " + serverShortName
	}
	_, err = s.ChatNotify(ctx, &pb.ChatNotifyReq{
		Uid:     0,
		Prefix:  fmt.Sprintf(AllChatPrefix, serverShortName),
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
