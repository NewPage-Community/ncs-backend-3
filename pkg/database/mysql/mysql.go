package mysql

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Config struct {
	Host     string
	User     string
	Password string
	DBName   string
	Charset  string
}

func Init(conf *Config) *gorm.DB {
	db, err := gorm.Open("mysql", getDSN(conf))
	if err != nil {
		panic(err)
	}
	err = db.DB().Ping()
	if err != nil {
		panic(err)
	}
	return db
}

func getDSN(conf *Config) string {
	return fmt.Sprintf(
		"%s:%s@%s/%s?charset=%s&parseTime=True&loc=Local",
		conf.User,
		conf.Password,
		conf.Host,
		conf.DBName,
		conf.Charset,
	)
}
