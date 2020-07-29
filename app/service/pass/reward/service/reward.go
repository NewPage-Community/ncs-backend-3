package service

import (
	pb "backend/app/service/pass/reward/api/grpc"
	"backend/app/service/pass/reward/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) GetRewards(ctx context.Context, req *pb.GetRewardsReq) (resp *pb.GetRewardsResp, err error) {
	if req.Level < req.Min {
		err = ecode.Errorf(codes.InvalidArgument, "Min can not bigger than Level")
		return
	}
	resp = &pb.GetRewardsResp{
		Season:      s.reward.Season,
		FreeRewards: s.getRewards(0, req.Level, req.Min),
		AdvRewards:  s.getRewards(1, req.Level, req.Min),
	}
	return
}

func (s *Service) getRewards(passType int32, level int32, min int32) []*pb.Item {
	var items []*pb.Item
	var list *[]model.Item

	if passType > 0 {
		list = &s.reward.AdvRewards
	} else {
		list = &s.reward.FreeRewards
	}

	for _, v := range *list {
		if level > 0 {
			if (level != v.Level && min == 0) || (level < v.Level || v.Level < min) {
				continue
			}
		}
		items = append(items, &pb.Item{
			Level:  v.Level,
			Id:     v.ID,
			Amount: v.Amount,
			Length: v.Length,
		})
	}
	return items
}
