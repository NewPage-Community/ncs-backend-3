package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
	"github.com/go-redis/redis/v7"
)

type Config struct {
	Log   *log.Config
	Mysql *mysql.Config
	Redis *redis.Options
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&mysql.Config{},
		&redis.Options{},
	}
	conf.Load(c)
	return
}
