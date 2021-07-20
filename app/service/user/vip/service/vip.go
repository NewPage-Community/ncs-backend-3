package service

import (
	pb "backend/app/service/user/vip/api/grpc/v1"
	"backend/app/service/user/vip/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%v)", req.Uid)
		return
	}

	info, err := s._Info(req.Uid)
	if info != nil {
		resp.Info = &pb.Info{
			Uid:        info.UID,
			Point:      int32(info.Point),
			Level:      int32(info.Level()),
			ExpireDate: info.ExpireDate,
		}
	}
	return
}

func (s *Service) _Info(uid int64) (info *model.VIP, err error) {
	info = &model.VIP{UID: uid}

	if !info.IsValid() {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%v)", info.UID)
		return
	}

	err = s.dao.Info(info)
	if err != nil {
		if ecode.GetError(err).Code == codes.NotFound {
			err = s.dao.Register(info)
		}
	}
	return
}

func (s *Service) Renewal(ctx context.Context, req *pb.RenewalReq) (resp *pb.RenewalResp, err error) {
	resp = &pb.RenewalResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%v)", req.Uid)
		return
	}
	if req.Length <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Length is invalid")
		return
	}

	exprTime, err := s.dao.Renewal(req.Uid, req.Length)
	resp.ExpireDate = exprTime
	return
}

func (s *Service) AddPoint(ctx context.Context, req *pb.AddPointReq) (resp *pb.AddPointResp, err error) {
	resp = &pb.AddPointResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%v)", req.Uid)
		return
	}
	if req.AddPoint <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "AddPoint is invalid")
		return
	}

	point, err := s.dao.AddPoint(req.Uid, int(req.AddPoint))
	resp.Point = int32(point)
	return
}
