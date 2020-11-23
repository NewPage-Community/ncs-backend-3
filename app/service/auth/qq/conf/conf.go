package conf

import (
	"backend/app/service/auth/jwt"
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
	"golang.org/x/oauth2"
)

// Config 定义配置结构
type Config struct {
	Log       *log.Config
	QQConnect *oauth2.Config
	JWT       *jwt.Config
	Mysql     *mysql.Config
}

// Init 初始化配置
func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&oauth2.Config{},
		&jwt.Config{},
		&mysql.Config{},
	}
	// 读取配置
	conf.Load(c)
	return
}
