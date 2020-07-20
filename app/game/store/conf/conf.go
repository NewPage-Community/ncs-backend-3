package conf

import (
	"backend/app/game/store/model"
	"backend/pkg/conf"
	"backend/pkg/log"
)

type Config struct {
	Log     *log.Config
	HotSale *model.HotSale
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&model.HotSale{},
	}
	conf.Load(c)
	return
}
