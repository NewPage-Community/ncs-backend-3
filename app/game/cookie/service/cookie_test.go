package service

import (
	pb "backend/app/game/cookie/api/grpc/v1"
	"backend/app/game/cookie/dao"
	"backend/app/game/cookie/model"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_GetAllCookie(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	cookie := make(map[string]string)
	cookie["test"] = "test"

	d := dao.NewMockDao(ctl)
	d.EXPECT().Get(int64(1)).Return(&model.Cookie{
		UID:    1,
		Cookie: cookie,
	}, nil)

	srv := &Service{dao: d}
	Convey("Test GetAllCookie", t, func() {
		Convey("Check it work", func() {
			res, err := srv.GetAllCookie(context.TODO(), &pb.GetAllCookieReq{Uid: 1})
			So(err, ShouldBeNil)
			So(res.Cookie, ShouldNotBeNil)
			t.Log(res)
		})
	})
}

func TestService_GetCookie(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	cookie := make(map[string]string)
	cookie["test"] = "test"

	d := dao.NewMockDao(ctl)
	d.EXPECT().Get(int64(1)).Return(&model.Cookie{
		UID:    1,
		Cookie: cookie,
	}, nil).Times(2)

	srv := &Service{dao: d}
	Convey("Test GetCookie", t, func() {
		Convey("Check exist key", func() {
			res, err := srv.GetCookie(context.TODO(), &pb.GetCookieReq{
				Uid: 1,
				Key: "test",
			})
			So(err, ShouldBeNil)
			So(res.Value, ShouldEqual, "test")
			So(res.Exist, ShouldBeTrue)
		})
		Convey("Check not exist key", func() {
			res, err := srv.GetCookie(context.TODO(), &pb.GetCookieReq{
				Uid: 1,
				Key: "test1",
			})
			So(err, ShouldBeNil)
			So(res.Value, ShouldEqual, "")
			So(res.Exist, ShouldBeFalse)
		})
	})
}

func TestService_SetCookie(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	cookie := make(map[string]string)
	cookie["test"] = "test"

	d := dao.NewMockDao(ctl)
	d.EXPECT().Set(int64(1), "test", "test").Return(nil)

	srv := &Service{dao: d}
	Convey("Test SetCookie", t, func() {
		Convey("Check it work", func() {
			_, err := srv.SetCookie(context.TODO(), &pb.SetCookieReq{
				Uid:   1,
				Key:   "test",
				Value: "test",
			})
			So(err, ShouldBeNil)
		})
	})
}
