package dao

import (
	cache "backend/pkg/cache/redis"
	db "backend/pkg/database/mysql"
	"github.com/DATA-DOG/go-sqlmock"
)

type mockDao struct {
	dao
	Mock sqlmock.Sqlmock
}

var testdao *mockDao

func init() {
	mock, db := db.InitMock()
	testdao = &mockDao{
		dao: dao{
			db:    db,
			cache: cache.InitMock(),
		},
		Mock: mock,
	}
}
