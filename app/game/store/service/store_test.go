package service

import (
	pb "backend/app/game/store/api/grpc"
	"backend/app/game/store/model"
	itemsService "backend/app/service/backpack/items/api/grpc"
	userService "backend/app/service/backpack/user/api/grpc"
	passService "backend/app/service/pass/user/api/grpc"
	moneyService "backend/app/service/user/money/api/grpc"
	vipService "backend/app/service/user/vip/api/grpc"
	ctx "context"
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
			{Id: 1, Available: true},
			{Id: 2, Available: true},
		},
	}, nil).Times(2)

	money := moneyService.NewMockMoneyClient(ctl)
	money.EXPECT().Pay(gomock.Any(), &moneyService.PayReq{
		Uid:    1,
		Price:  1,
		Reason: "购买 ",
	}).Return(nil, nil).Times(2)
	money.EXPECT().Give(gomock.Any(), &moneyService.GiveReq{
		Uid:    1,
		Money:  1,
		Reason: "退款 ",
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
			_, err := srv.BuyItem(ctx.TODO(), &pb.BuyItemReq{
				Uid:    1,
				ItemId: 1,
				Price:  1,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check return work", func() {
			_, err := srv.BuyItem(ctx.TODO(), &pb.BuyItemReq{
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
		res, err := srv.HotSaleList(ctx.TODO(), &pb.HotSaleListReq{})
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
	pass := passService.NewMockUserClient(ctl)
	pass.EXPECT().Info(gomock.Any(), &passService.InfoReq{
		Uid: 1,
	}).Return(&passService.InfoResp{
		Info: &passService.Info{
			Uid:      1,
			PassType: 1,
			Point:    1,
		},
	}, nil)

	srv := &Service{
		items: items,
		user:  user,
		money: money,
		pass:  pass,
	}

	Convey("Test SaleList", t, func() {
		res, err := srv.SaleList(ctx.TODO(), &pb.SaleListReq{
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

func TestService_BuyPass(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	money := moneyService.NewMockMoneyClient(ctl)
	money.EXPECT().Pay(gomock.Any(), &moneyService.PayReq{
		Uid:    1,
		Price:  Pass1Price,
		Reason: "购买通行证高级版",
	}).Return(nil, nil)
	money.EXPECT().Pay(gomock.Any(), &moneyService.PayReq{
		Uid:    1,
		Price:  Pass2Price,
		Reason: "购买通行证高级版",
	}).Return(nil, nil)
	money.EXPECT().Give(gomock.Any(), &moneyService.GiveReq{
		Uid:    1,
		Money:  Pass2Price,
		Reason: "通行证高级版退款",
	}).Return(nil, nil)

	pass := passService.NewMockUserClient(ctl)
	pass.EXPECT().UpgradePass(gomock.Any(), &passService.UpgradePassReq{
		Uid:  1,
		Type: 1,
	}).Return(nil, nil)
	pass.EXPECT().UpgradePass(gomock.Any(), &passService.UpgradePassReq{
		Uid:  1,
		Type: 2,
	}).Return(nil, errors.New("test"))

	srv := &Service{money: money, pass: pass}

	Convey("Test BuyPass", t, func() {
		Convey("Check it work", func() {
			_, err := srv.BuyPass(ctx.TODO(), &pb.BuyPassReq{
				Uid:  1,
				Type: 1,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check error", func() {
			_, err := srv.BuyPass(ctx.TODO(), &pb.BuyPassReq{
				Uid:  1,
				Type: 2,
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestService_BuyVIP(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	money := moneyService.NewMockMoneyClient(ctl)
	money.EXPECT().Pay(gomock.Any(), &moneyService.PayReq{
		Uid:    1,
		Price:  VIPMonthPrice,
		Reason: "续费 VIP 1个月",
	}).Return(nil, nil)
	money.EXPECT().Pay(gomock.Any(), &moneyService.PayReq{
		Uid:    1,
		Price:  VIPSemiannualPrice,
		Reason: "续费 VIP 6个月",
	}).Return(nil, nil)
	money.EXPECT().Give(gomock.Any(), &moneyService.GiveReq{
		Uid:    1,
		Money:  VIPSemiannualPrice,
		Reason: "VIP退款",
	}).Return(nil, nil)

	vip := vipService.NewMockVIPClient(ctl)
	vip.EXPECT().Renewal(gomock.Any(), &vipService.RenewalReq{
		Uid:    1,
		Length: Month,
	}).Return(nil, nil)
	vip.EXPECT().Renewal(gomock.Any(), &vipService.RenewalReq{
		Uid:    1,
		Length: Month * 6,
	}).Return(nil, errors.New("test"))
	vip.EXPECT().AddPoint(gomock.Any(), &vipService.AddPointReq{
		Uid:      1,
		AddPoint: BuyVIPPoint,
	}).Return(nil, nil)

	srv := &Service{money: money, vip: vip}

	Convey("Test BuyVIP", t, func() {
		Convey("Check it work", func() {
			_, err := srv.BuyVIP(ctx.TODO(), &pb.BuyVIPReq{
				Uid:   1,
				Month: 1,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check error", func() {
			_, err := srv.BuyVIP(ctx.TODO(), &pb.BuyVIPReq{
				Uid:   1,
				Month: 6,
			})
			So(err, ShouldBeNil)
		})
	})
}
