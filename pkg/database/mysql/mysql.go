package mysql

import (
	"backend/pkg/log"
	"fmt"
	l "log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	Host         string
	User         string
	Password     string
	DBName       string
	Charset      string
	Debug        bool
	MaxOpenConns int
	MaxIdleConns int
}

var client *gorm.DB

func (conf *Config) GetDSN() string {
	return fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=%s&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DBName,
		conf.Charset,
	)
}

func (conf *Config) Init() {
	if conf.Host == "" {
		conf.Host = "127.0.0.1"
	}
	if conf.User == "" {
		conf.User = "root"
	}
	if conf.Password == "" {
		conf.Password = "password"
	}
	if conf.DBName == "" {
		conf.DBName = "mysql"
	}
	if conf.Charset == "" {
		conf.Charset = "utf8mb4"
	}
}

func (conf *Config) GetLogger() logger.Interface {
	logLevel := logger.Silent
	if conf.Debug {
		logLevel = logger.Info
	}
	return logger.New(l.New(os.Stdout, "\r\n", l.LstdFlags), logger.Config{
		SlowThreshold: 100 * time.Millisecond,
		LogLevel:      logLevel,
		Colorful:      true,
	})
}

func Init(conf *Config) *gorm.DB {
	var err error
	conf.Init()
	conn := func() (*gorm.DB, error) {
		return gorm.Open(mysql.Open(conf.GetDSN()), &gorm.Config{
			Logger:                 conf.GetLogger(),
			SkipDefaultTransaction: true,
			PrepareStmt:            true,
		})
	}

	// Retry
	for i := 0; i < 3; i++ {
		client, err = conn()
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	// Set conns pool
	if err == nil {
		sqlDB, err := client.DB()
		if err == nil {
			sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
			sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		}
	}

	if err != nil {
		panic(err)
	}
	return client
}

func Healthy() bool {
	if client != nil {
		db, err := client.DB()
		if err == nil {
			err = db.Ping()
			if err == nil {
				return true
			}
		}
		log.Error(err)
	}
	return true
}

func Close() {

}
