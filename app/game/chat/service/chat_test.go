package service

import (
	pb "backend/app/game/chat/api/grpc/v1"
	"backend/app/game/chat/dao"
	server "backend/app/game/server/api/grpc/v1"
	"backend/pkg/log"
	"context"
	"sync"
	"testing"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestService_AllChat(t *testing.T) {
	log.Init(nil)
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	wg := sync.WaitGroup{}
	wg.Add(1)

	tName := "TEST"
	tMessage := "This is a test message"
	s := server.NewMockServerClient(ctl)
	s.EXPECT().Info(gomock.Any(), gomock.Any()).Return(&server.InfoResp{Info: &server.Info{
		ShortName: "测试服",
	}}, nil)
	d := dao.NewMockDao(ctl)
	d.EXPECT().CreateAllChatEvent(gomock.Any(), gomock.Any()).Return(nil).Do(func(ctx, in interface{}) { wg.Done() })

	srv := &Service{server: s, dao: d}
	Convey("Test AllChat", t, func() {
		Convey("Check it work", func() {
			_, err := srv.AllChat(context.TODO(), &pb.AllChatReq{
				Name:     tName,
				Message:  tMessage,
				ServerId: 1,
			})
			So(err, ShouldBeNil)
		})
	})
	wg.Wait()
}
