package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGroups_Add(t *testing.T) {
	g := &Groups{gID: []int64{1}}
	Convey("Test Groups add", t, func() {
		g.Add(2)
		success := g.InGroup(2)
		Convey("Add gid:2 to Groups", func() {
			So(success, ShouldEqual, true)
		})
	})
}

func TestGroups_Count(t *testing.T) {
	g := &Groups{gID: []int64{1, 2, 3}}
	Convey("Test Groups count", t, func() {
		count := g.Count()
		Convey("Check it work", func() {
			So(count, ShouldEqual, 3)
		})
	})
}

func TestGroups_GroupsToJSON(t *testing.T) {
	g := &Groups{gID: []int64{1, 2, 3}}
	Convey("Test Groups to JSON", t, func() {
		json, err := g.ToJSON()
		Convey("Check it work", func() {
			So(json, ShouldEqual, "[1,2,3]")
			So(err, ShouldBeNil)
		})
	})
}

func TestGroups_InGroup(t *testing.T) {
	g := &Groups{gID: []int64{1, 2, 3}}
	Convey("Test Groups in group", t, func() {
		in := g.InGroup(2)
		Convey("Check it work", func() {
			So(in, ShouldEqual, true)
		})
		in = g.InGroup(4)
		Convey("Check not in group", func() {
			So(in, ShouldEqual, false)
		})
	})
}

func TestGroups_Sort(t *testing.T) {
	g := &Groups{gID: []int64{3, 2, 1}}
	Convey("Test Groups sort", t, func() {
		g.Sort()
		json, _ := g.ToJSON()
		Convey("Check it work", func() {
			So(json, ShouldEqual, "[1,2,3]")
		})
	})
}

func TestNewGroups(t *testing.T) {
	Convey("Test NewGroups", t, func() {
		g, err := NewGroups("[3,2,1]")
		json, err := g.ToJSON()
		Convey("Check it work", func() {
			So(json, ShouldEqual, "[1,2,3]")
			So(err, ShouldBeNil)
		})
	})
}
