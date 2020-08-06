package conf

import (
	"backend/pkg/conf"
	"backend/pkg/log"
	"github.com/go-redis/redis/v7"
)

type Config struct {
	Log   *log.Config
	Redis *redis.Options
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&redis.Options{},
	}
	conf.Load(c)
	return
}
