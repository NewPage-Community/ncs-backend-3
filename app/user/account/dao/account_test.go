package dao

import (
	"backend/app/user/account/model"
	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

const (
	TestSteamID  = int64(76561198029350216)
	TestUID      = int64(1)
	TestUserName = "Gunslinger"
)

var info = model.Info{UID: TestUID, Username: TestUserName}

func Test_dao_UID(t *testing.T) {
	Mock_dao_UID()

	// Testing
	Convey("Get uid with steamID", t, func() {
		res, err := testdao.UID(TestSteamID)
		Convey("Then uid should be 1 and err should be nil", func() {
			So(res.UID, ShouldEqual, TestUID)
			So(err, ShouldBeNil)
		})
		// Cache
		res, err = testdao.UID(TestSteamID)
		Convey("It should be same", func() {
			So(res.UID, ShouldEqual, TestUID)
			So(err, ShouldBeNil)
		})
		So(testdao.Mock.ExpectationsWereMet(), ShouldBeNil)
	})
}

func Test_dao_Info(t *testing.T) {
	Mock_dao_Info()

	// Testing
	Convey("Get info with uid", t, func() {
		res, err := testdao.Info(TestUID)
		Convey("Then steamid should be 76561198029350216 and err should be nil", func() {
			So(res.SteamID, ShouldEqual, TestSteamID)
			So(err, ShouldBeNil)
		})
		// Cache
		res, err = testdao.Info(TestUID)
		Convey("It should be same", func() {
			So(res.SteamID, ShouldEqual, TestSteamID)
			So(err, ShouldBeNil)
		})
		So(testdao.Mock.ExpectationsWereMet(), ShouldBeNil)
	})
}

func Test_dao_Register(t *testing.T) {
	Mock_dao_Register()

	// Testing
	Convey("Register with steamID", t, func() {
		res, err := testdao.Register(TestSteamID)
		Convey("Then uid should be 1 and err should be nil", func() {
			So(res.UID, ShouldEqual, TestUID)
			So(err, ShouldBeNil)
		})
		So(testdao.Mock.ExpectationsWereMet(), ShouldBeNil)
	})
}

func Test_dao_ChangeName(t *testing.T) {
	Mock_dao_ChangeName()

	// Testing
	Convey("Change name with steamID", t, func() {
		err := testdao.ChangeName(&model.Info{UID: TestUID, Username: TestUserName})
		Convey("Then err should be nil", func() {
			So(err, ShouldBeNil)
		})
		So(testdao.Mock.ExpectationsWereMet(), ShouldBeNil)
	})
}

func Mock_dao_UID() {
	rows := testdao.Mock.NewRows([]string{"uid", "steamid", "username", "firstjoin"}).
		AddRow(TestUID, TestSteamID, TestUserName, time.Now().Unix())

	testdao.Mock.ExpectQuery("SELECT (.+) FROM `" + info.TableName() + "`").
		WithArgs(TestSteamID).
		WillReturnRows(rows)
}

func Mock_dao_Info() {
	rows := testdao.Mock.NewRows([]string{"uid", "steamid", "username", "firstjoin"}).
		AddRow(TestUID, TestSteamID, TestUserName, time.Now().Unix())

	testdao.Mock.ExpectQuery("SELECT (.+) FROM `" + info.TableName() + "`").
		WithArgs(TestUID).
		WillReturnRows(rows)
}

func Mock_dao_Register() {
	testdao.Mock.ExpectBegin()
	testdao.Mock.ExpectExec("INSERT INTO `"+info.TableName()+"`").
		WithArgs(TestSteamID, "", time.Now().Unix()).
		WillReturnResult(sqlmock.NewResult(1, 1))
	testdao.Mock.ExpectCommit()
}

func Mock_dao_ChangeName() {
	testdao.Mock.ExpectBegin()
	testdao.Mock.ExpectExec("UPDATE `"+info.TableName()+"` SET `username`").
		WithArgs(info.Username, TestUID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	testdao.Mock.ExpectCommit()
}
