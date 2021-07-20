package service

import (
	pb "backend/app/service/backpack/items/api/grpc/v1"
	"backend/app/service/backpack/items/dao"
	"backend/app/service/backpack/items/model"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_GetItems(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().GetItems().Return([]*model.Item{
		{
			ID:   1,
			Type: 1,
		},
		{
			ID:   3,
			Type: 2,
		},
	}, nil).Times(2)

	srv := &Service{dao: d}

	Convey("Test GetItems", t, func() {
		Convey("Check no type", func() {
			res, err := srv.GetItems(context.TODO(), &pb.GetItemsReq{Type: 0})
			So(err, ShouldBeNil)
			So(len(res.Items), ShouldEqual, 2)
			t.Log(res)
		})
		Convey("Check with type", func() {
			res, err := srv.GetItems(context.TODO(), &pb.GetItemsReq{Type: 2})
			So(err, ShouldBeNil)
			So(len(res.Items), ShouldEqual, 1)
			t.Log(res)
		})
	})
}
