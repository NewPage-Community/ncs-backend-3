package grpc

import (
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGrpc(t *testing.T) {
	ctx := context.Background()
	ctl := gomock.NewController(t)
	m := NewMockVIPServer(ctl)
	m.EXPECT().Info(gomock.Any(), gomock.Eq(&InfoReq{Uid: 1})).Return(&InfoResp{
		Info: &Info{
			Point:      1,
			Level:      1,
			ExpireDate: time.Now().Unix(),
		},
	}, nil)

	server := InitServer("tcp", "0.0.0.0:2333", m)
	client := InitClient("0.0.0.0:2333")

	Convey("Test grpc client & server", t, func() {
		resp, err := client.Info(ctx, &InfoReq{Uid: 1})
		Convey("Check it work", func() {
			So(resp, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
		t.Log(resp)
	})

	Close()
	server.Stop()
}
