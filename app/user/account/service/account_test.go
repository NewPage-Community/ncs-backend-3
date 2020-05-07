package service

import (
	pb "backend/app/user/account/api/grpc"
	"backend/app/user/account/dao"
	"backend/app/user/account/model"
	. "backend/app/user/account/test"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestService_ChangeName(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().ChangeName(gomock.Eq(&model.Info{
		UID:      TestUID,
		Username: TestUserName,
	})).Return(nil)

	srv := &Service{dao: m}
	Convey("Test service ChangeName", t, func() {
		res, err := srv.ChangeName(context.Background(),
			&pb.ChangeNameReq{Uid: TestUID, Username: TestUserName})
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
	m.EXPECT().Info(gomock.Eq(TestUID)).Return(&model.Info{
		SteamID:   TestSteamID,
		Username:  TestUserName,
		FirstJoin: time.Now().Unix(),
	}, nil)

	srv := &Service{dao: m}
	Convey("Test service Info", t, func() {
		res, err := srv.Info(context.Background(),
			&pb.InfoReq{Uid: TestUID})
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
	m.EXPECT().Register(gomock.Eq(TestSteamID)).Return(&model.Info{
		UID: TestUID,
	}, nil)

	srv := &Service{dao: m}
	Convey("Test service Register", t, func() {
		res, err := srv.Register(context.Background(),
			&pb.RegisterReq{SteamId: TestSteamID})
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
	m.EXPECT().UID(gomock.Eq(TestSteamID)).Return(&model.Info{
		UID: TestUID,
	}, nil)

	srv := &Service{dao: m}
	Convey("Test service UID", t, func() {
		res, err := srv.UID(context.Background(),
			&pb.UIDReq{SteamId: TestSteamID})
		Convey("Then check it work", func() {
			So(res, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
		t.Log(res)
	})
}
