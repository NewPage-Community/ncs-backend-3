package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
)

// TODO 默认值&单元测试

type Config struct {
	Log   *log.Config
	Mysql *mysql.Config
}

func Init() (c *Config) {
	c = &Config{}
	conf.Load(c)
	return
}
