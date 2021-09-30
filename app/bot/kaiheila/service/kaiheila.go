package service

import (
	pb "backend/app/bot/kaiheila/api/grpc/v1"
	"backend/app/game/chat"
	chatEvent "backend/app/game/chat/event"
	serverEvent "backend/app/game/server/event"
	donateEvent "backend/app/service/donate/event"
	"backend/pkg/log"
	"context"
	"fmt"
	"strconv"

	"github.com/lonelyevil/khl"
)

func (s *Service) SendChannelMsg(ctx context.Context, req *pb.SendMessageReq) (resp *pb.SendMessageResp, err error) {
	resp = &pb.SendMessageResp{}
	res, err := s.kaiheilaClient.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     khl.MessageType(req.Type),
			TargetID: req.ChannelId,
			Content:  req.Content,
			Quote:    req.Quote,
			Nonce:    req.Nonce,
		},
		TempTargetID: req.TempTargetId,
	})
	if err == nil {
		resp.Nonce = res.Nonce
		resp.MsgId = res.MsgID
		resp.MsgTimestamp = res.MegTimestamp.ToTime().Unix()
	}
	return
}

func (s *Service) EventHandler(event *khl.TextMessageContext) {
	// All chat event
	if !event.Extra.Author.Bot &&
		event.Common.TargetID == s.kaiheilaConfig.AllChatChannelID &&
		event.Common.Type == khl.MessageTypeText {
		log.CheckErr(s.dao.CreateAllChatEvent(context.Background(), &chatEvent.AllChatEventData{
			Name:       event.Extra.Author.Username + "#" + event.Extra.Author.IdentifyNum,
			Message:    event.Common.Content,
			ServerId:   chat.KaiheilaID,
			ServerName: chat.KaiheilaName,
		}))
	}
	if !event.Extra.Author.Bot &&
		event.Common.TargetID == s.kaiheilaConfig.BotChannelID &&
		event.Common.Type == khl.MessageTypeText {
		switch event.Common.Content {
		case "状态", "狀態", "/status":
			s.getServerStatus(event)
		}
	}
}

func (s *Service) AllChatEvent(ctx context.Context, data *chatEvent.AllChatEventData) {
	// Skip itself
	if data.ServerId == chat.KaiheilaID {
		return
	}

	card := (&khl.CardMessageCard{Theme: khl.CardThemeInfo}).
		AddModule((&khl.CardMessageSection{}).SetText(&khl.CardMessageElementKMarkdown{
			Content: fmt.Sprintf("📣 [%s](%s) `%s` :", data.ServerName, chat.GetUrl(int(data.ServerId)), chat.RemoveColor(data.Name)),
		})).
		AddModule((&khl.CardMessageSection{}).SetText(&khl.CardMessageElementText{
			Content: data.Message,
		}))

	_, err := s.kaiheilaClient.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     khl.MessageTypeCard,
			TargetID: s.kaiheilaConfig.AllChatChannelID,
			Content:  khl.CardMessage{card}.MustBuildMessage(),
		},
	})
	log.CheckErr(err)
}

func (s *Service) ChangeMapEvent(ctx context.Context, data *serverEvent.ChangeMapEventData) {
	card := (&khl.CardMessageCard{Theme: khl.CardThemeInfo}).
		AddModule((&khl.CardMessageHeader{Text: khl.CardMessageElementText{
			Content: "🗺 更换地图 🗺",
		}})).
		AddModule((&khl.CardMessageSection{}).SetText(&khl.CardMessageElementKMarkdown{
			Content: fmt.Sprintf("[%s](%s) `%s`", data.ServerName, chat.GetUrl(int(data.ServerId)), data.Map),
		}))

	_, err := s.kaiheilaClient.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     khl.MessageTypeCard,
			TargetID: s.kaiheilaConfig.BotChannelID,
			Content:  khl.CardMessage{card}.MustBuildMessage(),
		},
	})
	log.CheckErr(err)
}

func (s *Service) DonateEvent(ctx context.Context, data *donateEvent.DonateEventData) {
	steamUrl := "https://steamcommunity.com/profiles/" + strconv.FormatInt(data.SteamId, 10)
	card := (&khl.CardMessageCard{Theme: khl.CardThemeInfo}).
		AddModule((&khl.CardMessageHeader{Text: khl.CardMessageElementText{
			Content: "🙏 捐助信息 🙏",
		}})).
		AddModule((&khl.CardMessageSection{}).SetText(&khl.CardMessageElementKMarkdown{
			Content: fmt.Sprintf("感谢 [%s](%s) 捐助 **%d元**", data.Username, steamUrl, data.Amount),
		}))

	_, err := s.kaiheilaClient.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     khl.MessageTypeCard,
			TargetID: s.kaiheilaConfig.BotChannelID,
			Content:  khl.CardMessage{card}.MustBuildMessage(),
		},
	})
	log.CheckErr(err)
}
