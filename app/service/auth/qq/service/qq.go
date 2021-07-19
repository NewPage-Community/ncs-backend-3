package service

import (
	pb "backend/app/service/auth/qq/api/grpc/v1"
	"backend/app/service/auth/qq/model"
	"backend/pkg/ecode"
	"backend/pkg/jwt"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) AuthQQ(ctx context.Context, req *pb.AuthQQReq) (resp *pb.AuthQQResp, err error) {
	resp = &pb.AuthQQResp{}

	// Get OpenID
	openID := req.OpenId
	if len(openID) == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid openID")
		return
	}

	// Get UID
	qqConnect, err := s.dao.GetUID(openID)
	if err != nil {
		return
	}

	// Add to connect token
	tokenString, err := s.config.JWT.NewTokenString(*qqConnect.GetJWTPayload(&jwt.Payload{}))
	if err != nil {
		return
	}
	resp.JwtString = tokenString

	return
}

func (s *Service) BindQQ(ctx context.Context, req *pb.BindQQReq) (resp *pb.BindQQResp, err error) {
	resp = &pb.BindQQResp{}

	// Get info
	uid := model.GetQQConnectFromJWTPayload(jwt.PayloadFormContext(ctx)).UID
	if uid == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid uid")
		return
	}
	openID := req.OpenId
	if len(openID) == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid openID")
		return
	}

	// Bind account
	err = s.dao.BindQQ(model.QQConnect{
		UID:    uid,
		OpenID: req.OpenId,
	})
	return
}

func (s *Service) UnbindQQ(ctx context.Context, req *pb.UnbindQQReq) (resp *pb.UnbindQQResp, err error) {
	resp = &pb.UnbindQQResp{}

	// Get UID
	uid := model.GetQQConnectFromJWTPayload(jwt.PayloadFormContext(ctx)).UID
	if uid == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid uid")
		return
	}

	err = s.dao.UnbindQQ(model.QQConnect{UID: uid})
	return
}

func (s *Service) QQStatus(ctx context.Context, req *pb.QQStatusReq) (resp *pb.QQStatusResp, err error) {
	resp = &pb.QQStatusResp{}
	payload := jwt.PayloadFormContext(ctx)

	// Get UID
	uid := model.GetQQConnectFromJWTPayload(payload).UID
	if uid == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid uid")
		return
	}

	// Get status
	qqConnect, err := s.dao.GetStatus(uid)
	if err != nil {
		return
	}

	// Add to connect token
	tokenString, err := s.config.JWT.NewTokenString(*qqConnect.GetJWTPayload(payload))
	if err != nil {
		return
	}
	resp.JwtString = tokenString

	return
}
