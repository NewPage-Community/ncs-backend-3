package service

import (
	pb "backend/app/service/user/money/api/grpc/v1"
	"backend/app/service/user/money/dao"
	"backend/app/service/user/money/model"
	"context"
	"math/rand"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
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
	d.EXPECT().Give(int64(1), uint32(1), "test").Return(nil)

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
	d.EXPECT().Pay(int64(1), uint32(1), "test").Return(nil)

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
	d.EXPECT().GetRecords(int64(1), uint32(1)).Return(func() *model.Records {
		records := &model.Records{}
		for i := 0; i < 10; i++ {
			record := &model.Record{
				ID:     uint(i),
				UID:    1,
				Amount: rand.Int31(),
				Reason: "test",
			}
			record.CreatedAt = time.Now().Add(-time.Duration(rand.Intn(1000)))
			*records = append(*records, record)
		}
		return records
	}(), nil)

	srv := &Service{dao: d}

	Convey("Test Records", t, func() {
		Convey("Check it work", func() {
			res, err := srv.Records(context.TODO(), &pb.RecordsReq{
				Uid:  1,
				Days: 1,
			})
			So(err, ShouldBeNil)
			So(len(res.Records), ShouldEqual, 10)
			t.Log(res)
		})
	})
}

func TestService_Gift(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Gift(uint64(1), uint64(2), uint32(1)).Return(uint32(0), nil)

	srv := &Service{dao: d}

	Convey("Test Gift", t, func() {
		Convey("Check it work", func() {
			res, err := srv.Gift(context.TODO(), &pb.GiftReq{
				Uid:    1,
				Target: 2,
				Money:  1,
			})
			So(err, ShouldBeNil)
			t.Log(res)
		})
	})
}
