package conf

import (
	"backend/pkg/conf"
	"backend/pkg/log"
)

// Config 定义配置结构
// 已包含Mysql配置
type Config struct {
	Log *log.Config
}

// Init 初始化配置
func Init() (c *Config) {
	c = &Config{
		&log.Config{},
	}
	// 读取配置
	conf.Load(c)
	return
}
