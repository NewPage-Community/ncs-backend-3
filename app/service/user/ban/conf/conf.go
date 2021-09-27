package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
	"backend/pkg/steam"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Log   *log.Config
	Mysql *mysql.Config
	Redis *redis.Options
	Steam *steam.API
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&mysql.Config{},
		&redis.Options{},
		&steam.API{},
	}
	conf.Load(c)
	return
}
