package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
)

type Config struct {
	Log    *log.Config
	Alipay *Alipay
	Mysql  *mysql.Config
}

type Alipay struct {
	AppID        string
	PrivateKey   string
	IsProduction bool
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&Alipay{},
		&mysql.Config{},
	}
	conf.Load(c)
	return
}
