package service

import (
	pb "backend/app/service/user/account/api/grpc"
	"backend/app/service/user/account/model"
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

func (s *Service) GetAllUID(ctx context.Context, req *pb.GetAllUIDReq) (res *pb.GetAllUIDResp, err error) {
	res = &pb.GetAllUIDResp{}

	users, err := s.dao.GetAllUID()
	if err != nil {
		return
	}

	res.Uid = *users
	return
}
