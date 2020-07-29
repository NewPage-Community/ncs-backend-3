package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestBan_IsBanned(t *testing.T) {
	Convey("Test IsBanned", t, func() {
		Convey("Check expired ban", func() {
			b := &Ban{
				ID:         1,
				UID:        1,
				CreateTime: 0,
				ExpireTime: 0,
				Type:       0,
				ServerID:   1,
				ModID:      1,
				GameID:     1,
				Reason:     "test",
			}
			So(b.IsBanned(1, 1, 1), ShouldEqual, false)
		})
		Convey("Check BanTypeAll", func() {
			b := &Ban{
				ID:         1,
				UID:        1,
				CreateTime: 0,
				ExpireTime: time.Now().Add(time.Minute).Unix(),
				Type:       BanTypeAll,
				ServerID:   1,
				ModID:      1,
				GameID:     1,
				Reason:     "test",
			}
			So(b.IsBanned(1, 1, 1), ShouldEqual, true)
		})
		Convey("Check BanTypeMod", func() {
			b := &Ban{
				ID:         1,
				UID:        1,
				CreateTime: 0,
				ExpireTime: time.Now().Add(time.Minute).Unix(),
				Type:       BanTypeMod,
				ServerID:   1,
				ModID:      1,
				GameID:     1,
				Reason:     "test",
			}
			So(b.IsBanned(1, 1, 1), ShouldEqual, true)
		})
		Convey("Check BanTypeGame", func() {
			b := &Ban{
				ID:         1,
				UID:        1,
				CreateTime: 0,
				ExpireTime: time.Now().Add(time.Minute).Unix(),
				Type:       BanTypeGame,
				ServerID:   1,
				ModID:      1,
				GameID:     1,
				Reason:     "test",
			}
			So(b.IsBanned(1, 1, 1), ShouldEqual, true)
		})
		Convey("Check BanTypeServer", func() {
			b := &Ban{
				ID:         1,
				UID:        1,
				CreateTime: 0,
				ExpireTime: time.Now().Add(time.Minute).Unix(),
				Type:       BanTypeServer,
				ServerID:   1,
				ModID:      1,
				GameID:     1,
				Reason:     "test",
			}
			So(b.IsBanned(1, 1, 1), ShouldEqual, true)
		})
	})
}

func TestBan_IsExpired(t *testing.T) {
	Convey("Test IsExpired", t, func() {
		Convey("Check expired", func() {
			b := Ban{
				ExpireTime: time.Now().Unix() - 1,
			}
			So(b.IsExpired(), ShouldEqual, true)
		})
		Convey("Check not expired", func() {
			b := Ban{
				ExpireTime: time.Now().Add(time.Second).Unix(),
			}
			So(b.IsExpired(), ShouldEqual, false)
		})
	})
}
