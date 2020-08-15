package service

import (
	backpackService "backend/app/service/backpack/user/api/grpc"
	rewardService "backend/app/service/pass/reward/api/grpc"
	pb "backend/app/service/pass/user/api/grpc"
	"backend/app/service/pass/user/dao"
	"backend/app/service/pass/user/model"
	money "backend/app/service/user/money/api/grpc"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_AddPoint(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	// Upgrade
	d.EXPECT().AddPoint(int64(2), int32(1)).Return(&model.User{
		UID:      2,
		PassType: 0,
		Point:    model.PassLevelPoint,
	}, int32(1), nil)
	// Not upgrade
	d.EXPECT().AddPoint(int64(1), int32(1)).Return(&model.User{
		UID:      1,
		PassType: 0,
		Point:    1,
	}, int32(1), nil)
	reward := rewardService.NewMockRewardClient(ctl)
	reward.EXPECT().GetRewards(gomock.Any(), &rewardService.GetRewardsReq{
		Level: 2,
		Min:   2,
	}).Return(&rewardService.GetRewardsResp{
		Season:      1,
		FreeRewards: []*rewardService.Item{},
		AdvRewards:  []*rewardService.Item{},
	}, nil)

	srv := &Service{dao: d, rewardService: reward}

	Convey("Test AddPoint", t, func() {
		Convey("Check not upgrade", func() {
			_, err := srv.AddPoints(context.TODO(), &pb.AddPointsReq{
				Add: []*pb.AddPoints{
					{
						Uid:   1,
						Point: 1,
					},
				},
			})
			So(err, ShouldBeNil)
		})
		Convey("Check upgrade", func() {
			_, err := srv.AddPoints(context.TODO(), &pb.AddPointsReq{
				Add: []*pb.AddPoints{
					{
						Uid:   2,
						Point: 1,
					},
				},
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestService_GiveRewards(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	items := []*rewardService.Item{
		{
			Level:  1,
			Id:     1,
			Amount: 0,
			Length: 0,
		},
	}

	reward := rewardService.NewMockRewardClient(ctl)
	reward.EXPECT().GetRewards(gomock.Any(), &rewardService.GetRewardsReq{
		Level: 1,
		Min:   1,
	}).Return(&rewardService.GetRewardsResp{
		Season:      1,
		FreeRewards: items,
		AdvRewards:  items,
	}, nil).Times(2)
	backpack := backpackService.NewMockUserClient(ctl)
	backpack.EXPECT().AddItems(gomock.Any(), &backpackService.AddItemsReq{
		Uid: 1,
		Items: []*backpackService.Item{
			{
				Id:     1,
				Amount: 0,
				Length: 0,
			},
		},
	}).Return(nil, nil)
	backpack.EXPECT().AddItems(gomock.Any(), &backpackService.AddItemsReq{
		Uid: 2,
		Items: []*backpackService.Item{
			{
				Id:     1,
				Amount: 0,
				Length: 0,
			},
			{
				Id:     1,
				Amount: 0,
				Length: 0,
			},
		},
	}).Return(nil, nil)

	srv := &Service{rewardService: reward, backpackService: backpack}

	Convey("Test GiveRewards", t, func() {
		Convey("Check free pass", func() {
			err := srv.GiveRewards(context.TODO(), &model.User{
				UID:      1,
				Point:    1,
				PassType: 0,
			}, 0)
			So(err, ShouldBeNil)
		})
		Convey("Check adv pass", func() {
			err := srv.GiveRewards(context.TODO(), &model.User{
				UID:      2,
				Point:    1,
				PassType: 1,
			}, 0)
			So(err, ShouldBeNil)
		})
	})
}

func TestService_Info(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Info(int64(1)).Return(&model.User{
		UID:      1,
		PassType: 1,
		Point:    1,
	}, nil)

	srv := &Service{dao: d}

	Convey("Test Info", t, func() {
		Convey("Check it work", func() {
			res, err := srv.Info(context.TODO(), &pb.InfoReq{Uid: 1})
			So(err, ShouldBeNil)
			So(res.Info.Uid, ShouldEqual, 1)
			So(res.Info.PassType, ShouldEqual, 1)
			So(res.Info.Point, ShouldEqual, 1)
		})
	})
}

func TestService_UpgradePass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Info(int64(1)).Return(&model.User{
		UID:      1,
		PassType: 0,
		Point:    model.PassLevelPoint,
	}, nil).Times(2)
	// Invalid
	d.EXPECT().Info(int64(2)).Return(&model.User{
		UID:      2,
		PassType: 1,
		Point:    model.PassLevelPoint,
	}, nil)
	d.EXPECT().UpgradePass(int64(1), int32(1)).Return(nil)
	d.EXPECT().UpgradePass(int64(1), int32(2)).Return(nil)
	d.EXPECT().AddPoint(int64(1), int32(Pass2AddPoint)).Return(&model.User{
		UID:      1,
		PassType: 1,
		Point:    Pass2AddPoint,
	}, Pass2AddPoint/model.PassLevelPoint+1, nil)

	reward := rewardService.NewMockRewardClient(ctl)
	reward.EXPECT().GetRewards(gomock.Any(), &rewardService.GetRewardsReq{
		Level: 2,
		Min:   1,
	}).Return(&rewardService.GetRewardsResp{
		Season: 1,
		AdvRewards: []*rewardService.Item{
			{
				Level:  1,
				Id:     1,
				Amount: 0,
				Length: 0,
			},
			{
				Level:  2,
				Id:     2,
				Amount: 0,
				Length: 0,
			},
		},
	}, nil).Times(2)

	backpack := backpackService.NewMockUserClient(ctl)
	backpack.EXPECT().AddItems(gomock.Any(), &backpackService.AddItemsReq{
		Uid: 1,
		Items: []*backpackService.Item{
			{
				Id:     1,
				Amount: 0,
				Length: 0,
			},
			{
				Id:     2,
				Amount: 0,
				Length: 0,
			},
		},
	}).Return(nil, nil).Times(2)

	srv := &Service{
		dao:             d,
		rewardService:   reward,
		backpackService: backpack,
	}

	Convey("Test UpgradePass", t, func() {
		Convey("Check it work", func() {
			_, err := srv.UpgradePass(context.TODO(), &pb.UpgradePassReq{
				Uid:  1,
				Type: 1,
			})
			So(err, ShouldBeNil)
			_, err = srv.UpgradePass(context.TODO(), &pb.UpgradePassReq{
				Uid:  1,
				Type: 2,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check user already have adv pass", func() {
			_, err := srv.UpgradePass(context.TODO(), &pb.UpgradePassReq{
				Uid:  2,
				Type: 1,
			})
			So(err, ShouldNotBeNil)
			t.Log(err)
		})
	})
}

func TestService_GiveMoney(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := money.NewMockMoneyClient(ctl)
	m.EXPECT().Give(gomock.Any(), &money.GiveReq{
		Uid:    1,
		Money:  200,
		Reason: "通行证奖励",
	})

	srv := &Service{moneyService: m}

	Convey("Test GiveMoney", t, func() {
		Convey("Check it work", func() {
			err := srv.GiveMoney(context.TODO(), 1, []*rewardService.Item{
				{
					Level:  1,
					Id:     0,
					Amount: 100,
				},
				{
					Level:  2,
					Id:     0,
					Amount: 100,
				},
				{
					Level:  3,
					Id:     1,
					Amount: 100,
				},
			})
			So(err, ShouldBeNil)
		})
	})
}
