package service

import (
	qqBot "backend/app/bot/qq/api/grpc/v1"
	a2sSrv "backend/app/game/a2s/api/grpc/v1"
	pb "backend/app/game/server/api/grpc"
	"backend/app/game/server/model"
	"backend/pkg/ecode"
	"backend/pkg/log"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"sort"
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

	servers := make(map[string]*model.Info)
	ips := make([]string, 0)

	for _, server := range res {
		ips = append(ips, server.Address)
		servers[server.Address] = server
	}

	a2s, _ := s.a2s.A2SQuery(ctx, &a2sSrv.A2SQueryReq{Address: ips})

	for _, server := range a2s.Servers {
		info := servers[server.Address]
		resp.Info = append(resp.Info, &pb.Info{
			ServerId:  info.ServerID,
			ModId:     info.ModID,
			GameId:    info.GameID,
			Rcon:      info.Rcon,
			Hostname:  info.Hostname,
			Address:   server.Address,
			ShortName: info.ShortName,
			A2SInfo:   server.Info,
			A2SPlayer: server.Player,
		})
	}

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
		// Skip empty rcon password
		if len(server.Rcon) == 0 {
			continue
		}
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
