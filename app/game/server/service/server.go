package service

import (
	pb "backend/app/game/server/api/grpc"
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
		ServerId: res.ServerID,
		ModId:    res.ModID,
		GameId:   res.GameID,
		Rcon:     res.Rcon,
		Hostname: res.Hostname,
		Address:  res.Address,
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
		ServerId: res.ServerID,
		ModId:    res.ModID,
		GameId:   res.GameID,
		Rcon:     res.Rcon,
		Hostname: res.Hostname,
		Address:  res.Address,
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
		resp.Info = append(resp.Info, &pb.Info{
			ServerId: v.ServerID,
			ModId:    v.ModID,
			GameId:   v.GameID,
			Rcon:     v.Rcon,
			Hostname: v.Hostname,
			Address:  v.Address,
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
	resp.Response, err = res.Send(req.Cmd)
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
			_, err := server.Send(req.Cmd)
			if err == nil {
				log.Warn(err)
			}
		}()
	}

	return
}
