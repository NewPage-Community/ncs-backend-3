package service

import (
	"backend/app/service/pass/reward/model"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestService_getRewards(t *testing.T) {
	items := []model.Item{
		{
			Level:  1,
			ID:     1,
			Amount: 1,
			Length: 0,
		},
		{
			Level:  2,
			ID:     2,
			Amount: 0,
			Length: 1,
		},
	}
	srv := &Service{
		reward: &model.Reward{
			Season:      1,
			MaxLevel:    2,
			FreeRewards: items,
			AdvRewards:  items,
		}}
	Convey("Test getRewards", t, func() {
		Convey("Check AdvRewards", func() {
			res := srv.getRewards(1, 1, 0)
			So(len(res), ShouldEqual, 1)
			t.Log(res)
			res = srv.getRewards(1, 1, 1)
			So(len(res), ShouldEqual, 1)
			t.Log(res)
			res = srv.getRewards(1, 2, 1)
			So(len(res), ShouldEqual, 2)
			t.Log(res)
			res = srv.getRewards(1, 2, 0)
			So(len(res), ShouldEqual, 1)
			t.Log(res)
			res = srv.getRewards(1, 3, 0)
			So(len(res), ShouldEqual, 1)
			t.Log(res)
		})
		Convey("Check FreeRewards", func() {
			res := srv.getRewards(0, 1, 0)
			So(len(res), ShouldEqual, 1)
			t.Log(res)
			res = srv.getRewards(0, 1, 1)
			So(len(res), ShouldEqual, 1)
			t.Log(res)
			res = srv.getRewards(0, 2, 1)
			So(len(res), ShouldEqual, 2)
			t.Log(res)
			res = srv.getRewards(0, 2, 0)
			So(len(res), ShouldEqual, 1)
			t.Log(res)
		})
	})
}
