package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestItems_AddItems(t *testing.T) {
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

	Convey("Test AddItems", t, func() {
		Convey("Check it work", func() {
			items.AddItems(items)
			So((*items)[0].ExprTime, ShouldEqual, 2)
			So((*items)[1].Amount, ShouldEqual, 2)
			So((*items)[2].Amount, ShouldEqual, 0)
			So((*items)[2].ExprTime, ShouldEqual, 0)
		})
		Convey("Check add unlimited item", func() {
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
			So((*items)[0].Amount, ShouldEqual, 0)
			So((*items)[0].ExprTime, ShouldEqual, 0)
			So((*items)[1].Amount, ShouldEqual, 0)
			So((*items)[1].ExprTime, ShouldEqual, 0)
		})
		Convey("Check add not exist item", func() {
			items.AddItems(&Items{
				{
					ID:       4,
					Amount:   0,
					ExprTime: 1,
				},
			})
			So((*items)[3].ID, ShouldEqual, 4)
			So((*items)[3].Amount, ShouldEqual, 0)
			So((*items)[3].ExprTime, ShouldBeGreaterThanOrEqualTo, time.Now().Unix())
			t.Log(time.Now().Unix(), (*items)[3])
		})
		Convey("Check add item already is unlimited 1", func() {
			items.AddItems(&Items{
				{
					ID:       3,
					Amount:   1,
					ExprTime: 0,
				},
			})
			So((*items)[2].Amount, ShouldEqual, 0)
			So((*items)[2].ExprTime, ShouldEqual, 0)
		})
		Convey("Check add item already is unlimited 2", func() {
			items.AddItems(&Items{
				{
					ID:       3,
					Amount:   0,
					ExprTime: 1,
				},
			})
			So((*items)[2].Amount, ShouldEqual, 0)
			So((*items)[2].ExprTime, ShouldEqual, 0)
		})
	})
}

func BenchmarkItems_AddItems(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		items := &Items{}
		items.AddItems(&Items{
			{
				ID: 1,
			},
		})
	}
}

func BenchmarkItems_RemoveItem(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		items := &Items{
			{
				ID: 1,
			},
		}
		items.RemoveItem(Item{
			ID: 1,
		}, true)
	}
}

func BenchmarkItems_Check(b *testing.B) {
	t := time.Now().Add(time.Hour).Unix()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		items := &Items{
			{
				ID: 1,
			},
			{
				ID:       2,
				ExprTime: 1,
			},
			{
				ID:       3,
				ExprTime: t,
			},
		}
		items.Check()
	}
}
