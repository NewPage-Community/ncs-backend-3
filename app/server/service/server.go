package service

import (
	pb "backend/app/server/api/grpc"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	if req.Address == "" {
		err = ecode.Errorf(codes.InvalidArgument, "Address should not be empty")
		return
	}

	res, err := s.dao.Info(req.Address)

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
