package service

import (
	serverService "backend/app/game/server/api/grpc/v1"
	accountService "backend/app/service/user/account/api/grpc/v1"
	pb "backend/app/service/user/ban/api/grpc/v1"
	"backend/app/service/user/ban/dao"
	"backend/app/service/user/ban/model"
	"backend/pkg/log"
	"backend/pkg/steam"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"testing"
	"time"
)

func TestService_Add(t *testing.T) {
	log.Init(nil)
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	now := time.Now()

	server := serverService.NewMockServerClient(ctl)
	server.EXPECT().RconAll(gomock.Any(), &serverService.RconAllReq{
		Cmd: "ncs_ban_notify 1 1 \"test\"",
	}).Return(&serverService.RconAllResp{
		Success: 0,
	}, nil).Times(3)

	d := dao.NewMockDao(ctl)
	d.EXPECT().Add(&model.Ban{
		UID:        1,
		IP:         "127.0.0.1",
		CreateTime: now.Unix(),
		ExpireTime: now.Add(time.Minute).Unix(),
		Type:       1,
		Reason:     "test",
	}).Return(nil).Times(3)
	d.EXPECT().Add(&model.Ban{
		UID:        2,
		CreateTime: now.Unix(),
		ExpireTime: now.Add(time.Minute).Unix(),
		Type:       1,
		Reason:     "test" + LibOwnerBanReason,
	}).Return(nil)
	d.EXPECT().Add(&model.Ban{
		UID:        3,
		CreateTime: now.Unix(),
		ExpireTime: now.Add(time.Minute).Unix(),
		Type:       1,
		Reason:     "test" + LibOwnerBanReason,
	}).Return(nil)

	acc := accountService.NewMockAccountClient(ctl)
	acc.EXPECT().Info(gomock.Any(), &accountService.InfoReq{Uid: 1}).
		Return(&accountService.InfoResp{Info: &accountService.Info{
			SteamId: 1,
		}}, nil).Times(2)
	acc.EXPECT().UID(gomock.Any(), &accountService.UIDReq{SteamId: 2}).
		Return(&accountService.UIDResp{Uid: 2}, nil)
	acc.EXPECT().UID(gomock.Any(), &accountService.UIDReq{SteamId: 3}).
		Return(
			&accountService.UIDResp{},
			status.Error(codes.NotFound, ""),
		)
	acc.EXPECT().Register(gomock.Any(), &accountService.RegisterReq{
		SteamId: 3,
	}).Return(&accountService.RegisterResp{Uid: 3}, nil)

	steam2 := steam.NewMockAPIClient(ctl)
	steam2.EXPECT().IsPlayingSharedGame(uint64(1), 1).
		Return(steam.PlayingSharedGame{LenderSteamID: 2}, nil)
	steam2.EXPECT().IsPlayingSharedGame(uint64(1), 2).
		Return(steam.PlayingSharedGame{LenderSteamID: 3}, nil)

	srv := &Service{
		dao:     d,
		server:  server,
		account: acc,
		steam:   steam2,
	}

	Convey("Test Add", t, func() {
		Convey("Check it work", func() {
			_, err := srv.Add(context.TODO(), &pb.AddReq{
				Info: &pb.Info{
					Uid:        1,
					Ip:         "127.0.0.1",
					ExpireTime: now.Add(time.Minute).Unix(),
					Type:       1,
					Reason:     "test",
				},
			})
			So(err, ShouldBeNil)
		})
		Convey("add ban with lib owner", func() {
			_, err := srv.Add(context.TODO(), &pb.AddReq{
				Info: &pb.Info{
					Uid:        1,
					Ip:         "127.0.0.1",
					ExpireTime: now.Add(time.Minute).Unix(),
					Type:       1,
					Reason:     "test",
					AppId:      1,
				},
			})
			So(err, ShouldBeNil)
		})
		Convey("add ban with lib owner but owner do not have account", func() {
			_, err := srv.Add(context.TODO(), &pb.AddReq{
				Info: &pb.Info{
					Uid:        1,
					Ip:         "127.0.0.1",
					ExpireTime: now.Add(time.Minute).Unix(),
					Type:       1,
					Reason:     "test",
					AppId:      2,
				},
			})
			So(err, ShouldBeNil)
		})
	})
	// Wait routine
	time.Sleep(time.Second)
}

