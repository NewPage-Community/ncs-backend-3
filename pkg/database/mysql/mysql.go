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
	Host     string
	User     string
	Password string
	DBName   string
	Charset  string
	Debug    bool
}

func Init(conf *Config) *gorm.DB {
	level := 3
	if conf.Debug {
		level = 4
	}
	db, err := gorm.Open(mysql.Open(getDSN(conf)), &gorm.Config{
		Logger: logger.New(l.New(os.Stdout, "\r\n", l.LstdFlags), logger.Config{
			SlowThreshold: 100 * time.Millisecond,
			LogLevel:      logger.LogLevel(level),
			Colorful:      true,
		}),
		SkipDefaultTransaction: true,
	})
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
