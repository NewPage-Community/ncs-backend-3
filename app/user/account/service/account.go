package service

import (
	pb "backend/app/user/account/api/grpc"
	"context"
)

func (s *Service) UID(ctx context.Context, req *pb.UIDReq) (*pb.UIDResp, error) {
	res, err := s.dao.UID(req.SteamId)
	if err != nil {
		return nil, err
	}
	return &pb.UIDResp{Uid: res}, nil
}

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (*pb.InfoResp, error) {
	res, err := s.dao.Info(req.Uid)
	if err != nil {
		return nil, err
	}
	return &pb.InfoResp{
		Info: &pb.Info{
			SteamId:   res.SteamId,
			Username:  res.Username,
			FirstJoin: res.FirstJoin,
			LastSeen:  res.LastSeen,
		},
	}, nil
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterReq) (*pb.RegisterResp, error) {
	res, err := s.dao.Register(req.SteamId)
	if err != nil {
		return nil, err
	}
	return &pb.RegisterResp{
		Uid: res,
	}, nil
}
