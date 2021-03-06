package service

import (
	pb "backend/app/game/cookie/api/grpc"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) GetCookie(ctx context.Context, req *pb.GetCookieReq) (resp *pb.GetCookieResp, err error) {
	resp = &pb.GetCookieResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if len(req.Key) == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Key is empty")
		return
	}

	cookie, err := s.dao.Get(req.Uid)
	if err != nil {
		return
	}
	resp.Value, resp.Exist = cookie.Cookie[req.Key]
	return
}

func (s *Service) GetAllCookie(ctx context.Context, req *pb.GetAllCookieReq) (resp *pb.GetAllCookieResp, err error) {
	resp = &pb.GetAllCookieResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	cookie, err := s.dao.Get(req.Uid)
	if err != nil {
		return
	}
	resp.Cookie = cookie.Cookie
	return
}

func (s *Service) SetCookie(ctx context.Context, req *pb.SetCookieReq) (resp *pb.SetCookieResp, err error) {
	resp = &pb.SetCookieResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if len(req.Key) == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Key is empty")
		return
	}

	err = s.dao.Set(req.Uid, req.Key, req.Value)
	return
}
