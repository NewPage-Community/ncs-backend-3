package service

import (
	"backend/app/service/auth/jwt"
	pb "backend/app/service/auth/qq/api/grpc/v1"
	"backend/app/service/auth/qq/model"
	"backend/pkg/ecode"
	"backend/pkg/json"
	"context"
	"google.golang.org/grpc/codes"
	"io/ioutil"
	"net/http"
)

func (s *Service) SignInWithQQ(ctx context.Context, req *pb.SignInWithQQReq) (resp *pb.SignInWithQQResp, err error) {
	resp = &pb.SignInWithQQResp{
		AuthUrl: s.config.QQConnect.AuthCodeURL("state"),
	}
	return
}

func (s *Service) GetUIDWithCode(ctx context.Context, req *pb.GetUIDWithCodeReq) (resp *pb.GetUIDWithCodeResp, err error) {
	resp = &pb.GetUIDWithCodeResp{}
	// Get context token
	ctxToken, err := jwt.GetToken(req.JwtString, *s.config.JWT)
	if err != nil {
		ctxToken = &jwt.Token{}
	}

	// Get access token
	token, err := s.config.QQConnect.Exchange(ctx, req.AuthorizationCode)
	if err != nil {
		return
	}

	// Get OpenID
	openID, err := GetOpenID(token.AccessToken)
	if err != nil {
		return
	}

	// Get UID
	qqConnect, err := s.dao.GetUID(openID)
	if err != nil {
		return
	}

	// Add to connect token
	ctxToken.UID = qqConnect.UID
	tokenString, err := ctxToken.NewTokenString(*s.config.JWT)
	if err != nil {
		return
	}
	resp.JwtString = tokenString

	return
}

func (s *Service) BindQQ(ctx context.Context, req *pb.BindQQReq) (resp *pb.BindQQResp, err error) {
	resp = &pb.BindQQResp{}
	// Get context token
	ctxToken, err := jwt.GetToken(req.JwtString, *s.config.JWT)
	if err != nil {
		return
	}
	if ctxToken.UID == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid token")
		return
	}

	// Get access token
	token, err := s.config.QQConnect.Exchange(ctx, req.AuthorizationCode)
	if err != nil {
		return
	}

	// Get OpenID
	openID, err := GetOpenID(token.AccessToken)
	if err != nil {
		return
	}

	err = s.dao.BindQQ(model.QQConnect{
		UID:    ctxToken.UID,
		OpenID: openID,
	})
	return
}

func (s *Service) UnbindQQ(ctx context.Context, req *pb.UnbindQQReq) (resp *pb.UnbindQQResp, err error) {
	resp = &pb.UnbindQQResp{}
	// Get context token
	ctxToken, err := jwt.GetToken(req.JwtString, *s.config.JWT)
	if err != nil {
		return
	}
	if ctxToken.UID == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid token")
		return
	}

	err = s.dao.UnbindQQ(model.QQConnect{UID: ctxToken.UID})
	return
}

func (s *Service) GetQQConnectStatus(ctx context.Context, req *pb.GetQQConnectStatusReq) (resp *pb.GetQQConnectStatusResp, err error) {
	resp = &pb.GetQQConnectStatusResp{}
	// Get context token
	ctxToken, err := jwt.GetToken(req.JwtString, *s.config.JWT)
	if err != nil {
		return
	}
	if ctxToken.UID == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid token")
		return
	}

	qqConnect, err := s.dao.GetStatus(ctxToken.UID)
	if err != nil {
		return
	}
	resp.Uid = qqConnect.UID
	resp.Openid = qqConnect.OpenID
	return
}

func GetOpenID(token string) (openID string, err error) {
	resp, err := http.Get("https://graph.qq.com/oauth2.0/me?access_token=" + token)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	var info map[string]string
	err = json.Unmarshal(data, &info)
	if err != nil {
		return
	}
	openID, ok := info["openid"]
	if !ok {
		err = ecode.Errorf(codes.Unknown, "missing openid")
		return
	}
	return
}
