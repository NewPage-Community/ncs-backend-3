package service

import (
	serverService "backend/app/game/server/api/grpc/v1"
	accountService "backend/app/service/user/account/api/grpc/v1"
	pb "backend/app/service/user/ban/api/grpc/v1"
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
		// ASync request
		go func() {
			err := s.BanSharedLibOwner(context.Background(), req)
			if err != nil {
				log.Warn(err)
			}
		}()
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
	//blockIP := false
	//if len(req.Ip) > 0 {
	//	blockIP, err = s.dao.IsBlockIP(req.Ip)
	//	if err != nil {
	//		log.Warn(err)
	//	}
	//	resp.Info.BlockIp = blockIP
	//}

	// Check shared game lib owner
	if req.AppId > 0 {
		// ASync check
		go func() {
			res, err1 := s.CheckSharedLibOwnerBan(context.Background(), req)
			if err1 != nil {
				log.Error(err1)
			} else if res.ID > 0 {
				// res.ID > 0 -> ban is recorded
				_, err1 := s.server.RconAll(context.Background(), &serverService.RconAllReq{
					Cmd: fmt.Sprintf(BanNotifyCMD, res.UID, res.Type, res.Reason),
				})
				if err1 != nil {
					log.Error(err1)
				}
			}
		}()
	}
	return
}

func (s *Service) List(ctx context.Context, req *pb.ListReq) (resp *pb.ListResp, err error) {
	resp = &pb.ListResp{}

	uid := uint64(0)
	if req.SteamId > 0 {
		account, err := s.account.UID(ctx, &accountService.UIDReq{SteamId: req.SteamId})
		if err != nil {
			return resp, err
		}
		uid = uint64(account.Uid)
	}

	res, err := s.dao.List(uid)
	if err != nil {
		return
	}

	for _, v := range res {
		resp.List = append(resp.List, &pb.Info{
			Id:         v.ID,
			Uid:        v.UID,
			Ip:         v.IP,
			CreateTime: v.CreateTime,
			ExpireTime: v.ExpireTime,
			Type:       v.Type,
			ServerId:   v.ServerID,
			ModId:      v.ModID,
			GameId:     v.GameID,
			Reason:     v.Reason,
		})
	}
	return
}
