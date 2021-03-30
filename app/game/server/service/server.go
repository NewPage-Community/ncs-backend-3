package service

import (
	pb "backend/app/game/server/api/grpc"
	"backend/app/game/server/model"
	"backend/pkg/ecode"
	"backend/pkg/log"
	"context"
	"google.golang.org/grpc/codes"
	"strconv"
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

	for _, v := range res {
		a2sInfo := pb.A2S_Info{}
		a2sPlayer := make([]*pb.A2S_Player, 0)
		if req.A2S {
			if v.RequestA2S() == nil {
				a2sInfo = pb.A2S_Info{
					Protocol:     int32(v.A2SInfo.Protocol),
					Map:          model.FixA2SString(v.A2SInfo.Map),
					Folder:       model.FixA2SString(v.A2SInfo.Folder),
					Game:         model.FixA2SString(v.A2SInfo.Game),
					Id:           int32(v.A2SInfo.ID),
					Players:      int32(v.A2SInfo.Players),
					MaxPlayers:   int32(v.A2SInfo.MaxPlayers),
					Bots:         int32(v.A2SInfo.Bots),
					ServerType:   int32(v.A2SInfo.ServerType),
					Environment:  int32(v.A2SInfo.Environment),
					Visibility:   int32(v.A2SInfo.Visibility),
					Vac:          int32(v.A2SInfo.VAC),
					Version:      model.FixA2SString(v.A2SInfo.Version),
					Port:         int32(v.A2SInfo.Port),
					SteamId:      v.A2SInfo.SteamID,
					SourceTvPort: int32(v.A2SInfo.SourceTVPort),
					SourceTvName: model.FixA2SString(v.A2SInfo.SourceTVName),
					Keywords:     model.FixA2SString(v.A2SInfo.Keywords),
					GameId:       v.A2SInfo.GameID,
				}
				for _, v := range v.A2SPlayer.Players {
					a2sPlayer = append(a2sPlayer, &pb.A2S_Player{
						Name:     model.FixA2SString(v.Name),
						Score:    int32(v.Score),
						Duration: float32(v.Duration),
					})
				}
			}
		}
		resp.Info = append(resp.Info, &pb.Info{
			ServerId:  v.ServerID,
			ModId:     v.ModID,
			GameId:    v.GameID,
			Rcon:      v.Rcon,
			Hostname:  v.Hostname,
			Address:   v.Address,
			ShortName: v.ShortName,
			A2SInfo:   &a2sInfo,
			A2SPlayer: a2sPlayer,
		})
	}
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
