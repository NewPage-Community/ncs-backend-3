package service

import (
	"backend/app/game/chat"
	serverService "backend/app/game/server/api/grpc/v1"
	"backend/pkg/log"
	"context"
	"fmt"
	"time"

	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/lonelyevil/khl"
)

func (s *Service) getServerStatus(event *khl.TextMessageContext) {
	// Build message
	var msg string
	resp, err := s.serverSrv.AllInfo(context.Background(), &serverService.AllInfoReq{A2S: true}, grpc_retry.WithPerRetryTimeout(10*time.Second))
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range resp.Info {
		if v.A2SInfo.MaxPlayers > 0 {
			msg += fmt.Sprintf("%s 🗺 `%s` ", v.ShortName, v.A2SInfo.Map)
			if v.A2SInfo.Players == 0 {
				msg += "🈳"
			} else if v.A2SInfo.Players < v.A2SInfo.MaxPlayers {
				msg += fmt.Sprintf("♿️ `%d/%d`", v.A2SInfo.Players, v.A2SInfo.MaxPlayers)
			} else {
				msg += fmt.Sprintf("🈵 `%d`", v.A2SInfo.Players)
			}
			msg += "\n"
		}
	}

	card := (&khl.CardMessageCard{Theme: khl.CardThemeInfo}).
		AddModule((&khl.CardMessageHeader{Text: khl.CardMessageElementText{
			Content: "服务器状态",
		}})).
		AddModule(&khl.CardMessageDivider{}).
		AddModule((&khl.CardMessageSection{}).SetText(&khl.CardMessageElementKMarkdown{
			Content: msg,
		})).
		AddModule(&khl.CardMessageDivider{}).
		AddModule((&khl.CardMessageSection{}).SetText(&khl.CardMessageElementKMarkdown{
			Content: fmt.Sprintf("🔍 [查看服务器详细](%s)", chat.ServerURL),
		}))

	// Send message
	_, err = s.kaiheilaClient.MessageCreate(&khl.MessageCreate{
		MessageCreateBase: khl.MessageCreateBase{
			Type:     khl.MessageTypeCard,
			TargetID: event.Common.TargetID,
			Content:  khl.CardMessage{card}.MustBuildMessage(),
			Quote:    event.Common.MsgID,
		},
		TempTargetID: event.Common.AuthorID,
	})
	log.CheckErr(err)
}
