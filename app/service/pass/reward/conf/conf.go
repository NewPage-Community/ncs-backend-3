package conf

import (
	reward "backend/app/service/pass/reward/model"
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
)

type Config struct {
	Log    *log.Config
	Mysql  *mysql.Config
	Reward *reward.Reward
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&mysql.Config{},
		&reward.Reward{},
	}
	conf.Load(c)
	return
}
