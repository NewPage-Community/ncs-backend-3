package mysql

import (
	"backend/pkg/log"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	l "log"
	"os"
	"time"
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

func Init(conf *Config) (db *gorm.DB) {
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
		db, err = conn()
		if err == nil {
			break
		}
		time.Sleep(5 * time.Second)
	}

	// Set conns pool
	if err == nil {
		sqlDB, err := db.DB()
		if err == nil {
			sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
			sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		}
	}

	if err != nil {
		panic(err)
	}
	return
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
