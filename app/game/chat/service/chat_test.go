package service

import (
	pb "backend/app/game/chat/api/grpc"
	server "backend/app/game/server/api/grpc"
	"backend/pkg/log"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_AllChat(t *testing.T) {
	log.Init(nil)
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	tName := "TEST"
	tMessage := "This is a test message"
	tShortName := "测试服"
	tPrefix := fmt.Sprintf(AllChatPrefix, " · "+tShortName)
	s := server.NewMockServerClient(ctl)
	s.EXPECT().Info(gomock.Any(), gomock.Any()).Return(&server.InfoResp{Info: &server.Info{
		ShortName: "测试服",
	}}, nil)
	s.EXPECT().RconAll(gomock.Any(), &server.RconAllReq{
		Cmd: fmt.Sprintf("ncs_chat_notify 0 \"%s\" \"%s : %s\"", tPrefix, tName, tMessage),
	}).Return(&server.RconAllResp{}, nil)

	srv := &Service{server: s}
	Convey("Test AllChat", t, func() {
		Convey("Check it work", func() {
			_, err := srv.AllChat(context.TODO(), &pb.AllChatReq{
				Name:    tName,
				Message: tMessage,
			})
			So(err, ShouldBeNil)
		})
	})
}
