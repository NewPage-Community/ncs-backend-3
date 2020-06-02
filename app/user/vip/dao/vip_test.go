package dao

import (
	"backend/app/user/vip/model"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/smartystreets/goconvey/convey"
)

func TestAll(t *testing.T) {
	testdata := &model.VIP{
		UID:        1,
		Point:      1,
		ExpireDate: time.Now().UTC().Unix(),
	}

	testdao.Mock.ExpectBegin()
	testdao.Mock.ExpectExec("INSERT").WillReturnResult(sqlmock.NewResult(1, 1))
	testdao.Mock.ExpectCommit()
	testdao.Mock.ExpectQuery("SELECT").WillReturnRows(sqlmock.NewRows([]string{}))
	testdao.Mock.ExpectBegin()
	testdao.Mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	testdao.Mock.ExpectCommit()
	testdao.Mock.ExpectBegin()
	testdao.Mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
	testdao.Mock.ExpectCommit()

	testdao.Register(testdata)
	testdao.Info(testdata)
	testdao.ExpireTime(testdata)
	testdao.Point(testdata)
	testdao.Close()

	Convey("Test all dao of vip", t, func() {
		err := testdao.Mock.ExpectationsWereMet()
		Convey("Make sure everything all right", func() {
			So(err, ShouldBeNil)
		})
	})
}
