package service

import (
	donateService "backend/app/service/donate/api/grpc/v1"
	"backend/pkg/log"
	"context"
	"strconv"

	qqModel "github.com/miRemid/amy/websocket/model"
)

func (s *Service) donate(event qqModel.CQEvent, cmd []string) {
	ctx := context.Background()
	steamid, _ := strconv.ParseInt(cmd[0], 0, 0)
	amount, _ := strconv.Atoi(cmd[1])

	// Admin only
	if !s.IsAdmin(event.Map["sender"].(map[string]interface{})["user_id"].(int64)) {
		return
	}

	_, err := s.donateSrv.AddDonate(ctx, &donateService.AddDonateReq{
		SteamId: steamid,
		Amount:  int32(amount),
	})
	if err != nil {
		log.Error(err)
		return
	}

	// Send message
	err = s.Reply(event, SuccessReply)
	if err != nil {
		log.Error(err)
	}
}
