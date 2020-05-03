package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
)

type Config struct {
	Log   *log.Config
	Mysql *mysql.Config
}

func Init() (c *Config) {
	c = &Config{}
	conf.Load(c)
	return
}
