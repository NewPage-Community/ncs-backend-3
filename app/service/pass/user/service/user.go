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

	res, upgrade, err := s.dao.AddPoint(req.Uid, req.Point)
	if upgrade && res != nil {
		err = s.GiveRewards(ctx, res.UID, res.Level(), res.PassType)
	}

	return
}

func (s *Service) GiveRewards(ctx context.Context, uid int64, level int32, passType int32) (err error) {
	rewards, err := s.rewardService.GetRewards(ctx, &reward.GetRewardsReq{
		Level: level,
		Front: false,
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
	if passType > 0 {
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
			Uid:   uid,
			Items: items,
		})
	}
	return
}

func (s *Service) UpgradePass(ctx context.Context, req *pb.UpgradePassReq) (resp *pb.UpgradePassResp, err error) {
	resp = &pb.UpgradePassResp{}

	info, err := s.Info(ctx, &pb.InfoReq{Uid: req.Uid})
	if err != nil {
		return
	}
	if info.Info.PassType > 0 {
		err = ecode.Errorf(codes.Unknown, "User already is AdvPass UID(%d)", req.Uid)
		return
	}
	err = s.dao.UpgradePass(req.Uid)
	if err != nil {
		return
	}

	// give all adv reward before current level
	rewards, err := s.rewardService.GetRewards(ctx, &reward.GetRewardsReq{
		Level: info.Info.Point,
		Front: true,
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
