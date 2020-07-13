package service

import (
	backpack "backend/app/service/backpack/user/api/grpc"
	reward "backend/app/service/pass/reward/api/grpc"
	pb "backend/app/service/pass/user/api/grpc"
	"backend/app/service/pass/user/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	res, err := s.dao.Info(req.Uid)
	if err != nil {
		if ecode.GetError(err).Code != codes.NotFound {
			return
		}
		// not found -> create
		res = &model.User{
			UID:      req.Uid,
			PassType: 0,
			Point:    0,
		}
		err = s.dao.Create(res)
	}
	resp.Info = &pb.Info{
		Uid:      res.UID,
		PassType: res.PassType,
		Point:    res.Point,
	}
	return
}

func (s *Service) AddPoint(ctx context.Context, req *pb.AddPointReq) (resp *pb.AddPointResp, err error) {
	resp = &pb.AddPointResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if req.Point <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Point(%d)", req.Point)
		return
	}

	res, lastLevel, err := s.dao.AddPoint(req.Uid, req.Point)
	if err != nil && res != nil {
		// Upgrade~
		if lastLevel != res.Level() {
			err = s.GiveRewards(ctx, res, lastLevel)
		}
	}

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
	if info.PassType > 0 {
		for _, v := range rewards.AdvRewards {
			items = append(items, &backpack.Item{
				Id:     v.Id,
				Amount: v.Amount,
				Length: v.Length,
			})
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

	info, err := s.dao.Info(req.Uid)
	if err != nil {
		return
	}
	if info.PassType > 0 {
		err = ecode.Errorf(codes.Unknown, "User already is AdvPass UID(%d)", req.Uid)
		return
	}
	err = s.dao.UpgradePass(req.Uid)
	if err != nil {
		return
	}

	// give all adv reward before current level
	rewards, err := s.rewardService.GetRewards(ctx, &reward.GetRewardsReq{
		Level: info.Level(),
		Min:   1,
	})
	if err != nil {
		return
	}

	var items []*backpack.Item
	for _, v := range rewards.AdvRewards {
		items = append(items, &backpack.Item{
			Id:     v.Id,
			Amount: v.Amount,
			Length: v.Length,
		})
	}
	_, err = s.backpackService.AddItems(ctx, &backpack.AddItemsReq{
		Uid:   req.Uid,
		Items: items,
	})
	return
}
