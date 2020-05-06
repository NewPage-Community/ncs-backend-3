package dao

import (
	cache "backend/pkg/cache/redis"
	db "backend/pkg/database/mysql"
	"github.com/DATA-DOG/go-sqlmock"
)

type MockDao struct {
	dao
	Mock sqlmock.Sqlmock
}

var testdao *MockDao

func init() {
	mock, db := db.InitMock()
	testdao = &MockDao{
		dao: dao{
			db:    db,
			cache: cache.InitMock(),
		},
		Mock: mock,
	}
}
