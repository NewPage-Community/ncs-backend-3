package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestItems_addItem(t *testing.T) {
	items := &Items{
		{
			ID:       1,
			Amount:   0,
			ExprTime: 1,
		},
		{
			ID:       2,
			Amount:   1,
			ExprTime: 0,
		},
		{
			ID:       3,
			Amount:   0,
			ExprTime: 0,
		},
	}

	Convey("Test addItem", t, func() {
		items.AddItems(items)
		Convey("Then check it work", func() {
			So((*items)[0].ExprTime, ShouldEqual, 2)
			So((*items)[1].Amount, ShouldEqual, 2)
			So((*items)[2].Amount, ShouldEqual, 0)
			So((*items)[2].ExprTime, ShouldEqual, 0)
		})
		items.AddItems(&Items{
			{
				ID:       1,
				Amount:   0,
				ExprTime: 0,
			},
			{
				ID:       2,
				Amount:   0,
				ExprTime: 0,
			},
		})
		Convey("Then check add unlimited item", func() {
			So((*items)[0].Amount, ShouldEqual, 0)
			So((*items)[0].ExprTime, ShouldEqual, 0)
			So((*items)[1].Amount, ShouldEqual, 0)
			So((*items)[1].ExprTime, ShouldEqual, 0)
		})
	})
}
