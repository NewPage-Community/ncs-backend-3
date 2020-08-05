package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestSign_Sign(t *testing.T) {
	cst := time.FixedZone("CST", 8*3600)
	now := time.Now().In(cst)
	Convey("Test Sign", t, func() {
		Convey("Check continuous sign", func() {
			s := Sign{
				UID:      1,
				SignDate: GetDate(now.Add(-24 * time.Hour)),
				SignDays: 1,
			}
			s.Sign()
			So(s.IsSigned(), ShouldBeTrue)
			So(s.SignDays, ShouldEqual, 2)
			t.Log(s)
		})
		Convey("Check not continuous sign", func() {
			s := Sign{
				UID:      1,
				SignDate: GetDate(now.Add(-48 * time.Hour)),
				SignDays: 1,
			}
			s.Sign()
			So(s.IsSigned(), ShouldBeTrue)
			So(s.SignDays, ShouldEqual, 1)
			t.Log(s)
		})
	})
}

func TestSign_GetNowDate(t *testing.T) {
	Convey("Test GetNowDate", t, func() {
		Convey("Check it work", func() {
			s := Sign{}
			date := s.GetNowDate()
			So(date, ShouldNotBeNil)
			t.Log(date)
		})
	})
}

func TestSign_IsSigned(t *testing.T) {
	cst := time.FixedZone("CST", 8*3600)
	now := time.Now().In(cst)
	Convey("Test IsSigned", t, func() {
		Convey("Check is signed", func() {
			s := Sign{
				SignDate: GetDate(now),
			}
			So(s.IsSigned(), ShouldBeTrue)
		})
		Convey("Check is not signed", func() {
			s := Sign{
				SignDate: GetDate(now.Add(-24 * time.Hour)),
			}
			So(s.IsSigned(), ShouldBeFalse)
		})
	})
}
