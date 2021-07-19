package jwt

import (
	"context"
	. "github.com/smartystreets/goconvey/convey"
	"google.golang.org/grpc/metadata"
	"testing"
)

func TestJWT(t *testing.T) {
	jwt := JWT{
		ExpireTime: 3600,
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
			unaryHandler := func(ctx context.Context, req interface{}) (interface{}, error) {
				return PayloadFormContext(ctx), nil
			}
			md := metadata.MD{}
			md.Set(authorizationHeader, bearerPrefix+token)
			ctx := metadata.NewIncomingContext(context.Background(), md)
			payload, err := jwt.GetPayload(ctx, nil, nil, unaryHandler)
			So(err, ShouldBeNil)
			So((*payload.(*Payload))["test"], ShouldEqual, data["test"])
			t.Log(payload)
		})
	})
}