func TestService_BanCheck(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Info(uint64(1)).Return(&model.Ban{
		ID:         1,
		UID:        1,
		ExpireTime: time.Now().Add(time.Minute).Unix(),
		Type:       1,
		Reason:     "test",
	}, nil)
	d.EXPECT().Info(uint64(2)).Return(&model.Ban{
		UID: 2,
	}, nil).Times(3)
	//d.EXPECT().IsBlockIP("127.0.0.1").Return(true, nil)
	d.EXPECT().Info(uint64(3)).Return(&model.Ban{
		ID:         1,
		UID:        3,
		ExpireTime: time.Now().Add(time.Minute).Unix(),
		Type:       1,
		Reason:     "test",
	}, nil)
	d.EXPECT().Info(uint64(4)).Return(&model.Ban{}, nil)
	d.EXPECT().Add(gomock.Any()).DoAndReturn(func(info *model.Ban) error {
		info.ID = 2
		return nil
	})

	acc := accountService.NewMockAccountClient(ctl)
	acc.EXPECT().Info(gomock.Any(), &accountService.InfoReq{Uid: 2}).
		Return(&accountService.InfoResp{
			Info: &accountService.Info{
				SteamId: 2,
			},
		}, nil).Times(3)
	acc.EXPECT().UID(gomock.Any(), &accountService.UIDReq{SteamId: 3}).
		Return(&accountService.UIDResp{Uid: 3}, nil)
	acc.EXPECT().UID(gomock.Any(), &accountService.UIDReq{SteamId: 4}).
		Return(&accountService.UIDResp{Uid: 4}, nil)

	steam2 := steam.NewMockAPIClient(ctl)
	steam2.EXPECT().IsPlayingSharedGame(uint64(2), 1).
		Return(steam.PlayingSharedGame{LenderSteamID: 3}, nil)
	steam2.EXPECT().IsPlayingSharedGame(uint64(2), 2).
		Return(steam.PlayingSharedGame{LenderSteamID: 0}, nil).Times(1)
	steam2.EXPECT().IsPlayingSharedGame(uint64(2), 3).
		Return(steam.PlayingSharedGame{LenderSteamID: 4}, nil)

	server := serverService.NewMockServerClient(ctl)
	server.EXPECT().RconAll(gomock.Any(), &serverService.RconAllReq{
		Cmd: "ncs_ban_notify 2 1 \"test(共享者账号被封禁)\"",
	}).Return(&serverService.RconAllResp{
		Success: 0,
	}, nil)

	srv := &Service{
		dao:     d,
		account: acc,
		steam:   steam2,
		server:  server,
	}

	Convey("Test BanCheck", t, func() {
		Convey("Check it work", func() {
			res, err := srv.BanCheck(context.TODO(), &pb.Info2Req{
				Uid:      1,
				ServerId: 1,
				ModId:    1,
				GameId:   1,
			})
			So(err, ShouldBeNil)
			So(res.Info.Id, ShouldEqual, 1)
			t.Log(res)
		})
		Convey("Check record not found", func() {
			res, err := srv.BanCheck(context.TODO(), &pb.Info2Req{
				Uid:      2,
				ServerId: 1,
				ModId:    1,
				GameId:   1,
				AppId:    2,
			})
			So(err, ShouldBeNil)
			So(res.Info.Id, ShouldEqual, 0)
			t.Log(res)
		})
		/*Convey("Check record not found but block ip", func() {
			res, err := srv.BanCheck(context.TODO(), &pb.Info2Req{
				Uid:      2,
				Ip:       "127.0.0.1",
				ServerId: 1,
				ModId:    1,
				GameId:   1,
				AppId:    2,
			})
			So(err, ShouldBeNil)
			So(res.Info.Id, ShouldEqual, 0)
			So(res.Info.BlockIp, ShouldBeTrue)
			t.Log(res)
		})*/
		Convey("Check shared lib", func() {
			res, err := srv.BanCheck(context.TODO(), &pb.Info2Req{
				Uid:      2,
				ServerId: 1,
				ModId:    1,
				GameId:   1,
				AppId:    3,
			})
			So(err, ShouldBeNil)
			So(res.Info.ExpireTime, ShouldEqual, 0)
			t.Log(res)
		})
		Convey("Check shared lib but owner got ban", func() {
			res, err := srv.BanCheck(context.TODO(), &pb.Info2Req{
				Uid:      2,
				ServerId: 1,
				ModId:    1,
				GameId:   1,
				AppId:    1,
			})
			So(err, ShouldBeNil)
			t.Log(res)
		})
	})
	// Wait routine
	time.Sleep(time.Second)
}

func TestService_Info(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Info(uint64(1)).Return(&model.Ban{
		ID:         1,
		UID:        1,
		IP:         "127.0.0.1",
		CreateTime: 1,
		ExpireTime: time.Now().Add(time.Minute).Unix(),
		Type:       1,
		ServerID:   1,
		ModID:      1,
		GameID:     1,
		Reason:     "test",
	}, nil)

	srv := &Service{dao: d}

	Convey("Test Info", t, func() {
		Convey("Check it work", func() {
			res, err := srv.Info(context.TODO(), &pb.InfoReq{
				Uid: 1,
			})
			So(err, ShouldBeNil)
			So(res.Info.Id, ShouldEqual, 1)
			t.Log(res)
		})
	})
}

func TestService_Remove(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Remove(&model.Ban{ID: 1}).Return(nil)

	srv := &Service{dao: d}

	Convey("Test Remove", t, func() {
		Convey("Check it work", func() {
			_, err := srv.Remove(context.TODO(), &pb.RemoveReq{
				Id: 1,
			})
			So(err, ShouldBeNil)
		})
	})
}
