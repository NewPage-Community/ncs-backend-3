package model

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestUserModel_GetUser(t *testing.T) {
	user := &UserModel{UID: 1}
	Convey("Test UserModel GetUser", t, func() {
		Convey("Check it work", func() {
			res, err := user.GetUser()
			So(err, ShouldBeNil)
			So(res, ShouldNotBeNil)
			t.Logf("%v", res)
		})
	})
}

func TestUser_AddItems(t *testing.T) {
	user := &User{UID: 1}
	Convey("Test User AddItems", t, func() {
		Convey("Check it work", func() {
			user.AddItems(&Items{
				{
					ID:       2,
					Amount:   0,
					ExprTime: 0,
				},
				{
					ID:       1,
					Amount:   0,
					ExprTime: 0,
				},
			})
			So(len(*user.Items), ShouldEqual, 2)
			t.Logf("%v", user.Items)
		})
	})
}

func TestUser_CheckItem(t *testing.T) {
	user := &User{
		UID: 1,
	}
	user.AddItems(&Items{
		{
			ID:       1,
			Amount:   0,
			ExprTime: 0,
		},
		{
			ID:       2,
			Amount:   0,
			ExprTime: 1,
		},
	})

	Convey("Test User CheckItem", t, func() {
		Convey("Check last item expired", func() {
			time.Sleep(2 * time.Second)
			user.CheckItem()
			So(len(*user.Items), ShouldEqual, 1)
		})
	})
}

func TestUser_GetModel(t *testing.T) {
	user := &User{
		UID: 1,
	}
	user.AddItems(&Items{
		{
			ID:       1,
			Amount:   0,
			ExprTime: 0,
		},
	})
	Convey("Test User GetModel", t, func() {
		Convey("Check it work", func() {
			res, err := user.GetModel()
			So(err, ShouldBeNil)
			So(res.UID, ShouldEqual, 1)
			So(string(res.Items), ShouldEqual, `{"1":{"id":1,"amount":0,"expr_time":0}}`)
		})
	})
}

func TestUser_RemoveItem(t *testing.T) {
	Convey("Test User RemoveItem", t, func() {
		Convey("Check remove ones", func() {
			user := &User{
				UID: 1,
			}
			user.AddItems(&Items{
				{
					ID:       1,
					Amount:   2,
					ExprTime: 0,
				},
			})
			user.RemoveItem(Item{ID: 1}, false)
			So((*user.Items)[1].Amount, ShouldEqual, 1)
		})
		Convey("Check remove all", func() {
			user := &User{
				UID: 1,
			}
			user.AddItems(&Items{
				{
					ID:       1,
					Amount:   2,
					ExprTime: 0,
				},
			})
			user.RemoveItem(Item{ID: 1}, true)
			So(len(*user.Items), ShouldEqual, 0)
		})
	})
}

func TestUser_SearchItem(t *testing.T) {
	user := &User{
		UID: 1,
	}
	user.AddItems(&Items{
		{
			ID:       1,
			Amount:   2,
			ExprTime: 0,
		},
		{
			ID:       2,
			Amount:   0,
			ExprTime: 0,
		},
	})
	Convey("Test User SearchItem", t, func() {
		Convey("Check it work", func() {
			res, ok := user.SearchItem(2)
			So(ok, ShouldBeTrue)
			So(res.ID, ShouldEqual, 2)
		})
		Convey("Check not found", func() {
			res, ok := user.SearchItem(3)
			So(ok, ShouldBeFalse)
			So(res.ID, ShouldEqual, 0)
		})
	})
}
