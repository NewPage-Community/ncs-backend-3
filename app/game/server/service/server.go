package service

import (
	pb "backend/app/game/server/api/grpc"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
	"strconv"
)

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	if req.Address == "" || req.Port == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Address or port should not be empty")
		return
	}

	res, err := s.dao.Info(req.Address + ":" + strconv.Itoa(int(req.Port)))

	if res == nil {
		return
	}

	// Random rcon password
	res.GenerateRcon()
	s.dao.UpdateRcon(res)

	resp = &pb.InfoResp{
		Info: &pb.Info{
			ServerId: res.ServerID,
			ModId:    res.ModID,
			GameId:   res.GameID,
			Rcon:     res.Rcon,
			Hostname: res.Hostname,
			Address:  res.Address,
		},
	}
	return
}
