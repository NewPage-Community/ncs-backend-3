package service

import (
	pb "backend/app/game/cookie/api/grpc/v1"
	"backend/app/game/cookie/model"
	serverSrv "backend/app/game/server/api/grpc/v1"
	"backend/pkg/ecode"
	"backend/pkg/rpc/gateway"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
)

func (s *Service) GetCookie(ctx context.Context, req *pb.GetCookieReq) (resp *pb.GetCookieResp, err error) {
	resp = &pb.GetCookieResp{}

	// Web gateway force cover UID
	if id := gateway.GetID(ctx); id == "gateway-web" {
		req.Uid = s.jwt.PayloadFormContext(ctx).GetInt64("uid")
	}

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
	var model model.CookieModel
	model, resp.Exist = cookie.Cookie[req.Key]
	resp.Value = model.Value
	resp.LastUpdated = model.LastUpdated
	return
}

func (s *Service) GetAllCookie(ctx context.Context, req *pb.GetAllCookieReq) (resp *pb.GetAllCookieResp, err error) {
	resp = &pb.GetAllCookieResp{}

	// Web gateway force cover UID
	if id := gateway.GetID(ctx); id == "gateway-web" {
		req.Uid = s.jwt.PayloadFormContext(ctx).GetInt64("uid")
	}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	cookie, err := s.dao.Get(req.Uid)
	if err != nil {
		return
	}
	resp.Cookie = make(map[string]string)
	for key, model := range cookie.Cookie {
		resp.Cookie[key] = model.Value
	}
	return
}

func (s *Service) SetCookie(ctx context.Context, req *pb.SetCookieReq) (resp *pb.SetCookieResp, err error) {
	resp = &pb.SetCookieResp{}

	// Web gateway force cover UID
	if id := gateway.GetID(ctx); id == "gateway-web" {
		req.Uid = s.jwt.PayloadFormContext(ctx).GetInt64("uid")
	}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if len(req.Key) == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Key is empty")
		return
	}
	if req.Timestamp <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Timestamp is empty")
		return
	}

	err = s.dao.Set(req.Uid, req.Key, req.Value, req.Timestamp)

	// Game server notify
	if err == nil && req.Notify {
		_, _ = s.server.RconAll(ctx, &serverSrv.RconAllReq{
			Cmd: fmt.Sprintf("ncs_cookie_notify %d", req.Uid),
		})
	}
	return
}
