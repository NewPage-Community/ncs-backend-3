package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"

	goredis "github.com/go-redis/redis/v8"
)

type Config struct {
	Log   *log.Config
	Mysql *mysql.Config
	Redis *goredis.Options
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&mysql.Config{},
		&goredis.Options{},
	}
	conf.Load(c)
	return
}
