package service

import (
	pb "backend/app/service/user/admin/api/grpc/v1"
	"backend/app/service/user/admin/dao"
	"backend/app/service/user/admin/model"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_Info(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Info(int64(1)).Return(&model.Admin{
		UID:      1,
		Flag:     "abc",
		Immunity: 99,
	}, nil)
	d.EXPECT().Info(int64(2)).Return(&model.Admin{
		UID:      2,
		Flag:     "",
		Immunity: 0,
	}, nil)

	srv := &Service{dao: d}

	Convey("Test Info", t, func() {
		Convey("Check it work", func() {
			res, err := srv.Info(context.TODO(), &pb.InfoReq{
				Uid: 1,
			})
			So(err, ShouldBeNil)
			So(res.Info.Uid, ShouldEqual, 1)
			So(res.Info.Immunity, ShouldEqual, 99)
			So(res.Info.Flag, ShouldEqual, 1<<0|1<<1|1<<2)
			t.Log(res)
		})
		Convey("Check user is not admin", func() {
			res, err := srv.Info(context.TODO(), &pb.InfoReq{
				Uid: 2,
			})
			So(err, ShouldBeNil)
			So(res.Info.Uid, ShouldEqual, 2)
			So(res.Info.Immunity, ShouldEqual, 0)
			So(res.Info.Flag, ShouldEqual, 0)
			t.Log(res)
		})
	})
}
