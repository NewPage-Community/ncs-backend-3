package service

import (
	pb "backend/app/service/user/money/api/grpc/v1"
	"backend/app/service/user/money/dao"
	"backend/app/service/user/money/model"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestService_Get(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Get(int64(1)).Return(&model.Money{
		UID: 1,
		RMB: 1,
	}, nil)

	srv := &Service{dao: d}

	Convey("Test Get", t, func() {
		Convey("Check it work", func() {
			res, err := srv.Get(context.TODO(), &pb.GetReq{Uid: 1})
			So(err, ShouldBeNil)
			So(res.Uid, ShouldEqual, 1)
			So(res.Money, ShouldEqual, 1)
			t.Log(res)
		})
	})
}

func TestService_Give(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Give(int64(1), int32(1)).Return(nil)
	d.EXPECT().AddRecord(int64(1), int32(1), "test").Return(nil)

	srv := &Service{dao: d}

	Convey("Test Give", t, func() {
		Convey("Check it work", func() {
			_, err := srv.Give(context.TODO(), &pb.GiveReq{
				Uid:    1,
				Money:  1,
				Reason: "test",
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestService_Pay(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Pay(int64(1), int32(1)).Return(nil)
	d.EXPECT().AddRecord(int64(1), int32(-1), "test").Return(nil)

	srv := &Service{dao: d}

	Convey("Test Pay", t, func() {
		Convey("Check it work", func() {
			_, err := srv.Pay(context.TODO(), &pb.PayReq{
				Uid:    1,
				Price:  1,
				Reason: "test",
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestService_Records(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().GetRecords(int64(1)).Return(&model.Records{
		{
			Timestamp: time.Now().Add(time.Second).Unix(),
			Amount:    1,
			Reason:    "test",
		},
		{
			Timestamp: time.Now().Add(5 * time.Second).Unix(),
			Amount:    -1,
			Reason:    "test",
		},
	}, nil)

	srv := &Service{dao: d}

	Convey("Test Records", t, func() {
		Convey("Check it work", func() {
			res, err := srv.Records(context.TODO(), &pb.RecordsReq{
				Uid: 1,
			})
			So(err, ShouldBeNil)
			So(len(res.Records), ShouldEqual, 2)
			t.Log(res)
		})
	})
}
