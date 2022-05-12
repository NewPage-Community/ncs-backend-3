package conf

import (
	"backend/pkg/conf"
	"backend/pkg/log"
	"backend/pkg/steam"
)

type Config struct {
	Log   *log.Config
	Steam *steam.API
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&steam.API{},
	}
	conf.Load(c)
	return
}
