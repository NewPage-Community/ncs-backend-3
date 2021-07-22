package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/jwt"
	"backend/pkg/log"
	"backend/pkg/rpc"
	"golang.org/x/oauth2"
)

// Config 定义配置结构
type Config struct {
	Log       *log.Config
	QQConnect *oauth2.Config
	JWT       *jwt.JWT
	Mysql     *mysql.Config
	Rpc       *rpc.ServerConfig
}

// Init 初始化配置
func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&oauth2.Config{},
		&jwt.JWT{},
		&mysql.Config{},
		&rpc.ServerConfig{},
	}
	// 读取配置
	conf.Load(c)
	return
}
