package service

import (
	kaiheilaBot "backend/app/bot/kaiheila/api/grpc/v1"
	qqBot "backend/app/bot/qq/api/grpc/v1"
	"backend/app/game/chat"
	pb "backend/app/game/chat/api/grpc/v1"
	"backend/app/game/chat/dao"
	server "backend/app/game/server/api/grpc/v1"
	"backend/pkg/log"
	"context"
	"fmt"
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
	wg.Add(5)

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
	s.EXPECT().RconAll(gomock.Any(), &server.RconAllReq{
		Cmd: fmt.Sprintf("ncs_chat_notify 0 \"%s\" \"%s : %s\"", fmt.Sprintf(AllChatPrefix, " · QQ"), tName, tMessage),
	}).Return(&server.RconAllResp{}, nil)
	kaiheila := kaiheilaBot.NewMockKaiheilaClient(ctl)
	kaiheila.EXPECT().SendChannelMsg(gomock.Any(), gomock.Any()).Return(nil, nil).Do(func(ctx, in interface{}, opts ...interface{}) { wg.Done() })
	qq := qqBot.NewMockQQClient(ctl)
	//qq.EXPECT().SendGroupMessage(gomock.Any(), gomock.Any()).Return(nil, nil)
	d := dao.NewMockDao(ctl)
	d.EXPECT().CreateChatEvent(gomock.Any(), gomock.Any()).Return(nil).Times(4).Do(func(ctx, in interface{}) { wg.Done() })

	srv := &Service{server: s, kaiheila: kaiheila, qq: qq, dao: d}
	Convey("Test AllChat", t, func() {
		Convey("Check it work", func() {
			_, err := srv.AllChat(context.TODO(), &pb.AllChatReq{
				Name:     tName,
				Message:  tMessage,
				ServerId: 1,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check kaiheila chat", func() {
			_, err := srv.AllChat(context.TODO(), &pb.AllChatReq{
				Name:     tName,
				Message:  tMessage,
				ServerId: chat.KaiheilaID,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check discord chat", func() {
			_, err := srv.AllChat(context.TODO(), &pb.AllChatReq{
				Name:     tName,
				Message:  tMessage,
				ServerId: chat.DiscordID,
			})
			So(err, ShouldBeNil)
		})
		Convey("Check qq chat", func() {
			_, err := srv.AllChat(context.TODO(), &pb.AllChatReq{
				Name:     tName,
				Message:  tMessage,
				ServerId: chat.QQID,
			})
			So(err, ShouldBeNil)
		})
	})
	wg.Wait()
}
