package service

import (
	kaiheilaBot "backend/app/bot/kaiheila/api/grpc/v1"
	pb "backend/app/game/chat/api/grpc"
	server "backend/app/game/server/api/grpc"
	"backend/pkg/log"
	"context"
	"fmt"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestService_AllChat(t *testing.T) {
	log.Init(nil)
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	tName := "TEST"
	tMessage := "This is a test message"
	s := server.NewMockServerClient(ctl)
	s.EXPECT().Info(gomock.Any(), gomock.Any()).Return(&server.InfoResp{Info: &server.Info{
		ShortName: "测试服",
	}}, nil)
	s.EXPECT().RconAll(gomock.Any(), &server.RconAllReq{
		Cmd: fmt.Sprintf("ncs_chat_notify 0 \"%s\" \"%s : %s\"", fmt.Sprintf(AllChatPrefix, " · 测试服"), tName, tMessage),
	}).Return(&server.RconAllResp{}, nil)
	s.EXPECT().RconAll(gomock.Any(), &server.RconAllReq{
		Cmd: fmt.Sprintf("ncs_chat_notify 0 \"%s\" \"%s : %s\"", fmt.Sprintf(AllChatPrefix, " · 开黑啦"), tName, tMessage),
	}).Return(&server.RconAllResp{}, nil)
	s.EXPECT().RconAll(gomock.Any(), &server.RconAllReq{
		Cmd: fmt.Sprintf("ncs_chat_notify 0 \"%s\" \"%s : %s\"", fmt.Sprintf(AllChatPrefix, " · Discord"), tName, tMessage),
	}).Return(&server.RconAllResp{}, nil)
	kaiheila := kaiheilaBot.NewMockKaiheilaClient(ctl)
	kaiheila.EXPECT().SendChannelMsg(gomock.Any(), gomock.Any()).Return(nil, nil).Times(2)

	srv := &Service{server: s, kaiheila: kaiheila}
	Convey("Test AllChat", t, func() {
		Convey("Check it work", func() {
			_, err := srv.AllChat(context.TODO(), &pb.AllChatReq{
				Name:    tName,
				Message: tMessage,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check kaiheila chat", func() {
			_, err := srv.AllChat(context.TODO(), &pb.AllChatReq{
				Name:     tName,
				Message:  tMessage,
				ServerId: -1,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check discord chat", func() {
			_, err := srv.AllChat(context.TODO(), &pb.AllChatReq{
				Name:     tName,
				Message:  tMessage,
				ServerId: -100,
			})
			So(err, ShouldBeNil)
		})
	})
	time.Sleep(time.Second)
}
