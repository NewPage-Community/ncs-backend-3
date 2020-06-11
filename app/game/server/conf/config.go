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
	c = &Config{
		&log.Config{},
		&mysql.Config{},
	}
	conf.Load(c)
	return
}
