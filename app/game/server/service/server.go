package service

import (
	qqBot "backend/app/bot/qq/api/grpc/v1"
	pb "backend/app/game/server/api/grpc"
	"backend/app/game/server/model"
	"backend/pkg/ecode"
	"backend/pkg/log"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"sort"
	"strconv"
	"sync"
)

func (s *Service) Init(ctx context.Context, req *pb.InitReq) (resp *pb.InitResp, err error) {
	resp = &pb.InitResp{}

	if req.Address == "" || req.Port == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Address or port should not be empty")
		return
	}

	res, err := s.dao.Info(req.Address + ":" + strconv.Itoa(int(req.Port)))

	if err != nil {
		return
	}

	// Random rcon password
	res.GenerateRcon()
	err = s.dao.UpdateRcon(res)

	resp.Info = &pb.Info{
		ServerId:  res.ServerID,
		ModId:     res.ModID,
		GameId:    res.GameID,
		Rcon:      res.Rcon,
		Hostname:  res.Hostname,
		Address:   res.Address,
		ShortName: res.ShortName,
	}
	return
}

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{}

	if req.ServerId <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid ServerID(%d)", req.ServerId)
		return
	}

	res, err := s.dao.InfoWithID(req.ServerId)
	if err != nil {
		return
	}

	resp.Info = &pb.Info{
		ServerId:  res.ServerID,
		ModId:     res.ModID,
		GameId:    res.GameID,
		Rcon:      res.Rcon,
		Hostname:  res.Hostname,
		Address:   res.Address,
		ShortName: res.ShortName,
	}
	return
}

func (s *Service) AllInfo(ctx context.Context, req *pb.AllInfoReq) (resp *pb.AllInfoResp, err error) {
	resp = &pb.AllInfoResp{}

	res, err := s.dao.AllInfo()
	if err != nil {
		return
	}

	var lock sync.Mutex
	var wg sync.WaitGroup
	for _, v := range res {
		wg.Add(1)
		server := v
		// Async query A2S
		go func() {
			a2sInfo := pb.A2S_Info{}
			a2sPlayer := make([]*pb.A2S_Player, 0)
			if req.A2S {
				// Query A2S
				err := server.RequestA2S()
				// Push data
				a2sInfo = pb.A2S_Info{
					Protocol:     int32(server.A2SInfo.Protocol),
					Map:          model.FixA2SString(server.A2SInfo.Map),
					Folder:       model.FixA2SString(server.A2SInfo.Folder),
					Game:         model.FixA2SString(server.A2SInfo.Game),
					Id:           int32(server.A2SInfo.ID),
					Players:      int32(server.A2SInfo.Players),
					MaxPlayers:   int32(server.A2SInfo.MaxPlayers),
					Bots:         int32(server.A2SInfo.Bots),
					ServerType:   int32(server.A2SInfo.ServerType),
					Environment:  int32(server.A2SInfo.Environment),
					Visibility:   int32(server.A2SInfo.Visibility),
					Vac:          int32(server.A2SInfo.VAC),
					Version:      model.FixA2SString(server.A2SInfo.Version),
					Port:         int32(server.A2SInfo.Port),
					SteamId:      server.A2SInfo.SteamID,
					SourceTvPort: int32(server.A2SInfo.SourceTVPort),
					SourceTvName: model.FixA2SString(server.A2SInfo.SourceTVName),
					Keywords:     model.FixA2SString(server.A2SInfo.Keywords),
					GameId:       server.A2SInfo.GameID,
				}
				for _, v := range server.A2SPlayer.Players {
					a2sPlayer = append(a2sPlayer, &pb.A2S_Player{
						Name:     model.FixA2SString(v.Name),
						Score:    int32(v.Score),
						Duration: float32(v.Duration),
					})
				}
				if err != nil {
					log.Error(server.Address, err)
				}
			}
			// Push result to slice
			lock.Lock()
			resp.Info = append(resp.Info, &pb.Info{
				ServerId:  server.ServerID,
				ModId:     server.ModID,
				GameId:    server.GameID,
				Rcon:      server.Rcon,
				Hostname:  server.Hostname,
				Address:   server.Address,
				ShortName: server.ShortName,
				A2SInfo:   &a2sInfo,
				A2SPlayer: a2sPlayer,
			})
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()

	// Sort server list (by ServerID)
	sort.Slice(resp.Info, func(i, j int) bool {
		return resp.Info[i].ServerId < resp.Info[j].ServerId
	})
	return
}

func (s *Service) Rcon(ctx context.Context, req *pb.RconReq) (resp *pb.RconResp, err error) {
	resp = &pb.RconResp{}

	if req.ServerId <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid ServerID(%d)", req.ServerId)
		return
	}

	res, err := s.dao.InfoWithID(req.ServerId)
	if err != nil {
		return
	}
	resp.Response, err = res.RCON(req.Cmd)
	return
}

func (s *Service) RconAll(ctx context.Context, req *pb.RconAllReq) (resp *pb.RconAllResp, err error) {
	resp = &pb.RconAllResp{}

	res, err := s.dao.AllInfo()
	if err != nil {
		return
	}

	// Async
	for i := range res {
		server := res[i]
		go func() {
			_, err := server.RCON(req.Cmd)
			if err != nil {
				log.Warn(err)
			}
		}()
	}

	return
}

func (s *Service) ChangeMapNotify(ctx context.Context, req *pb.ChangeMapNotifyReq) (resp *pb.ChangeMapNotifyResp, err error) {
	resp = &pb.ChangeMapNotifyResp{}

	res, err := s.dao.InfoWithID(req.ServerId)
	if err != nil {
		return
	}

	_, err = s.qq.SendGroupMessage(ctx, &qqBot.SendGroupMessageReq{
		Message:    fmt.Sprintf("%s 更换地图 %s", res.ShortName, req.Map),
		AutoEscape: false,
	})
	return
}
