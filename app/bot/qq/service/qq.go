package service

import (
	pb "backend/app/bot/qq/api/grpc/v1"
	"backend/app/game/chat"
	chatService "backend/app/game/chat/api/grpc/v1"
	serverService "backend/app/game/server/api/grpc/v1"
	"backend/pkg/log"
	"context"
	"encoding/json"
	"fmt"
	qqMessage "github.com/miRemid/amy/message"
	qqModel "github.com/miRemid/amy/websocket/model"
)

func (s *Service) SendGroupMessage(ctx context.Context, req *pb.SendGroupMessageReq) (resp *pb.SendGroupMessageResp, err error) {
	resp = &pb.SendGroupMessageResp{}
	builder := qqMessage.NewCQMsgBuilder()
	for _, v := range s.qqGroups {
		_, err1 := s.qqAPIClient.SendGroupMsg(builder.GroupMsg(v, req.Message, req.AutoEscape), true)
		if err1 != nil {
			log.Error(err1)
		}
	}
	return
}

func (s *Service) OnMessage(event qqModel.CQEvent) {
	msg, ok := event.Map["raw_message"]
	if !ok {
		return
	}

	switch msg.(string) {
	case "状态":
		s.getServerStatus(event)
	}

	if msg.(string)[0] == '#' {
		_, err := s.chatSrv.AllChat(context.Background(), &chatService.AllChatReq{
			Name:     event.Map["sender"].(map[string]interface{})["nickname"].(string),
			Message:  msg.(string)[1:],
			ServerId: chat.QQID,
		})
		if err != nil {
			log.Error(err)
		}
	}
}

func (s *Service) getServerStatus(event qqModel.CQEvent) {
	// Group only
	if t := event.Map["message_type"].(string); t != "group" {
		return
	}
	// Get event
	groupID, err := event.Map["group_id"].(json.Number).Int64()
	if err != nil {
		return
	}

	// Build message
	var msg string
	resp, err := s.serverSrv.AllInfo(context.Background(), &serverService.AllInfoReq{A2S: true})
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range resp.Info {
		if v.A2SInfo.Id != 0 {
			msg += fmt.Sprintf("%s | %s (%d/%d)\n", v.ShortName, v.A2SInfo.Map, v.A2SInfo.Players, v.A2SInfo.MaxPlayers)
		}
	}
	msg += "仪表盘：https://game.new-page.xyz"

	// Send message
	builder := qqMessage.NewCQMsgBuilder()
	m := builder.GroupMsg(int(groupID), fmt.Sprintf("[CQ:reply,id=%s]%s", event.Map["message_id"], msg), false)
	_, err = s.qqAPIClient.SendGroupMsg(m, true)
	if err != nil {
		log.Error(err)
	}
}
