package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/jwt"
	"backend/pkg/log"
)

type Config struct {
	Log   *log.Config
	Mysql *mysql.Config
	JWT   *jwt.JWT
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&mysql.Config{},
		&jwt.JWT{},
	}
	conf.Load(c)
	return
}
