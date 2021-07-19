package conf

import (
	"backend/pkg/conf"
	"backend/pkg/jwt"
	"backend/pkg/log"
)

type OpenID struct {
	Endpoint  string
	Host      string
	HttpProxy string
}

// Config 定义配置结构
type Config struct {
	Log         *log.Config
	JWT         *jwt.JWT
	SteamOpenID *OpenID
}

// Init 初始化配置
func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&jwt.JWT{},
		&OpenID{},
	}
	// 读取配置
	conf.Load(c)
	return
}
