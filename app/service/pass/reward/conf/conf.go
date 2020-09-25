package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
)

type Config struct {
	Log        *log.Config
	Mysql      *mysql.Config
	RewardConf *RewardConf
}

type RewardConf struct {
	Season   int32
	MaxLevel int32
	Free     string
	Adv      string
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&mysql.Config{},
		&RewardConf{},
	}
	conf.Load(c)
	return
}
