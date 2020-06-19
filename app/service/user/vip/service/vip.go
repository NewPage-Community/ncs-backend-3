package service

import (
	"backend/app/service/user/vip/api/grpc"
	"backend/app/service/user/vip/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) Info(ctx context.Context, req *grpc.InfoReq) (resp *grpc.InfoResp, err error) {
	resp = &grpc.InfoResp{}
	info, err := s._Info(req.Uid)
	if info != nil {
		resp.Info = &grpc.Info{
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

func (s *Service) Renewal(ctx context.Context, req *grpc.RenewalReq) (resp *grpc.RenewalResp, err error) {
	resp = &grpc.RenewalResp{}

	if req.Length <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Length is invalid")
		return
	}

	info, err := s._Info(req.Uid)
	if err != nil {
		return
	}

	info.Renewal(req.Length)
	resp.ExpireDate = info.ExpireDate
	err = s.dao.ExpireTime(info)
	return
}

func (s *Service) AddPoint(ctx context.Context, req *grpc.AddPointReq) (resp *grpc.AddPointResp, err error) {
	resp = &grpc.AddPointResp{}

	if req.AddPoint <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "AddPoint is invalid")
		return
	}

	info, err := s._Info(req.Uid)
	if err != nil {
		return
	}

	info.Point += int(req.AddPoint)
	resp.Point = int32(info.Point)
	err = s.dao.Point(info)
	return
}
