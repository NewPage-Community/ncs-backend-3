package jwt

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestJWT(t *testing.T) {
	config := Config{
		ExpireTime: 3600,
		SecretKey:  "abc",
	}
	data := map[string]interface{}{
		"test": "test",
	}
	token, err := config.NewTokenString(data)
	Convey("Test JWT", t, func() {
		Convey("Create a token", func() {
			So(err, ShouldBeNil)
			So(token, ShouldNotBeEmpty)
			t.Log(token)
		})
		Convey("Get payload", func() {
			payload, err := config.GetTokenPayload(token)
			So(err, ShouldBeNil)
			So(payload["test"], ShouldEqual, data["test"])
			t.Log(payload)
		})

	})
}
