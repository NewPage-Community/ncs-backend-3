package conf

import (
	"backend/pkg/conf"
	"backend/pkg/log"
)

type Config struct {
	Log *log.Config
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
	}
	conf.Load(c)
	return
}
