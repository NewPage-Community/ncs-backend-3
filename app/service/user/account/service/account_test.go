package service

import (
	pb "backend/app/service/user/account/api/grpc"
	"backend/app/service/user/account/dao"
	"backend/app/service/user/account/model"
	"backend/app/service/user/account/test"
	"context"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestService_ChangeName(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().ChangeName(gomock.Eq(&model.Info{
		UID:      test.UID,
		Username: test.UserName,
	})).Return(nil)

	srv := &Service{dao: m}
	Convey("Test service ChangeName", t, func() {
		res, err := srv.ChangeName(context.Background(),
			&pb.ChangeNameReq{Uid: test.UID, Username: test.UserName})
		Convey("Then check it work", func() {
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
	})
}

func TestService_Info(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().Info(gomock.Eq(test.UID)).Return(&model.Info{
		SteamID:   test.SteamID,
		Username:  test.UserName,
		FirstJoin: time.Now().UTC().Unix(),
	}, nil)

	srv := &Service{dao: m}
	Convey("Test service Info", t, func() {
		res, err := srv.Info(context.Background(),
			&pb.InfoReq{Uid: test.UID})
		Convey("Then check it work", func() {
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
		t.Log(res)
	})
}

func TestService_Register(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().Register(gomock.Eq(test.SteamID)).Return(&model.Info{
		UID: test.UID,
	}, nil)

	srv := &Service{dao: m}
	Convey("Test service Register", t, func() {
		res, err := srv.Register(context.Background(),
			&pb.RegisterReq{SteamId: test.SteamID})
		Convey("Then check it work", func() {
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
		t.Log(res)
	})
}

func TestService_UID(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().UID(gomock.Eq(test.SteamID)).Return(&model.Info{
		UID: test.UID,
	}, nil)

	srv := &Service{dao: m}
	Convey("Test service UID", t, func() {
		res, err := srv.UID(context.Background(),
			&pb.UIDReq{SteamId: test.SteamID})
		Convey("Then check it work", func() {
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
		t.Log(res)
	})
}
