package service

import (
	qqBot "backend/app/bot/qq/api/grpc/v1"
	a2sSrv "backend/app/game/a2s/api/grpc/v1"
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
	a2s := a2sSrv.NewMockA2SClient(ctl)
	a2s.EXPECT().A2SQuery(gomock.Any(), &a2sSrv.A2SQueryReq{Address: []string{
		"127.0.0.1:27015",
		"127.0.0.1:27016",
	}}).Return(&a2sSrv.A2SQueryResp{
		Servers: []*a2sSrv.ServerInfo{
			{
				Address: "127.0.0.1:27015",
				Info: &a2sSrv.A2SInfo{
					Hostname: "test",
				},
				Player: make([]*a2sSrv.A2SPlayer, 0),
			},
			{
				Address: "127.0.0.1:27016",
				Info: &a2sSrv.A2SInfo{
					Hostname: "test",
				},
				Player: make([]*a2sSrv.A2SPlayer, 0),
			},
		},
	}, nil)
	srv := &Service{
		dao: m,
		a2s: a2s,
	}

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

func TestService_ChangeMapNotify(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	m := dao.NewMockDao(ctl)
	m.EXPECT().InfoWithID(int32(1)).Return(&model.Info{
		ServerID:  1,
		ModID:     1,
		GameID:    1,
		Rcon:      "test",
		Hostname:  "test",
		Address:   "127.0.0.1:27015",
		ShortName: "test",
	}, nil)
	qq := qqBot.NewMockQQClient(ctl)
	qq.EXPECT().SendGroupMessage(gomock.Any(), &qqBot.SendGroupMessageReq{
		Message:    "test 更换地图 test",
		AutoEscape: false,
	}).Return(nil, nil)

	srv := &Service{dao: m, qq: qq}
	Convey("Test ChangeMapNotify", t, func() {
		_, err := srv.ChangeMapNotify(context.Background(), &pb.ChangeMapNotifyReq{
			ServerId: 1,
			Map:      "test",
		})
		Convey("Check it work!", func() {
			So(err, ShouldBeNil)
		})
	})
}
