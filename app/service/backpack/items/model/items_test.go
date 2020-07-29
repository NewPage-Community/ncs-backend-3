package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAttributes_GetSet(t *testing.T) {
	var attr = &Attributes{}
	attr.Set("test", "test")
	Convey("Test Attributes Get", t, func() {
		Convey("Check it work", func() {
			res, ok := attr.Get("test")
			So(ok, ShouldBeTrue)
			So(res, ShouldEqual, "test")
		})
	})
}

func TestItem_GetAttributes(t *testing.T) {
	item := &Item{
		ID: 1,
	}
	Convey("Test Get Attributes", t, func() {
		Convey("Check it work", func() {
			attr, err := item.GetAttributes()
			So(err, ShouldBeNil)
			So(attr, ShouldNotBeNil)
		})
	})
}

func TestItem_SetAttributes(t *testing.T) {
	var attr = &Attributes{}
	attr.Set("test", "test")
	item := &Item{
		ID: 1,
	}
	Convey("Test Set Attributes", t, func() {
		Convey("Check it work", func() {
			err := item.SetAttributes(attr)
			So(err, ShouldBeNil)
			So(item.Attributes, ShouldNotBeEmpty)
			t.Log(string(item.Attributes))
		})
	})
}
