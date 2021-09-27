package conf

import (
	"backend/pkg/conf"
	"backend/pkg/log"

	goredis "github.com/go-redis/redis/v8"
)

// Config 定义配置结构
// 已包含Mysql配置
type Config struct {
	Log      *log.Config
	QQConfig *QQConfig
	Redis    *goredis.Options
}

type QQConfig struct {
	Address  string
	APIPort  int
	WSPort   int
	Token    string
	QQGroups []int
}

// Init 初始化配置
func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&QQConfig{},
		&goredis.Options{},
	}
	// 读取配置
	conf.Load(c)
	return
}
