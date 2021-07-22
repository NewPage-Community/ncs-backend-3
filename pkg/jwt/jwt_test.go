package jwt

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestJWT(t *testing.T) {
	jwt := JWT{
		ExpireTime: 720,
		SecretKey:  "key",
	}
	data := map[string]interface{}{
		"test": "test",
	}
	token, err := jwt.NewTokenString(data)
	Convey("Test JWT", t, func() {
		Convey("Create a token", func() {
			So(err, ShouldBeNil)
			So(token, ShouldNotBeEmpty)
			t.Log(token)
		})
		Convey("Get payload", func() {
			payload, err := jwt.GetTokenPayload(token)
			So(err, ShouldBeNil)
			So(payload["test"], ShouldEqual, data["test"])
			t.Log(payload)
		})
		Convey("GRPC interceptor", func() {
			md := metadata.MD{}
			md.Set(authorizationHeader, bearerPrefix+token)
			ctx := metadata.NewIncomingContext(context.Background(), md)
			payload := jwt.PayloadFormContext(ctx)
			So(err, ShouldBeNil)
			So(payload.Get("test"), ShouldEqual, data["test"])
			t.Log(payload)
		})
	})
}
