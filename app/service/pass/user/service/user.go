package service

import (
	server "backend/app/game/server/api/grpc/v1"
	backpack "backend/app/service/backpack/user/api/grpc/v1"
	reward "backend/app/service/pass/reward/api/grpc/v1"
	pb "backend/app/service/pass/user/api/grpc/v1"
	"backend/app/service/pass/user/model"
	money_pb "backend/app/service/user/money/api/grpc/v1"
	"backend/pkg/ecode"
	"backend/pkg/log"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"sync"
)

const (
	Pass2AddPoint    = model.PassLevelPoint * 10
	UpgradeNotifyCMD = "ncs_pass_upgrade_notify %d %d"
)

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	res, err := s.dao.Info(req.Uid)
	if err != nil {
		return
	}
	resp.Info = &pb.Info{
		Uid:      res.UID,
		PassType: res.PassType,
		Point:    res.Point,
	}
	return
}

func (s *Service) AddPoints(ctx context.Context, req *pb.AddPointsReq) (resp *pb.AddPointsResp, err error) {
	resp = &pb.AddPointsResp{}

	if len(req.Add) == 0 {
		return
	}

	wg := sync.WaitGroup{}
	for _, v := range req.Add {
		v := v
		if v.Uid <= 0 || v.Point <= 0 {
			log.Warn("Invalid data", "UID:", v.Uid, "Point:", v.Point)
			continue
		}
		wg.Add(1)
		go func() {
			res, lastLevel, err := s.dao.AddPoint(v.Uid, v.Point)
			if err == nil {
				// Upgrade~
				if lastLevel != res.Level() {
					if err := s.GiveRewards(ctx, res, lastLevel); err != nil {
						log.Error(err)
					}
					if err := s.UpgradeNotify(ctx, v.Uid, res.Level()); err != nil {
						log.Error(err)
					}
				}
			} else {
				log.Error(err)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	return
}

func (s *Service) GiveRewards(ctx context.Context, info *model.User, lastLevel int32) (err error) {
	rewards, err := s.rewardService.GetRewards(ctx, &reward.GetRewardsReq{
		Level: info.Level(),
		Min:   lastLevel + 1,
	})
	if err != nil {
		return
	}

	var items []*backpack.Item
	for _, v := range rewards.FreeRewards {
		items = append(items, &backpack.Item{
			Id:     v.Id,
			Amount: v.Amount,
			Length: v.Length,
		})
	}
	err = s.GiveMoney(ctx, info.UID, rewards.FreeRewards)
	if err != nil {
		log.Error(err)
	}
	if info.PassType > 0 {
		for _, v := range rewards.AdvRewards {
			items = append(items, &backpack.Item{
				Id:     v.Id,
				Amount: v.Amount,
				Length: v.Length,
			})
		}
		err = s.GiveMoney(ctx, info.UID, rewards.AdvRewards)
		if err != nil {
			log.Error(err)
		}
	}

	if len(items) > 0 {
		_, err = s.backpackService.AddItems(ctx, &backpack.AddItemsReq{
			Uid:   info.UID,
			Items: items,
		})
	}
	return
}

func (s *Service) UpgradePass(ctx context.Context, req *pb.UpgradePassReq) (resp *pb.UpgradePassResp, err error) {
	resp = &pb.UpgradePassResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if req.Type != 1 && req.Type != 2 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Type(%d)", req.Type)
		return
	}

	info, err := s.dao.Info(req.Uid)
	if err != nil {
		return
	}
	if info.PassType > 0 {
		err = ecode.Errorf(codes.Unknown, "User already is AdvPass UID(%d)", req.Uid)
		return
	}
	err = s.dao.UpgradePass(req.Uid, req.Type)
	if err != nil {
		return
	}

	// Level 0 do not need to give reward
	if info.Level() > 0 {
		var rewards *reward.GetRewardsResp
		// give all adv rewards before current level
		rewards, err = s.rewardService.GetRewards(ctx, &reward.GetRewardsReq{
			Level: info.Level(),
			Min:   1,
		})
		if err != nil {
			return
		}

		// Give rewards
		var items []*backpack.Item
		for _, v := range rewards.AdvRewards {
			items = append(items, &backpack.Item{
				Id:     v.Id,
				Amount: v.Amount,
				Length: v.Length,
			})
		}
		err = s.GiveMoney(ctx, req.Uid, rewards.AdvRewards)
		if err != nil {
			log.Error(err)
		}
		_, err = s.backpackService.AddItems(ctx, &backpack.AddItemsReq{
			Uid:   req.Uid,
			Items: items,
		})
		if err != nil {
			log.Error(err)
		}
	}

	if req.Type == 2 {
		_, err = s.AddPoints(ctx, &pb.AddPointsReq{
			Add: []*pb.AddPoints{
				{
					Uid:   req.Uid,
					Point: Pass2AddPoint,
				},
			},
		})
	}
	return
}

func (s *Service) GiveMoney(ctx context.Context, uid int64, rewards []*reward.Item) (err error) {
	money := int32(0)
	for _, v := range rewards {
		if v.Id == 0 {
			money += v.Amount
		}
	}

	if money > 0 {
		_, err = s.moneyService.Give(ctx, &money_pb.GiveReq{
			Uid:    uid,
			Money:  money,
			Reason: "通行证奖励",
		})
	}
	return
}

func (s *Service) UpgradeNotify(ctx context.Context, uid int64, level int32) (err error) {
	_, err = s.serverService.RconAll(ctx, &server.RconAllReq{
		Cmd: fmt.Sprintf(UpgradeNotifyCMD, uid, level),
	})
	return
}
