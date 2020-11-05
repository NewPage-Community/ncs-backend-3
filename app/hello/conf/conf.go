package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
)

// Config 定义配置结构
// 已包含Mysql配置
type Config struct {
	Log    *log.Config
	Mysql  *mysql.Config
}

// Init 初始化配置
func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&mysql.Config{},
	}
	// 读取配置
	conf.Load(c)
	return
}
