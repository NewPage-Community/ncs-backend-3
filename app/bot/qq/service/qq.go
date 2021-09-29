package service

import (
	pb "backend/app/bot/qq/api/grpc/v1"
	"backend/app/game/chat"
	chatEvent "backend/app/game/chat/event"
	donateEvent "backend/app/service/donate/event"

	serverEvent "backend/app/game/server/event"
	"backend/pkg/log"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	qqMessage "github.com/miRemid/amy/message"
	qqModel "github.com/miRemid/amy/websocket/model"
)

const (
	QQMessage    = "[📣%s] %s:\n%s"
	ServerURL    = "https://game.new-page.xyz"
	SuccessReply = "✅成功处理✅"
)

func (s *Service) SendGroupMessage(ctx context.Context, req *pb.SendGroupMessageReq) (resp *pb.SendGroupMessageResp, err error) {
	resp = &pb.SendGroupMessageResp{}
	builder := qqMessage.NewCQMsgBuilder()
	for _, v := range s.qqConfig.QQGroups {
		_, err1 := s.qqAPIClient.SendGroupMsg(builder.GroupMsg(int(v), req.Message, req.AutoEscape), true)
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

	// Group only
	if t := event.Map["message_type"].(string); t != "group" {
		return
	}

	// Get event
	groupID, err := event.Map["group_id"].(json.Number).Int64()
	if err != nil {
		return
	}

	// Service auth group only
	if !s.InGroup(groupID) && !s.IsAdminGroup(groupID) {
		return
	}

	if msg.(string)[0] == '#' {
		log.CheckErr(s.dao.CreateAllChatEvent(context.Background(), &chatEvent.AllChatEventData{
			Name:       event.Map["sender"].(map[string]interface{})["nickname"].(string),
			Message:    msg.(string)[1:],
			ServerId:   chat.QQID,
			ServerName: chat.QQName,
		}))
	}

	cmd := strings.Split(msg.(string), " ")
	switch cmd[0] {
	case "状态", "狀態", "/status":
		s.getServerStatus(event, groupID)
	}

	if s.IsAdminGroup(groupID) {
		switch cmd[0] {
		case "/ban":
			s.banPlayer(event, cmd[1:])
		case "/unban":
			s.unBanPlayer(event, cmd[1:])
		case "/map":
			s.changeMap(event, cmd[1:])
		case "/donate":
			s.donate(event, cmd[1:])
		}
	}
}

func (s *Service) Reply(event qqModel.CQEvent, message string) (err error) {
	group, _ := event.Map["group_id"].(json.Number).Int64()
	builder := qqMessage.NewCQMsgBuilder()
	m := builder.GroupMsg(int(group), fmt.Sprintf("[CQ:reply,id=%s]%s", event.Map["message_id"], message), false)
	_, err = s.qqAPIClient.SendGroupMsg(m, true)
	return
}

func (s *Service) AllChatEvent(ctx context.Context, data *chatEvent.AllChatEventData) {
	// Skip itself
	if data.ServerId == chat.QQID {
		return
	}

	_, err := s.SendGroupMessage(ctx, &pb.SendGroupMessageReq{
		Message:    fmt.Sprintf("📣%s📣\n%s : %s", data.ServerName, data.Name, data.Message),
		AutoEscape: false,
	})
	log.CheckErr(err)
}

func (s *Service) ChangeMapEvent(ctx context.Context, data *serverEvent.ChangeMapEventData) {
	_, err := s.SendGroupMessage(ctx, &pb.SendGroupMessageReq{
		Message:    fmt.Sprintf("🗺%s🗺\n更换地图 %s", data.ServerName, data.Map),
		AutoEscape: false,
	})
	log.CheckErr(err)
}

func (s *Service) DonateEvent(ctx context.Context, data *donateEvent.DonateEventData) {
	_, err := s.SendGroupMessage(ctx, &pb.SendGroupMessageReq{
		Message:    fmt.Sprintf("🙏捐助信息🙏\n感谢 %s 捐助了 %d元", data.Username, data.Amount),
		AutoEscape: false,
	})
	log.CheckErr(err)
}

func (s *Service) InGroup(id int64) bool {
	for _, v := range s.qqConfig.QQGroups {
		if v == id {
			return true
		}
	}
	return false
}

func (s *Service) IsAdminGroup(id int64) bool {
	for _, v := range s.qqConfig.AdminGroups {
		if v == id {
			return true
		}
	}
	return false
}

func (s *Service) IsAdmin(id int64) bool {
	for _, v := range s.qqConfig.Admins {
		if v == id {
			return true
		}
	}
	return false
}
