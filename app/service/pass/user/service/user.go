package service

import (
	pb "backend/app/service/pass/user/api/grpc"
	"backend/app/service/pass/user/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	res, err := s.dao.Info(req.Uid)
	if err != nil {
		if ecode.GetError(err).Code != codes.NotFound {
			return
		}
		// not found -> create
		res = &model.User{
			UID:      req.Uid,
			PassType: 0,
			Point:    0,
		}
		err = s.dao.Create(res)
	}
	resp.Info = &pb.Info{
		Uid:      res.UID,
		PassType: res.PassType,
		Point:    res.Point,
	}
	return
}

func (s *Service) AddPoint(ctx context.Context, req *pb.AddPointReq) (resp *pb.AddPointResp, err error) {
	resp = &pb.AddPointResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if req.Point <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Point(%d)", req.Point)
		return
	}

	res, upgrade, err := s.dao.AddPoint(req.Uid, req.Point)
	if upgrade && res != nil {
		err = s.Upgrade(ctx, res.UID, res.Level(), res.PassType)
	}

	return
}

func (s *Service) Upgrade(ctx context.Context, uid int64, level int32, passType int32) (err error) {
	// TODO: give gift
	return
}

func (s *Service) UpgradePass(ctx context.Context, req *pb.UpgradePassReq) (resp *pb.UpgradePassResp, err error) {
	resp = &pb.UpgradePassResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	err = s.dao.UpgradePass(req.Uid)
	if err != nil {
		return
	}

	// TODO: give all reward before current level
	return
}
