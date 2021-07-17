package service

import (
	pb "backend/app/game/a2s/api/grpc/v1"
	"context"
	"github.com/NewPage-Community/go-steam"
	"net"
	"sync"
	"time"
)

func (s *Service) A2SQuery(ctx context.Context, req *pb.A2SQueryReq) (resp *pb.A2SQueryResp, err error) {
	resp = &pb.A2SQueryResp{}
	var wg sync.WaitGroup
	var lock sync.Mutex

	for _, ip := range req.Address {
		ip := ip
		wg.Add(1)

		go func() {
			res := &pb.ServerInfo{Address: ip}
			defer func() {
				lock.Lock()
				resp.Servers = append(resp.Servers, res)
				lock.Unlock()
				wg.Done()
			}()

			server, err := steam.Connect(ip, &steam.ConnectOptions{
				Dial: (&net.Dialer{
					Timeout: 100 * time.Millisecond,
				}).Dial,
			})
			if err != nil {
				return
			}

			a2sInfo, err := server.Info()
			res.Info = &pb.A2SInfo{}
			if err == nil && a2sInfo != nil {
				res.Info.Protocol = int32(a2sInfo.Protocol)
				res.Info.Hostname = a2sInfo.Name
				res.Info.Map = a2sInfo.Map
				res.Info.Folder = a2sInfo.Folder
				res.Info.Game = a2sInfo.Game
				res.Info.Id = int32(a2sInfo.ID)
				res.Info.Players = int32(a2sInfo.Players)
				res.Info.MaxPlayers = int32(a2sInfo.MaxPlayers)
				res.Info.Bots = int32(a2sInfo.Bots)
				res.Info.ServerType = int32(a2sInfo.ServerType)
				res.Info.Environment = int32(a2sInfo.Environment)
				res.Info.Visibility = int32(a2sInfo.Visibility)
				res.Info.Vac = int32(a2sInfo.VAC)
				res.Info.Version = a2sInfo.Version
				res.Info.Port = int32(a2sInfo.Port)
				res.Info.SteamId = a2sInfo.SteamID
				res.Info.SourceTvPort = int32(a2sInfo.SourceTVPort)
				res.Info.SourceTvName = a2sInfo.SourceTVName
				res.Info.Keywords = a2sInfo.Keywords
				res.Info.GameId = a2sInfo.GameID
			}
			a2sPlayer, err := server.PlayersInfo()
			res.Player = make([]*pb.A2SPlayer, 0)
			if err == nil && a2sPlayer != nil {
				for _, player := range a2sPlayer.Players {
					res.Player = append(res.Player, &pb.A2SPlayer{
						Name:     player.Name,
						Score:    int32(player.Score),
						Duration: float32(player.Duration),
					})
				}
			}
			server.Close()
		}()
	}

	wg.Wait()
	return
}
