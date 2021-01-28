package conf

import (
	"backend/pkg/conf"
	"backend/pkg/log"
)

// Config 定义配置结构
// 已包含Mysql配置
type Config struct {
	Log      *log.Config
	Kaiheila *Kaiheila
}

type Kaiheila struct {
	Token string
}

// Init 初始化配置
func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&Kaiheila{},
	}
	// 读取配置
	conf.Load(c)
	return
}
