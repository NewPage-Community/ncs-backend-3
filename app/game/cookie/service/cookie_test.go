package service

import (
	pb "backend/app/game/cookie/api/grpc/v1"
	"backend/app/game/cookie/dao"
	"backend/app/game/cookie/model"
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestService_GetAllCookie(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	cookie := make(map[string]model.CookieModel)
	cookie["test"] = model.CookieModel{
		Value:       "test",
		LastUpdated: 0,
	}

	d := dao.NewMockDao(ctl)
	d.EXPECT().Get(int64(1)).Return(&model.Cookie{
		UID:    1,
		Cookie: cookie,
	}, nil)

	srv := &Service{dao: d}
	Convey("Test GetAllCookie", t, func() {
		Convey("Check it work", func() {
			res, err := srv.GetAllCookie(context.TODO(), &pb.GetAllCookieReq{Uid: 1})
			So(err, ShouldBeNil)
			So(res.Cookie, ShouldNotBeNil)
			t.Log(res)
		})
	})
}

func TestService_GetCookie(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	cookie := make(map[string]model.CookieModel)
	cookie["test"] = model.CookieModel{
		Value:       "test",
		LastUpdated: 0,
	}

	d := dao.NewMockDao(ctl)
	d.EXPECT().Get(int64(1)).Return(&model.Cookie{
		UID:    1,
		Cookie: cookie,
	}, nil).Times(2)

	srv := &Service{dao: d}
	Convey("Test GetCookie", t, func() {
		Convey("Check exist key", func() {
			res, err := srv.GetCookie(context.TODO(), &pb.GetCookieReq{
				Uid: 1,
				Key: "test",
			})
			So(err, ShouldBeNil)
			So(res.Value, ShouldEqual, "test")
			So(res.Exist, ShouldBeTrue)
		})
		Convey("Check not exist key", func() {
			res, err := srv.GetCookie(context.TODO(), &pb.GetCookieReq{
				Uid: 1,
				Key: "test1",
			})
			So(err, ShouldBeNil)
			So(res.Value, ShouldEqual, "")
			So(res.Exist, ShouldBeFalse)
		})
	})
}

func TestService_SetCookie(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	cookie := make(map[string]model.CookieModel)
	cookie["test"] = model.CookieModel{
		Value:       "test",
		LastUpdated: 0,
	}

	time := time.Now().Unix()
	d := dao.NewMockDao(ctl)
	d.EXPECT().Set(int64(1), "test", "test", time).Return(nil)

	srv := &Service{dao: d}
	Convey("Test SetCookie", t, func() {
		Convey("Check it work", func() {
			_, err := srv.SetCookie(context.TODO(), &pb.SetCookieReq{
				Uid:       1,
				Key:       "test",
				Value:     "test",
				Timestamp: time,
			})
			So(err, ShouldBeNil)
		})
	})
}

func TestMergeCookieData(t *testing.T) {
	// Test case 1: valid input
	input1 := `{"key1": {"value": "val1", "last_updated": 1234567890}}`
	expectedOutput1 := &model.Cookie{
		UID: 123,
		Cookie: map[string]model.CookieModel{
			"key1": {
				Value:       "val1",
				LastUpdated: 1234567890,
			},
		},
	}
	c1 := &model.Cookie{UID: 123}
	err1 := c1.MergeCookieData(input1)
	if err1 != nil {
		t.Errorf("Test case 1 failed: expected nil error, but got %v", err1)
	}
	if !reflect.DeepEqual(c1, expectedOutput1) {
		t.Errorf("Test case 1 failed: expected %+v, but got %+v", expectedOutput1, c1)
	}

	// Test case 2: invalid JSON input
	input2 := `{"key1": {"value": "val1", "last_updated": 1234567890}`
	c2 := &model.Cookie{UID: 123}
	err2 := c2.MergeCookieData(input2)
	if err2 == nil {
		t.Errorf("Test case 2 failed: expected non-nil error, but got nil")
	}

	// Test case 3: old format input
	input3 := `{"key1": "val1"}`
	expectedOutput3 := &model.Cookie{
		UID: 123,
		Cookie: map[string]model.CookieModel{
			"key1": {
				Value:       "val1",
				LastUpdated: 0,
			},
		},
	}
	c3 := &model.Cookie{UID: 123}
	err3 := c3.MergeCookieData(input3)
	if err3 != nil {
		t.Errorf("Test case 3 failed: expected nil error, but got %v", err3)
	}
	if !reflect.DeepEqual(c3, expectedOutput3) {
		t.Errorf("Test case 3 failed: expected %+v, but got %+v", expectedOutput3, c3)
	}
}
