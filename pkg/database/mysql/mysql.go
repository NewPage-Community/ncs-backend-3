package mysql

import (
	"backend/pkg/log"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Charset  string
}

func Init(conf *Config) *gorm.DB {
	db, err := gorm.Open(mysql.Open(getDSN(conf)), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db
}

func getDSN(conf *Config) string {
	return fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DBName,
		conf.Charset,
	)
}

func InitMock() (sqlmock.Sqlmock, *gorm.DB) {
	test, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	db, err := gorm.Open(mysql.Open(""), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic(err)
	}
	db.ConnPool = test
	return mock, db
}

func Healthy(db *gorm.DB) bool {
	if db == nil {
		return false
	}
	d, err := db.DB()
	if err != nil {
		log.Error(err)
		return false
	}
	if err = d.Ping(); err != nil {
		log.Error(err)
		return false
	}
	return true
}
