package service

import (
	pb "backend/app/user/account/api/grpc"
	"backend/app/user/account/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) UID(ctx context.Context, req *pb.UIDReq) (res *pb.UIDResp, err error) {
	res = &pb.UIDResp{}

	// Invalid
	if req.SteamId == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "SteamID invalid")
		return
	}

	r, err := s.dao.UID(req.SteamId)
	if err != nil {
		return
	}

	res.Uid = r.UID
	return
}

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (res *pb.InfoResp, err error) {
	res = &pb.InfoResp{}

	// Invalid
	if req.Uid == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "UID invalid")
		return
	}

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

	// Invalid
	if req.SteamId == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "SteamID invalid")
		return
	}

	r, err := s.dao.Register(req.SteamId)
	if err != nil {
		return
	}

	res.Uid = r.UID
	return
}

func (s *Service) ChangeName(ctx context.Context, req *pb.ChangeNameReq) (res *pb.ChangeNameResp, err error) {
	res = &pb.ChangeNameResp{}

	// Invalid
	if req.Uid == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "UID invalid")
		return
	}

	err = s.dao.ChangeName(&model.Info{
		UID:      req.Uid,
		Username: req.Username,
	})
	return
}

func (s *Service) PlayerConnect(ctx context.Context, req *pb.PlayerConnectReq) (res *pb.PlayerConnectResp, err error) {
	res = &pb.PlayerConnectResp{}

	// Get UID
	r, err := s.UID(ctx, &pb.UIDReq{
		SteamId: req.SteamId,
	})
	if err != nil {
		return
	}
	// Get info
	r2, err := s.Info(ctx, &pb.InfoReq{
		Uid: r.Uid,
	})
	if err != nil {
		return
	}
	// Update username
	_, err = s.ChangeName(ctx, &pb.ChangeNameReq{
		Uid:      r.Uid,
		Username: req.Username,
	})
	if err != nil {
		return
	}

	res.Uid = r.Uid
	res.FirstJoin = r2.Info.FirstJoin
	return
}
