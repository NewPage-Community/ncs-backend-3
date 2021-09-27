package conf

import (
	"backend/pkg/conf"
	"backend/pkg/jwt"
	"backend/pkg/log"
	"github.com/go-redis/redis/v8"
)

type Config struct {
	Log   *log.Config
	Redis *redis.Options
	JWT   *jwt.JWT
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&redis.Options{},
		&jwt.JWT{},
	}
	conf.Load(c)
	return
}
