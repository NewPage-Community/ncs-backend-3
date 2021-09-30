package service

import (
	serverService "backend/app/game/server/api/grpc/v1"
	"backend/pkg/log"
	"context"
	"fmt"
	"strconv"

	qqModel "github.com/miRemid/amy/websocket/model"
)

func (s *Service) getServerStatus(event qqModel.CQEvent, groupID int64) {
	// Build message
	var msg string
	resp, err := s.serverSrv.AllInfo(context.Background(), &serverService.AllInfoReq{A2S: true})
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range resp.Info {
		if v.A2SInfo.MaxPlayers > 0 {
			if s.IsAdminGroup(groupID) {
				msg += fmt.Sprintf("#%d ", v.ServerId)
			}
			msg += fmt.Sprintf("%s 🗺 %s ", v.ShortName, v.A2SInfo.Map)
			if v.A2SInfo.Players > 0 {
				msg += fmt.Sprintf("♿️ %d/%d", v.A2SInfo.Players, v.A2SInfo.MaxPlayers)
			} else {
				msg += "🈳"
			}
			msg += "\n"
		}
	}
	msg += "仪表盘: " + ServerURL

	// Send message
	err = s.Reply(event, msg)
	if err != nil {
		log.Error(err)
	}
}

func (s *Service) changeMap(event qqModel.CQEvent, cmd []string) {
	serverID, _ := strconv.Atoi(cmd[0])
	mapName := cmd[1]

	// Invalid args
	if serverID <= 0 || len(mapName) == 0 {
		return
	}

	_, err := s.serverSrv.Rcon(context.Background(), &serverService.RconReq{
		ServerId: int32(serverID),
		Cmd:      fmt.Sprintf("changelevel %s", mapName),
	})
	if err != nil {
		log.Error(err)
		return
	}
}
