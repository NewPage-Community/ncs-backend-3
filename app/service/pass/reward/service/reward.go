package service

import (
	pb "backend/app/service/pass/reward/api/grpc"
	"backend/app/service/pass/reward/model"
	"context"
)

func (s *Service) GetRewards(ctx context.Context, req *pb.GetRewardsReq) (resp *pb.GetRewardsResp, err error) {
	resp = &pb.GetRewardsResp{
		Season:      s.config.Reward.Season,
		FreeRewards: s.getRewards(0, req.Level, req.Front),
		AdvRewards:  s.getRewards(1, req.Level, req.Front),
	}
	return
}

func (s *Service) getRewards(passType int32, level int32, front bool) []*pb.Item {
	var items []*pb.Item
	var list *[]model.Item

	if passType > 0 {
		list = &s.config.Reward.AdvRewards
	} else {
		list = &s.config.Reward.FreeRewards
	}

	for _, v := range *list {
		if level > 0 {
			if (level != v.Level && !front) || (front && level < v.Level) {
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
