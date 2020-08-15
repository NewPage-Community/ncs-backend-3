package service

import (
	pb "backend/app/game/pass/api/grpc"
	itemService "backend/app/service/backpack/items/api/grpc"
	userItemService "backend/app/service/backpack/user/api/grpc"
	rewardService "backend/app/service/pass/reward/api/grpc"
	passService "backend/app/service/pass/user/api/grpc"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_GetRewards(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	item := itemService.NewMockItemsClient(ctl)
	item.EXPECT().GetItems(gomock.Any(), &itemService.GetItemsReq{}).Return(&itemService.GetItemsResp{
		Items: []*itemService.Item{
			{
				Id:   1,
				Name: "test",
			},
			{
				Id:   2,
				Name: "test",
			},
		},
	}, nil)
	reward := rewardService.NewMockRewardClient(ctl)
	reward.EXPECT().GetRewards(gomock.Any(), &rewardService.GetRewardsReq{}).Return(&rewardService.GetRewardsResp{
		Season: 1,
		FreeRewards: []*rewardService.Item{
			{
				Id:    1,
				Level: 1,
			},
			{
				Id:    0,
				Level: 2,
			},
		},
		AdvRewards: []*rewardService.Item{
			{
				Id:    2,
				Level: 1,
			},
		},
	}, nil)

	srv := &Service{item: item, reward: reward}
	Convey("Test GetRewards", t, func() {
		Convey("Check it work", func() {
			rewards, err := srv.GetRewards(context.TODO(), &pb.GetRewardsReq{})
			So(err, ShouldBeNil)
			So(rewards.Season, ShouldEqual, 1)
			So(len(rewards.FreeRewards), ShouldEqual, 2)
			So(len(rewards.AdvRewards), ShouldEqual, 1)
			t.Log(rewards)
		})
	})
}

func TestService_UsePointBox(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	userItem := userItemService.NewMockUserClient(ctl)
	userItem.EXPECT().GetItems(gomock.Any(), &userItemService.GetItemsReq{
		Uid: 1,
	}).Return(&userItemService.GetItemsResp{
		Uid: 1,
		Items: []*userItemService.Item{
			{
				Id:     BoxId,
				Amount: 10,
			},
		},
	}, nil).Times(2)
	userItem.EXPECT().RemoveItem(gomock.Any(), &userItemService.RemoveItemReq{
		Uid:  1,
		Item: &userItemService.Item{Id: BoxId},
		All:  false,
	}).Return(nil, nil)
	userItem.EXPECT().RemoveItem(gomock.Any(), &userItemService.RemoveItemReq{
		Uid:  1,
		Item: &userItemService.Item{Id: BoxId},
		All:  true,
	}).Return(nil, nil)

	passUser := passService.NewMockUserClient(ctl)
	passUser.EXPECT().AddPoints(gomock.Any(), &passService.AddPointsReq{
		Add: []*passService.AddPoints{
			{
				Uid:   1,
				Point: BoxPoint,
			},
		},
	}).Return(nil, nil)
	passUser.EXPECT().AddPoints(gomock.Any(), &passService.AddPointsReq{
		Add: []*passService.AddPoints{
			{
				Uid:   1,
				Point: BoxPoint * 10,
			},
		},
	}).Return(nil, nil)

	srv := &Service{userPass: passUser, userItem: userItem}
	Convey("Test UsePointBox", t, func() {
		Convey("Check use 1 box", func() {
			_, err := srv.UsePointBox(context.TODO(), &pb.UsePointBoxReq{
				Uid: 1,
				All: false,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check use all boxs", func() {
			_, err := srv.UsePointBox(context.TODO(), &pb.UsePointBoxReq{
				Uid: 1,
				All: true,
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestService_Info(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	passUser := passService.NewMockUserClient(ctl)
	passUser.EXPECT().Info(gomock.Any(), &passService.InfoReq{
		Uid: 1,
	}).Return(&passService.InfoResp{
		Info: &passService.Info{
			Uid:      1,
			PassType: 1,
			Point:    1,
		},
	}, nil)

	userItem := userItemService.NewMockUserClient(ctl)
	userItem.EXPECT().GetItems(gomock.Any(), &userItemService.GetItemsReq{
		Uid: 1,
	}).Return(&userItemService.GetItemsResp{
		Uid: 1,
		Items: []*userItemService.Item{
			{
				Id:     BoxId,
				Amount: 10,
			},
		},
	}, nil)

	srv := &Service{userItem: userItem, userPass: passUser}
	Convey("Test Info", t, func() {
		Convey("Check it work", func() {
			info, err := srv.Info(context.TODO(), &pb.InfoReq{
				Uid: 1,
			})
			So(err, ShouldBeNil)
			So(info.Uid, ShouldEqual, 1)
			So(info.BoxAmount, ShouldEqual, 10)
			So(info.Type, ShouldEqual, 1)
			So(info.Point, ShouldEqual, 1)
		})
	})
}
