package dao

import (
	"backend/app/user/account/model"
	. "backend/app/user/account/test"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

var info = model.Info{UID: TestUID, Username: TestUserName}

func Test_dao_UID(t *testing.T) {
	MockDaoUid()

	// Testing
	Convey("Test with normal data", t, func() {
		res, err := testdao.UID(TestSteamID)
		Convey("Then uid should be 1 and err should be nil", func() {
			So(res.UID, ShouldEqual, TestUID)
			So(err, ShouldBeNil)
		})
	})
	Convey("Test cache", t, func() {
		res, err := testdao.UID(TestSteamID)
		Convey("Then uid should be 1 and err should be nil", func() {
			So(res.UID, ShouldEqual, TestUID)
			So(err, ShouldBeNil)
		})
	})
	Convey("Test error", t, func() {
		_, err := testdao.UID(2)
		Convey("Then error should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
}

func Test_dao_Info(t *testing.T) {
	MockDaoInfo()

	// Testing
	Convey("Test with normal data", t, func() {
		res, err := testdao.Info(TestUID)
		Convey("Then steamid should be 76561198029350216 and err should be nil", func() {
			So(res.SteamID, ShouldEqual, TestSteamID)
			So(err, ShouldBeNil)
		})
	})
	Convey("Test cache", t, func() {
		res, err := testdao.Info(TestUID)
		Convey("Then steamid should be 76561198029350216 and err should be nil", func() {
			So(res.SteamID, ShouldEqual, TestSteamID)
			So(err, ShouldBeNil)
		})
	})
	Convey("Test error", t, func() {
		_, err := testdao.Info(2)
		Convey("Then err should not be nil", func() {
			So(err, ShouldNotBeNil)
		})
	})
}

func Test_dao_Register(t *testing.T) {
	MockDaoRegister()

	// Testing
	Convey("Register with steamID", t, func() {
		Convey("Then uid should be 1 and err should be nil", func() {
			res, err := testdao.Register(TestSteamID)
			So(res.UID, ShouldEqual, TestUID)
			So(err, ShouldBeNil)
		})
	})
}

func Test_dao_ChangeName(t *testing.T) {
	MockDaoChangeName()

	// Testing
	Convey("Change name with steamID", t, func() {
		Convey("Then err should be nil", func() {
			err := testdao.ChangeName(&model.Info{UID: TestUID, Username: TestUserName})
			So(err, ShouldBeNil)
		})
	})
}

func MockDaoUid() {
	rows := testdao.Mock.NewRows([]string{"uid", "steamid", "username", "firstjoin"}).
		AddRow(TestUID, TestSteamID, TestUserName, time.Now().Unix())

	testdao.Mock.ExpectQuery("SELECT (.+) FROM `" + info.TableName() + "`").
		WithArgs(TestSteamID).
		WillReturnRows(rows)
	testdao.Mock.ExpectQuery("SELECT (.+) FROM `" + info.TableName() + "`").
		WithArgs(2).
		WillReturnError(fmt.Errorf("not found"))
}

func MockDaoInfo() {
	rows := testdao.Mock.NewRows([]string{"uid", "steamid", "username", "firstjoin"}).
		AddRow(TestUID, TestSteamID, TestUserName, time.Now().Unix())

	testdao.Mock.ExpectQuery("SELECT (.+) FROM `" + info.TableName() + "`").
		WithArgs(TestUID).
		WillReturnRows(rows)
	testdao.Mock.ExpectQuery("SELECT (.+) FROM `" + info.TableName() + "`").
		WithArgs(2).
		WillReturnError(fmt.Errorf("not found"))
}

func MockDaoRegister() {
	testdao.Mock.ExpectBegin()
	testdao.Mock.ExpectExec("INSERT INTO `"+info.TableName()+"`").
		WithArgs(TestSteamID, "", time.Now().Unix()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	testdao.Mock.ExpectCommit()
}

func MockDaoChangeName() {
	testdao.Mock.ExpectBegin()
	testdao.Mock.ExpectExec("UPDATE `"+info.TableName()+"` SET `username`").
		WithArgs(info.Username, TestUID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	testdao.Mock.ExpectCommit()
}
