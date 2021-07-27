package service

import (
	accountService "backend/app/service/user/account/api/grpc/v1"
	pb "backend/app/service/user/ban/api/grpc/v1"
	"backend/app/service/user/ban/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
	"time"
)

// CheckSharedLibOwnerBan
// if not found returns nil, nil
// if found returns *model.Ban, nil
func (s *Service) CheckSharedLibOwnerBan(ctx context.Context, req *pb.Info2Req) (ban *model.Ban, err error) {
	ban = &model.Ban{}

	checkPlayer, err := s.account.Info(ctx, &accountService.InfoReq{Uid: int64(req.Uid)})
	if err != nil {
		return
	}

	owner, err := s.steam.IsPlayingSharedGame(uint64(checkPlayer.Info.SteamId), int(req.AppId))
	if err != nil {
		return
	}
	if owner.LenderSteamID == 0 {
		return
	}

	ownerUID, err := s.account.UID(ctx, &accountService.UIDReq{SteamId: int64(owner.LenderSteamID)})
	if err != nil {
		if ecode.GetError(err).Code == codes.NotFound {
			// can not found owner account (no banning info)
			// not error just nil it =w=
			err = nil
		}
		return
	}

	ownerBan, err := s.dao.Info(uint64(ownerUID.Uid))
	if err != nil {
		return
	}
	if !ownerBan.IsBanned(req.ServerId, req.ModId, req.GameId) {
		return
	}

	// Lib owner was got banned
	// ban the player who is checked
	ban = &model.Ban{
		UID:        req.Uid,
		IP:         req.Ip,
		CreateTime: time.Now().Unix(),
		ExpireTime: ownerBan.ExpireTime,
		Type:       ownerBan.Type,
		ServerID:   req.ServerId,
		ModID:      req.ModId,
		GameID:     req.GameId,
		Reason:     ownerBan.Reason + LibSharedBanReason,
	}
	err = s.dao.Add(ban)
	return ban, err
}

func (s *Service) BanSharedLibOwner(ctx context.Context, req *pb.AddReq) error {
	checkPlayer, err := s.account.Info(ctx, &accountService.InfoReq{Uid: int64(req.Info.Uid)})
	if err != nil {
		return err
	}

	owner, err := s.steam.IsPlayingSharedGame(uint64(checkPlayer.Info.SteamId), int(req.Info.AppId))
	if err != nil {
		return err
	}
	if owner.LenderSteamID == 0 {
		return nil
	}

	ownerUID, err := s.account.UID(ctx, &accountService.UIDReq{SteamId: int64(owner.LenderSteamID)})
	if err != nil {
		if ecode.GetError(err).Code != codes.NotFound {
			return err
		}
		// Register account to ban
		reg, err := s.account.Register(ctx, &accountService.RegisterReq{SteamId: int64(owner.LenderSteamID)})
		if err != nil {
			return err
		}
		ownerUID.Uid = reg.Uid
	}

	// Ban check player
	ban := &model.Ban{
		UID:        uint64(ownerUID.Uid),
		CreateTime: time.Now().Unix(),
		ExpireTime: req.Info.ExpireTime,
		Type:       req.Info.Type,
		ServerID:   req.Info.ServerId,
		ModID:      req.Info.ModId,
		GameID:     req.Info.GameId,
		Reason:     req.Info.Reason + LibOwnerBanReason,
	}
	return s.dao.Add(ban)
}
