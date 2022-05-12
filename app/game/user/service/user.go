package service

import (
	pb "backend/app/game/user/api/grpc/v1"
	backpack "backend/app/service/backpack/user/api/grpc/v1"
	account_pb "backend/app/service/user/account/api/grpc/v1"
	"backend/pkg/ecode"
	"context"

	"google.golang.org/grpc/codes"
)

func (s *Service) PlayerConnect(ctx context.Context, req *pb.PlayerConnectReq) (resp *pb.PlayerConnectResp, err error) {
	resp = &pb.PlayerConnectResp{}

	// Get UID
	// If player is first time join in our server
	// he do not have UID in our database
	// we help him to register
	uid, err := s.account.UID(ctx, &account_pb.UIDReq{
		SteamId: req.SteamId,
	})
	if err != nil {
		if ecode.GetError(err).Code != codes.NotFound {
			return
		}
		// Not found -> Register -> Get UID
		var info *account_pb.RegisterResp
		info, err = s.account.Register(ctx, &account_pb.RegisterReq{
			SteamId: req.SteamId,
		})
		if err != nil {
			return
		}
		_, err = s.backpack.Init(ctx, &backpack.InitReq{
			Uid: info.Uid,
		})
		if err != nil {
			return
		}
		resp.Uid = info.Uid
	} else {
		// Found UID
		resp.Uid = uid.Uid
	}

	// Get steam groups
	groups, err := s.steam.GetUserGroupList(uint64(req.SteamId))
	if err != nil {
		return
	}
	for _, v := range groups.Groups {
		resp.Groups = append(resp.Groups, v.GID)
	}

	// Change name
	_, err = s.account.ChangeName(ctx, &account_pb.ChangeNameReq{
		Uid:      resp.Uid,
		Username: req.Username,
	})
	return
}

func (s *Service) PlayerDisconnect(ctx context.Context, req *pb.PlayerDisconnectReq) (resp *pb.PlayerDisconnectResp, err error) {
	resp = &pb.PlayerDisconnectResp{}

	// Get UID
	//uid, err := s.account.UID(ctx, &account_pb.UIDReq{
	//	SteamId: req.SteamId,
	//})
	//if err != nil {
	//	return
	//}

	// do some things??
	return
}
