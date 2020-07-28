package service

import (
	pb "backend/app/game/store/api/grpc"
	"backend/app/game/store/model"
	itemsService "backend/app/service/backpack/items/api/grpc"
	userService "backend/app/service/backpack/user/api/grpc"
	moneyService "backend/app/service/user/money/api/grpc"
	"errors"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_BuyItem(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	items := itemsService.NewMockItemsClient(ctl)
	items.EXPECT().GetItems(gomock.Any(), &itemsService.GetItemsReq{
		Type: 0,
	}).Return(&itemsService.GetItemsResp{
		Items: []*itemsService.Item{
			{Id: 1},
			{Id: 2},
		},
	}, nil).Times(2)

	money := moneyService.NewMockMoneyClient(ctl)
	money.EXPECT().Pay(gomock.Any(), &moneyService.PayReq{
		Uid:    1,
		Price:  1,
		Reason: "购买商品 ",
	}).Return(nil, nil).Times(2)
	money.EXPECT().Give(gomock.Any(), &moneyService.GiveReq{
		Uid:    1,
		Money:  1,
		Reason: "商品退款 ",
	}).Return(nil, nil)

	user := userService.NewMockUserClient(ctl)
	user.EXPECT().AddItems(gomock.Any(), &userService.AddItemsReq{
		Uid: 1,
		Items: []*userService.Item{{
			Id:     1,
			Amount: 0,
			Length: 0,
		}},
	}).Return(nil, nil)
	user.EXPECT().AddItems(gomock.Any(), &userService.AddItemsReq{
		Uid: 1,
		Items: []*userService.Item{{
			Id:     2,
			Amount: 0,
			Length: 0,
		}},
	}).Return(nil, errors.New("test"))
	srv := &Service{items: items, money: money, user: user}

	Convey("Test BuyItem", t, func() {
		Convey("Check it work", func() {
			_, err := srv.BuyItem(nil, &pb.BuyItemReq{
				Uid:    1,
				ItemId: 1,
				Price:  1,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check return work", func() {
			_, err := srv.BuyItem(nil, &pb.BuyItemReq{
				Uid:    1,
				ItemId: 2,
				Price:  1,
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestService_HotSaleList(t *testing.T) {
	srv := &Service{
		hotSale: model.HotSale{
			ItemIDs: []int32{1, 2, 3},
		},
	}

	Convey("Test HotSaleList", t, func() {
		res, err := srv.HotSaleList(nil, &pb.HotSaleListReq{})
		Convey("Check it work", func() {
			So(err, ShouldBeNil)
			So(len(res.ItemsId), ShouldEqual, 3)
			t.Log(res)
		})
	})
}

func TestService_SaleList(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	user := userService.NewMockUserClient(ctl)
	user.EXPECT().GetItems(gomock.Any(), &userService.GetItemsReq{
		Uid: 1,
	}).Return(&userService.GetItemsResp{
		Uid: 1,
		Items: []*userService.Item{
			{
				Id:       1,
				Amount:   0,
				ExprTime: 0,
			},
			{
				Id:       2,
				Amount:   1,
				ExprTime: 0,
			},
		},
	}, nil)
	money := moneyService.NewMockMoneyClient(ctl)
	money.EXPECT().Get(gomock.Any(), &moneyService.GetReq{
		Uid: 1,
	}).Return(&moneyService.GetResp{
		Uid:   1,
		Money: 1,
	}, nil)
	items := itemsService.NewMockItemsClient(ctl)
	items.EXPECT().GetItems(gomock.Any(), &itemsService.GetItemsReq{
		Type: 0,
	}).Return(&itemsService.GetItemsResp{
		Items: []*itemsService.Item{
			{Id: 1},
			{Id: 2},
		},
	}, nil)

	srv := &Service{
		items: items,
		user:  user,
		money: money,
	}

	Convey("Test SaleList", t, func() {
		res, err := srv.SaleList(nil, &pb.SaleListReq{
			Uid: 1,
		})
		Convey("Check it work", func() {
			So(err, ShouldBeNil)
			So(res.Money, ShouldEqual, 1)
			So(res.Items[0].AlreadyHave, ShouldEqual, true)
			So(res.Items[1].AlreadyHave, ShouldEqual, false)
			t.Log(res)
		})
	})
}
