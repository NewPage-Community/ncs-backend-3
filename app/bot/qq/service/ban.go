package service

import (
	serverService "backend/app/game/server/api/grpc/v1"
	accountService "backend/app/service/user/account/api/grpc/v1"
	banService "backend/app/service/user/ban/api/grpc/v1"
	"backend/pkg/log"
	"context"
	"fmt"
	"strconv"
	"time"

	qqModel "github.com/miRemid/amy/websocket/model"
)

func (s *Service) banPlayer(event qqModel.CQEvent, cmd []string) {
	ctx := context.Background()
	steamid, _ := strconv.ParseInt(cmd[0], 0, 0)
	days, _ := strconv.Atoi(cmd[1])
	reason := ""
	for _, v := range cmd[2:] {
		reason += v + " "
	}
	reason = reason[:len(reason)-1]

	// Invalid args
	if steamid <= 0 || days <= 0 {
		return
	}

	// Get uid
	account, err := s.accountSrv.UID(ctx, &accountService.UIDReq{SteamId: steamid})
	if err != nil {
		log.Error(err)
		return
	}

	// Add ban
	_, err = s.banSrv.Add(ctx, &banService.AddReq{
		Info: &banService.Info{
			Uid:        uint64(account.Uid),
			ExpireTime: time.Now().AddDate(0, 0, days).Unix(),
			Reason:     reason,
			Type:       1,
		},
	})
	if err != nil {
		log.Error(err)
		return
	}

	// Reply
	err = s.Reply(event, SuccessReply)
	if err != nil {
		log.Error(err)
	}
}

func (s *Service) unBanPlayer(event qqModel.CQEvent, cmd []string) {
	ctx := context.Background()
	steamid, _ := strconv.ParseInt(cmd[0], 0, 0)

	// Invalid args
	if steamid <= 0 {
		return
	}

	// Get uid
	account, err := s.accountSrv.UID(ctx, &accountService.UIDReq{SteamId: steamid})
	if err != nil {
		log.Error(err)
		return
	}

	// Remove ban
	ban, err := s.banSrv.Info(ctx, &banService.InfoReq{Uid: uint64(account.Uid)})
	if ban.Info.Id == 0 {
		if err != nil {
			log.Error(err)
		}
		return
	}
	_, err = s.banSrv.Remove(ctx, &banService.RemoveReq{Id: ban.Info.Id})
	if err != nil {
		log.Error(err)
		return
	}

	// Reply
	err = s.Reply(event, SuccessReply)
	if err != nil {
		log.Error(err)
	}
}

func (s *Service) kickPlayer(event qqModel.CQEvent, cmd []string) {
	ctx := context.Background()
	steamid, _ := strconv.ParseInt(cmd[0], 0, 0)
	reason := ""
	for _, v := range cmd[1:] {
		reason += v + " "
	}
	reason = reason[:len(reason)-1]

	// Invalid args
	if steamid <= 0 {
		return
	}

	// Kick notify
	_, err := s.serverSrv.RconAll(ctx, &serverService.RconAllReq{
		Cmd: fmt.Sprintf("ncs_kick_notify %d %s", steamid, reason),
	})
	if err != nil {
		log.Error(err)
		return
	}

	// Reply
	err = s.Reply(event, SuccessReply)
	if err != nil {
		log.Error(err)
	}
}
