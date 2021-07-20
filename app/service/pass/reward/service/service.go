package service

import (
	pb "backend/app/service/pass/reward/api/grpc/v1"
	"backend/app/service/pass/reward/conf"
	"backend/app/service/pass/reward/model"
	"backend/pkg/json"
	"backend/pkg/log"
)

type Service struct {
	reward *model.Reward
	pb.UnimplementedRewardServer
}

func Init(config *conf.Config) (s *Service) {
	s = &Service{
		reward: &model.Reward{
			Season:   config.RewardConf.Season,
			MaxLevel: config.RewardConf.MaxLevel,
		},
	}
	//json
	if err := json.Unmarshal([]byte(config.RewardConf.Free), &s.reward.FreeRewards); err != nil {
		log.Error(err)
	}
	if err := json.Unmarshal([]byte(config.RewardConf.Adv), &s.reward.AdvRewards); err != nil {
		log.Error(err)
	}
	return
}

func (s *Service) Healthy() bool {
	return true
}

func (s *Service) Close() {
}
