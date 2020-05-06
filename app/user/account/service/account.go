package service

import (
	pb "backend/app/user/account/api/grpc"
	"backend/app/user/account/model"
	"context"
)

func (s *Service) UID(ctx context.Context, req *pb.UIDReq) (res *pb.UIDResp, err error) {
	res = &pb.UIDResp{}

	r, err := s.dao.UID(req.SteamId)
	if err != nil {
		return
	}

	res.Uid = r.UID
	return
}

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (res *pb.InfoResp, err error) {
	res = &pb.InfoResp{}

	r, err := s.dao.Info(req.Uid)
	if err != nil {
		return
	}

	res.Info = &pb.Info{
		SteamId:   r.SteamID,
		Username:  r.Username,
		FirstJoin: r.FirstJoin,
	}
	return
}

func (s *Service) Register(ctx context.Context, req *pb.RegisterReq) (res *pb.RegisterResp, err error) {
	res = &pb.RegisterResp{}

	r, err := s.dao.Register(req.SteamId)
	if err != nil {
		return
	}

	res.Uid = r.UID
	return
}

func (s *Service) ChangeName(ctx context.Context, req *pb.ChangeNameReq) (res *pb.ChangeNameResp, err error) {
	res = &pb.ChangeNameResp{}

	err = s.dao.ChangeName(&model.Info{
		UID:      req.Uid,
		Username: req.Username,
	})
	return
}
