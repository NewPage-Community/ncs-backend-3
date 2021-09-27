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
	Kaiheila *Kaiheila
	Redis    *goredis.Options
}

type Kaiheila struct {
	Token            string
	AllChatChannelID string
}

// Init 初始化配置
func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&Kaiheila{},
		&goredis.Options{},
	}
	// 读取配置
	conf.Load(c)
	return
}
