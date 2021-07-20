package service

import (
	pb "backend/app/game/cvar/api/grpc/v1"
	"backend/app/game/cvar/dao"
	"backend/app/game/cvar/model"
	serverService "backend/app/game/server/api/grpc/v1"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/robfig/cron/v3"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestService_CheckCVars(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().GetCVars().Return([]*model.CVar{
		{
			ID:         1,
			Key:        "test",
			Value:      "test",
			Type:       0,
			TypeArea:   0,
			NeedUpdate: true,
		},
	}, nil)
	d.EXPECT().UpdatedCVar(1).Return(nil)

	server := serverService.NewMockServerClient(ctl)
	server.EXPECT().AllInfo(gomock.Any(), &serverService.AllInfoReq{}).Return(
		&serverService.AllInfoResp{
			Info: []*serverService.Info{
				{
					ServerId: 1,
					ModId:    1,
					GameId:   1,
				},
			}}, nil)
	server.EXPECT().Rcon(gomock.Any(), gomock.Any()).Return(nil, nil)

	srv := &Service{dao: d, server: server}

	Convey("Test CheckCVars", t, func() {
		Convey("Check it work", func() {
			srv.CheckCVars()
			// Wait for Goroutines
			time.Sleep(time.Second)
		})
	})
}

func TestService_UpdateCVars(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	d := dao.NewMockDao(ctl)
	d.EXPECT().GetCVars().Return([]*model.CVar{
		{
			ID:    1,
			Key:   "test",
			Value: "test",
		},
	}, nil)

	srv := &Service{dao: d}

	Convey("Test UpdateCVars", t, func() {
		Convey("Check it work", func() {
			res, err := srv.UpdateCVars(context.TODO(), &pb.UpdateCVarsReq{
				GameId:   1,
				ModId:    1,
				ServerId: 1,
			})
			So(err, ShouldBeNil)
			So(len(res.Cvars), ShouldEqual, 1)
			t.Log(res.Cvars)
		})
	})
}

func TestService_regCron(t *testing.T) {
	srv := &Service{cron: cron.New()}
	Convey("Test regCron", t, func() {
		Convey("Check it work", func() {
			srv.regCron()
		})
	})
}
