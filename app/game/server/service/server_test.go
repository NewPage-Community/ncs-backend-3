package service

import (
	pb "backend/app/game/server/api/grpc"
	"backend/app/game/server/dao"
	"backend/app/game/server/model"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_AllInfo(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().AllInfo().Return([]*model.Info{
		{
			ServerID: 1,
			ModID:    1,
			GameID:   1,
			Rcon:     "test",
			Hostname: "test",
			Address:  "127.0.0.1:27015",
		},
		{
			ServerID: 2,
			ModID:    1,
			GameID:   1,
			Rcon:     "test",
			Hostname: "test",
			Address:  "127.0.0.1:27016",
		},
	}, nil)
	srv := &Service{dao: m}

	Convey("Test AllInfo", t, func() {
		res, err := srv.AllInfo(context.TODO(), &pb.AllInfoReq{})
		Convey("Check it work!", func() {
			So(err, ShouldBeNil)
			So(len(res.Info), ShouldEqual, 2)
			t.Log(res)
		})
	})
}

func TestService_Info(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().InfoWithID(int32(1)).Return(&model.Info{
		ServerID: 1,
		ModID:    1,
		GameID:   1,
		Rcon:     "test",
		Hostname: "test",
		Address:  "127.0.0.1:27015",
	}, nil)
	srv := &Service{dao: m}

	Convey("Test Info", t, func() {
		res, err := srv.Info(context.TODO(), &pb.InfoReq{
			ServerId: 1,
		})
		Convey("Check it work!", func() {
			So(err, ShouldBeNil)
			So(res.Info.ServerId, ShouldEqual, 1)
			t.Log(res)
		})
	})
}

func TestService_Init(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().Info("127.0.0.1:27015").Return(&model.Info{
		ServerID: 1,
		ModID:    1,
		GameID:   1,
		Rcon:     "test",
		Hostname: "test",
		Address:  "127.0.0.1:27015",
	}, nil)
	m.EXPECT().UpdateRcon(gomock.Any()).Return(nil)
	srv := &Service{dao: m}

	Convey("Test Init", t, func() {
		res, err := srv.Init(context.Background(), &pb.InitReq{
			Address: "127.0.0.1",
			Port:    27015,
		})
		Convey("Check it work!", func() {
			So(err, ShouldBeNil)
			So(res.Info.ServerId, ShouldEqual, 1)
			t.Log(res)
		})
	})
}
