package grpc

import (
	. "backend/app/service/user/account/test"
	"context"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGrpc(t *testing.T) {
	ctx := context.Background()
	req := &UIDReq{
		SteamId: TestSteamID,
	}
	ctl := gomock.NewController(t)
	m := NewMockAccountServer(ctl)
	m.EXPECT().UID(gomock.Any(), gomock.Eq(req)).Return(&UIDResp{
		Uid: TestUID,
	}, nil)

	server := InitServer("tcp", "0.0.0.0:2333", m)
	client := InitClient("0.0.0.0:2333")

	Convey("Test grpc client & server", t, func() {
		resp, err := client.UID(ctx, req)
		Convey("Check it work", func() {
			So(resp, ShouldNotBeNil)
			So(err, ShouldBeNil)
		})
		t.Log(resp)
	})

	Close()
	server.Stop()
}
