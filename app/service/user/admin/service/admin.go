package service

import (
	pb "backend/app/service/user/admin/api/grpc"
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
			Flag:     int32(res.GetFlagBits()),
			Immunity: res.Immunity,
		}
	}
	return
}
