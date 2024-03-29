package service

import (
	pb "backend/app/service/user/vip/api/grpc/v1"
	"backend/app/service/user/vip/dao"
	"backend/app/service/user/vip/model"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestService_AddPoint(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().AddPoint(gomock.Eq(int64(1)), gomock.Eq(1)).Return(1, nil)

	srv := &Service{dao: m}
	Convey("Test service AddPoint", t, func() {
		res, err := srv.AddPoint(context.Background(), &pb.AddPointReq{
			Uid:      1,
			AddPoint: 1,
		})
		Convey("Check it work", func() {
			So(res.Point, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})
}

func TestService_Renewal(t *testing.T) {
	length := int64(10)
	expireTime := time.Now().UTC().Unix() + length
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().Renewal(gomock.Eq(int64(1)), gomock.Eq(length)).Return(expireTime, nil)

	srv := &Service{dao: m}
	Convey("Test service Renewal", t, func() {
		res, err := srv.Renewal(context.Background(), &pb.RenewalReq{
			Uid:    1,
			Length: length,
		})
		Convey("Check it work", func() {
			So(res.ExpireDate, ShouldEqual, expireTime)
			So(err, ShouldBeNil)
		})
	})
}

func TestService_Info(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().Info(gomock.Eq(&model.VIP{
		UID: 1,
	})).Return(nil)

	srv := &Service{dao: m}
	Convey("Test service Info", t, func() {
		res, err := srv.Info(context.Background(), &pb.InfoReq{
			Uid: 1,
		})
		Convey("Check it work", func() {
			So(res.Info, ShouldNotBeNil)
			So(err, ShouldBeNil)
			t.Log(res.Info.String())
		})
	})
}
