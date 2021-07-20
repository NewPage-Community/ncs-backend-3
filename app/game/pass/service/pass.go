package service

import (
	pb "backend/app/game/pass/api/grpc/v1"
	itemService "backend/app/service/backpack/items/api/grpc/v1"
	userItemService "backend/app/service/backpack/user/api/grpc/v1"
	rewardService "backend/app/service/pass/reward/api/grpc/v1"
	userPassService "backend/app/service/pass/user/api/grpc/v1"
	passModel "backend/app/service/pass/user/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

const (
	BoxId    = passModel.PassPointBoxID
	BoxPoint = passModel.PassLevelPoint
)

func (s *Service) GetRewards(ctx context.Context, req *pb.GetRewardsReq) (resp *pb.GetRewardsResp, err error) {
	resp = &pb.GetRewardsResp{}

	items, err := s.item.GetItems(ctx, &itemService.GetItemsReq{})
	if err != nil {
		return
	}
	nameMap := make(map[int32]string)
	for _, v := range items.Items {
		nameMap[v.Id] = v.Name
	}
	// 0 = rmb
	nameMap[0] = "软妹币"

	rewards, err := s.reward.GetRewards(ctx, &rewardService.GetRewardsReq{})
	if err != nil {
		return
	}

	// free rewards
	for _, v := range rewards.FreeRewards {
		resp.FreeRewards = append(resp.FreeRewards, &pb.Item{
			Id:     v.Id,
			Level:  v.Level,
			Name:   nameMap[v.Id],
			Amount: v.Amount,
			Length: v.Length,
		})
	}
	// adv rewards
	for _, v := range rewards.AdvRewards {
		resp.AdvRewards = append(resp.AdvRewards, &pb.Item{
			Id:     v.Id,
			Level:  v.Level,
			Name:   nameMap[v.Id],
			Amount: v.Amount,
			Length: v.Length,
		})
	}

	resp.Season = rewards.Season
	return
}

func (s *Service) UsePointBox(ctx context.Context, req *pb.UsePointBoxReq) (resp *pb.UsePointBoxResp, err error) {
	resp = &pb.UsePointBoxResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	items, err := s.userItem.GetItems(ctx, &userItemService.GetItemsReq{Uid: req.Uid})
	if err != nil {
		return
	}
	boxNum := int32(0)
	for _, v := range items.Items {
		if v.Id == BoxId {
			boxNum = v.Amount
			break
		}
	}

	_, err = s.userItem.RemoveItem(ctx, &userItemService.RemoveItemReq{
		Uid:  req.Uid,
		Item: &userItemService.Item{Id: BoxId},
		All:  req.All,
	})
	if err != nil {
		return
	}

	if !req.All {
		boxNum = 1
	}
	_, err = s.userPass.AddPoints(ctx, &userPassService.AddPointsReq{Add: []*userPassService.AddPoints{
		{
			Uid:   req.Uid,
			Point: BoxPoint * boxNum,
		},
	}})
	return
}

func (s *Service) Info(ctx context.Context, req *pb.InfoReq) (resp *pb.InfoResp, err error) {
	resp = &pb.InfoResp{}

	info, err := s.userPass.Info(ctx, &userPassService.InfoReq{Uid: req.Uid})
	if err != nil {
		return
	}

	items, err := s.userItem.GetItems(ctx, &userItemService.GetItemsReq{Uid: req.Uid})
	if err != nil {
		return
	}

	boxNum := int32(0)
	for _, v := range items.Items {
		if v.Id == BoxId {
			boxNum = v.Amount
			break
		}
	}

	resp.Uid = req.Uid
	resp.Point = info.Info.Point
	resp.Type = info.Info.PassType
	resp.BoxAmount = boxNum
	return
}
