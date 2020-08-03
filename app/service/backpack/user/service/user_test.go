package service

import (
	pb "backend/app/service/backpack/user/api/grpc"
	"backend/app/service/backpack/user/dao"
	"backend/app/service/backpack/user/model"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_AddItems(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	dao := dao.NewMockDao(ctl)
	dao.EXPECT().AddItems(int64(1), &model.Items{
		{
			ID: 1,
		},
	})

	srv := &Service{dao: dao}

	Convey("Test AddItems", t, func() {
		Convey("Check it work", func() {
			_, err := srv.AddItems(context.TODO(), &pb.AddItemsReq{
				Uid: 1,
				Items: []*pb.Item{
					{
						Id: 1,
					},
				},
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestService_GetItems(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	dao := dao.NewMockDao(ctl)
	dao.EXPECT().Get(int64(1)).Return(&model.User{
		UID: 1,
		Items: &model.Items{
			{
				ID: 1,
			},
		},
	}, nil)

	srv := &Service{dao: dao}

	Convey("Test GetItems", t, func() {
		Convey("Check it work", func() {
			res, err := srv.GetItems(context.TODO(), &pb.GetItemsReq{Uid: 1})
			So(err, ShouldBeNil)
			So(res.Uid, ShouldEqual, 1)
			So(len(res.Items), ShouldEqual, 1)
			So(res.Items[0].Id, ShouldEqual, 1)
		})
	})
}

func TestService_RemoveItem(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	dao := dao.NewMockDao(ctl)
	dao.EXPECT().RemoveItem(int64(1), model.Item{ID: 1}, true).Return(nil)

	srv := &Service{dao: dao}

	Convey("Test RemoveItem", t, func() {
		Convey("Check it work", func() {
			_, err := srv.RemoveItem(context.TODO(), &pb.RemoveItemReq{
				Uid: 1,
				Item: &pb.Item{
					Id: 1,
				},
				All: true,
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestService_Init(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	dao := dao.NewMockDao(ctl)
	dao.EXPECT().Create(int64(1)).Return(nil, nil)

	srv := &Service{dao: dao}

	Convey("Test Init", t, func() {
		Convey("Check it work", func() {
			_, err := srv.Init(context.TODO(), &pb.InitReq{Uid: 1})
			So(err, ShouldBeNil)
		})
	})
}
