package mysql

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func Test_getDSN(t *testing.T) {
	Convey("Call getDSN", t, func() {
		dsn := (&Config{
			Host:     "db",
			User:     "db",
			Password: "db",
			DBName:   "db",
			Charset:  "utf8mb4",
		}).GetDSN()
		Convey("Then dsn should not be empty", func() {
			So(dsn, ShouldNotBeEmpty)
		})
	})
}
