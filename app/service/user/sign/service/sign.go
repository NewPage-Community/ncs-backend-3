package service

import (
	pb "backend/app/service/user/sign/api/grpc"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	res, err := s.dao.Info(req.Uid)
	if res != nil {
		resp.Info = &pb.Info{
			Uid:      res.UID,
			SignTime: res.SignTime.Unix(),
			SignDays: int32(res.SignDays),
		}
	}
	return
}

func (s *Service) Sign(ctx context.Context, req *pb.SignReq) (resp *pb.SignResp, err error) {
	resp = &pb.SignResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	resp.AlreadySigned, err = s.dao.Sign(req.Uid)
	return
}
