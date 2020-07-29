package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSign_Sign(t *testing.T) {
	Convey("Test Sign", t, func() {
		Convey("Check it work", func() {
			s := Sign{
				UID:      1,
				SignDays: 1,
			}
			s.Sign()
			So(s.SignTime, ShouldNotBeNil)
			So(s.SignDays, ShouldEqual, 2)
			t.Log(s)
		})
	})
}

func TestSign_GetNowTime(t *testing.T) {
	Convey("Test GetNowTime", t, func() {
		Convey("Check it work", func() {
			s := Sign{}
			So(s.GetNowTime(), ShouldNotBeNil)
			t.Log(s.GetNowTime())
		})
	})
}

func TestSign_IsSigned(t *testing.T) {
	Convey("Test IsSigned", t, func() {
		Convey("Check is signed", func() {
			s := Sign{
				SignTime: time.Now(),
			}
			So(s.IsSigned(), ShouldBeTrue)
		})
		Convey("Check is not signed", func() {
			s := Sign{
				SignTime: time.Now().Add(-24 * time.Hour),
			}
			So(s.IsSigned(), ShouldBeFalse)
		})
	})
}
