package service

import (
	serverService "backend/app/game/server/api/grpc"
	pb "backend/app/service/user/ban/api/grpc"
	"backend/app/service/user/ban/dao"
	"backend/app/service/user/ban/model"
	"backend/pkg/log"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestService_Add(t *testing.T) {
	log.Init(nil)
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	server := serverService.NewMockServerClient(ctl)
	server.EXPECT().RconAll(gomock.Any(), &serverService.RconAllReq{
		Cmd: "ncs_ban_notify 1 1 \"test\"",
	}).Return(&serverService.RconAllResp{
		Success: 0,
	}, nil)
	d := dao.NewMockDao(ctl)
	d.EXPECT().Add(gomock.Any()).Return(nil)

	srv := &Service{
		dao:    d,
		server: server,
	}

	Convey("Test Add", t, func() {
		Convey("Check it work", func() {
			_, err := srv.Add(context.TODO(), &pb.AddReq{
				Info: &pb.Info{
					Id:         1,
					Uid:        1,
					ExpireTime: time.Now().Add(time.Minute).Unix(),
					Type:       1,
					ServerId:   1,
					ModId:      1,
					GameId:     1,
					Reason:     "test",
				},
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestService_BanCheck(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Info(int64(1)).Return(&model.Ban{
		ID:         1,
		UID:        1,
		CreateTime: 1,
		ExpireTime: time.Now().Add(time.Minute).Unix(),
		Type:       1,
		ServerID:   1,
		ModID:      1,
		GameID:     1,
		Reason:     "test",
	}, nil)
	d.EXPECT().Info(int64(2)).Return(&model.Ban{
		UID: 2,
	}, nil)

	srv := &Service{
		dao: d,
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
			})
			So(err, ShouldBeNil)
			So(res.Info.Id, ShouldEqual, 0)
			t.Log(res)
		})
	})
}

func TestService_Info(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().Info(int64(1)).Return(&model.Ban{
		ID:         1,
		UID:        1,
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
