package service

import (
	"backend/app/service/auth/jwt"
	pb "backend/app/service/auth/qq/api/grpc/v1"
	"backend/app/service/auth/qq/model"
	"backend/pkg/ecode"
	"backend/pkg/json"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"io/ioutil"
	"net/http"
	"strings"
)

func (s *Service) GetUID(ctx context.Context, req *pb.GetUIDReq) (resp *pb.GetUIDResp, err error) {
	resp = &pb.GetUIDResp{}

	// TODO: Get OpenID
	var openID string

	// Get UID
	qqConnect, err := s.dao.GetUID(openID)
	if err != nil {
		return
	}

	// Add to connect token
	payload := &jwt.Payload{}
	payload.SetUID(qqConnect.UID)
	tokenString, err := s.config.JWT.NewTokenString(*payload)
	if err != nil {
		return
	}
	resp.JwtString = tokenString

	return
}

func (s *Service) BindQQ(ctx context.Context, req *pb.BindQQReq) (resp *pb.BindQQResp, err error) {
	resp = &pb.BindQQResp{}

	// TODO: Get OpenID
	var openID string

	err = s.dao.BindQQ(model.QQConnect{
		UID:    req.Uid,
		OpenID: openID,
	})
	return
}

func (s *Service) UnbindQQ(ctx context.Context, req *pb.UnbindQQReq) (resp *pb.UnbindQQResp, err error) {
	resp = &pb.UnbindQQResp{}

	err = s.dao.UnbindQQ(model.QQConnect{UID: req.Uid})
	return
}

func (s *Service) GetQQConnectStatus(ctx context.Context, req *pb.GetQQConnectStatusReq) (resp *pb.GetQQConnectStatusResp, err error) {
	resp = &pb.GetQQConnectStatusResp{}
	// Get context token
	payload, err := s.config.JWT.GetTokenPayload(req.JwtString)
	if err != nil {
		return
	}

	// Get UID
	uid, err := payload.GetUID()
	if err != nil {
		return
	}
	if uid == 0 {
		err = ecode.Errorf(codes.Unauthenticated, "invalid uid")
		return
	}

	// Get status
	qqConnect, err := s.dao.GetStatus(uid)
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
	if err != nil {
		return
	}
	if callback.Error > 0 && len(callback.ErrorDescription) > 0 {
		err = ecode.Errorf(codes.Unknown, "%d - %s", callback.Error, callback.ErrorDescription)
		return
	}
	openID = callback.OpenID
	return
}
