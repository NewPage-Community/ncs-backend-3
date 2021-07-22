package service

import (
	pb "backend/app/service/auth/qq/api/grpc/v1"
	"backend/app/service/auth/qq/model"
	"backend/pkg/ecode"
	"backend/pkg/json"
	"backend/pkg/jwt"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"io/ioutil"
	"net/http"
	"strings"
)

func (s *Service) AuthQQ(ctx context.Context, req *pb.AuthQQReq) (resp *pb.AuthQQResp, err error) {
	resp = &pb.AuthQQResp{}

	// Get OpenID
	openID, err := GetOpenID(req.AccessToken)
	if err != nil {
		err = ecode.Errorf(codes.Unauthenticated, "failed to get openID: %v", err)
		return
	}

	// Get UID
	qqConnect, err := s.dao.GetUID(openID)
	if err != nil {
		return
	}

	// Add to connect token
	resp.JwtString, err = s.config.JWT.NewTokenString(*qqConnect.GetJWTPayload(&jwt.Payload{}))

	return
}

func (s *Service) BindQQ(ctx context.Context, req *pb.BindQQReq) (resp *pb.BindQQResp, err error) {
	resp = &pb.BindQQResp{}

	// Get info
	uid := model.GetQQConnectFromJWTPayload(s.config.JWT.PayloadFormContext(ctx)).UID
	if uid == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid uid")
		return
	}
	openID, err := GetOpenID(req.AccessToken)
	if err != nil {
		err = ecode.Errorf(codes.Unauthenticated, "failed to get openID: %v", err)
		return
	}

	// Bind account
	err = s.dao.BindQQ(model.QQConnect{
		UID:    uid,
		OpenID: openID,
	})
	return
}

func (s *Service) UnbindQQ(ctx context.Context, req *pb.UnbindQQReq) (resp *pb.UnbindQQResp, err error) {
	resp = &pb.UnbindQQResp{}

	// Get UID
	uid := model.GetQQConnectFromJWTPayload(s.config.JWT.PayloadFormContext(ctx)).UID
	if uid == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid uid")
		return
	}

	err = s.dao.UnbindQQ(model.QQConnect{UID: uid})
	return
}

func (s *Service) QQStatus(ctx context.Context, req *pb.QQStatusReq) (resp *pb.QQStatusResp, err error) {
	resp = &pb.QQStatusResp{}
	payload := s.config.JWT.PayloadFormContext(ctx)

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
	resp.JwtString, err = s.config.JWT.NewTokenString(*qqConnect.GetJWTPayload(payload))

	return
}
func GetOpenID(accessToken string) (openID string, err error) {
	resp, err := http.Get("https://graph.qq.com/oauth2.0/me?access_token=" + accessToken)
	if err != nil {
		return
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return
	}

	if !strings.Contains(string(data), "callback") {
		err = ecode.Errorf(codes.Unknown, "invalid response")
		return
	}
	data = data[10 : len(data)-4]
	fmt.Println(string(data))

	var callback struct {
		Error            int    `json:"error"`
		ErrorDescription string `json:"error_description"`
		OpenID           string `json:"openid"`
		ClientID         int    `json:"client_id"`
	}
	err = json.Unmarshal(data, &callback)

	if callback.Error > 0 && len(callback.ErrorDescription) > 0 {
		err = ecode.Errorf(codes.Unknown, "%d - %s", callback.Error, callback.ErrorDescription)
		return
	}
	openID = callback.OpenID

	return
}
