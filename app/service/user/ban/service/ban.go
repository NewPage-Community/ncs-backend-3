package service

import (
	serverService "backend/app/game/server/api/grpc"
	accountService "backend/app/service/user/account/api/grpc"
	pb "backend/app/service/user/ban/api/grpc"
	"backend/app/service/user/ban/model"
	"backend/pkg/ecode"
	"backend/pkg/log"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"time"
)

const (
	BanNotifyCMD       = "ncs_ban_notify %d %d \"%s\""
	LibSharedBanReason = "(共享者账号被封禁)"
	LibOwnerBanReason  = "(被共享者账号被封禁)"
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
			Id:         res.ID,
			Uid:        res.UID,
			Ip:         res.IP,
			CreateTime: res.CreateTime,
			ExpireTime: res.ExpireTime,
			Type:       res.Type,
			ServerId:   res.ServerID,
			ModId:      res.ModID,
			GameId:     res.GameID,
			Reason:     res.Reason,
		}
	}
	return
}

func (s *Service) Add(ctx context.Context, req *pb.AddReq) (resp *pb.AddResp, err error) {
	resp = &pb.AddResp{}
	now := time.Now().UTC()

	if req.Info == nil {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Info")
		return
	}
	if req.Info.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Info.UID")
		return
	}
	if req.Info.Type == model.BanTypeNone {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Info.Type")
		return
	}
	if req.Info.ExpireTime <= now.Unix() {
		err = ecode.Errorf(codes.InvalidArgument, "Expire time must be after %s", now.Format(time.UnixDate))
		return
	}

	err = s.dao.Add(&model.Ban{
		UID:        req.Info.Uid,
		IP:         req.Info.Ip,
		CreateTime: now.Unix(),
		ExpireTime: req.Info.ExpireTime,
		Type:       req.Info.Type,
		ServerID:   req.Info.ServerId,
		ModID:      req.Info.ModId,
		GameID:     req.Info.GameId,
		Reason:     req.Info.Reason,
	})
	if err != nil {
		return
	}

	// ban shared game lib owner
	if req.Info.AppId > 0 {
		err = s.BanSharedLibOwner(ctx, req)
		if err != nil {
			log.Warn(err)
		}
	}

	rcon, err := s.server.RconAll(context.Background(), &serverService.RconAllReq{
		Cmd: fmt.Sprintf(BanNotifyCMD, req.Info.Uid, req.Info.Type, req.Info.Reason),
	})
	if err != nil {
		log.Warn(err)
	} else if rcon.Success == 0 {
		log.Warn("None server notify!")
	}
	return
}

func (s *Service) Remove(ctx context.Context, req *pb.RemoveReq) (resp *pb.RemoveResp, err error) {
	resp = &pb.RemoveResp{}

	err = s.dao.Remove(&model.Ban{
		ID: req.Id,
	})
	return
}

func (s *Service) BanCheck(ctx context.Context, req *pb.Info2Req) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{
		Info: &pb.Info{
			Uid: req.Uid,
		}}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	res, err := s.dao.Info(req.Uid)
	if err != nil {
		return
	}
	if res.IsBanned(req.ServerId, req.ModId, req.GameId) {
		resp.Info.Id = res.ID
		resp.Info.ExpireTime = res.ExpireTime
		resp.Info.Type = res.Type
		resp.Info.Reason = res.Reason
		return
	}

	// Check block ip
	blockIP := false
	if len(req.Ip) > 0 {
		blockIP, err = s.dao.IsBlockIP(req.Ip)
		if err != nil {
			log.Warn(err)
		}
		resp.Info.BlockIp = blockIP
	}

	// Check shared game lib owner
	if req.AppId > 0 {
		res, err1 := s.CheckSharedLibOwnerBan(ctx, req)
		if err1 != nil {
			log.Warn(err1)
		} else {
			// if owner not banned or can not found
			// just give 0 or empty string to them ;)
			resp.Info.Id = res.ID
			resp.Info.ExpireTime = res.ExpireTime
			resp.Info.Type = res.Type
			resp.Info.Reason = res.Reason
		}
	}
	return
}

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
