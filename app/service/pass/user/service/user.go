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

	upgrade, err := s.dao.AddPoint(req.Uid, req.Point)
	if upgrade > 0 {
		err = s.Upgrade(ctx, req.Uid, upgrade)
	}

	return
}

func (s *Service) Update(ctx context.Context, req *pb.UpdateReq) (resp *pb.UpdateResp, err error) {
	resp = &pb.UpdateResp{}

	if req.Info == nil {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Info")
		return
	}
	if req.Info.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Info.Uid)
		return
	}

	err = s.dao.Update(&model.User{
		UID:      req.Info.Uid,
		PassType: req.Info.PassType,
		Point:    req.Info.Point,
	})
	return
}

func (s *Service) Upgrade(ctx context.Context, uid int64, level int32) (err error) {
	// TODO: give gift
	return
}
