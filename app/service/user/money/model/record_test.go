package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	TestJSON = `[{"timestamp":123,"amount":1,"reason":"test"}]`
)

func TestGetRecordsFromJSON(t *testing.T) {
	Convey("Test GetRecordsFromJSON", t, func() {
		Convey("Check it work", func() {
			res, err := GetRecordsFromJSON([]byte(TestJSON))
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So((*res)[0].Amount, ShouldEqual, 1)
			t.Log((*res)[0])
		})
	})
}

func TestRecordsDB_Add(t *testing.T) {
	var records = &RecordsDB{
		UID:     1,
		Records: []byte(TestJSON),
	}
	Convey("Test RecordsDB Get", t, func() {
		Convey("Check it work", func() {
			err := records.Add(-1, "test")
			So(err, ShouldBeNil)
			t.Log(string(records.Records))
		})
	})
}

func TestRecordsDB_Get(t *testing.T) {
	var records = &RecordsDB{
		UID:     1,
		Records: []byte(TestJSON),
	}
	Convey("Test RecordsDB Get", t, func() {
		Convey("Check it work", func() {
			res, err := records.Get()
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			So((*res)[0].Amount, ShouldEqual, 1)
			t.Log((*res)[0])
		})
	})
}

func TestRecords_Add(t *testing.T) {
	records := &Records{}
	Convey("Test Records Add", t, func() {
		Convey("Check it work", func() {
			records.Add(1, "test")
			So(len(*records), ShouldEqual, 1)
			So((*records)[0].Amount, ShouldEqual, 1)
			t.Log((*records)[0])
		})
	})
}

func TestRecords_RemoveExpr(t *testing.T) {
	records := &Records{
		{
			Timestamp: 123,
			Amount:    1,
			Reason:    "test",
		},
	}
	Convey("Test Records RemoveExpr", t, func() {
		Convey("Check it work", func() {
			records.RemoveExpr()
			So(len(*records), ShouldEqual, 0)
		})
	})
}

func TestRecords_ToJSON(t *testing.T) {
	records := &Records{
		{
			Timestamp: 123,
			Amount:    1,
			Reason:    "test",
		},
	}
	Convey("Test Records ToJSON", t, func() {
		Convey("Check it work", func() {
			res, err := records.ToJSON()
			So(err, ShouldBeNil)
			So(string(res), ShouldEqual, TestJSON)
			t.Log(string(res))
		})
	})
}
