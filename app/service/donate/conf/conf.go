package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"

	goredis "github.com/go-redis/redis/v8"
)

type Config struct {
	Log    *log.Config
	Alipay *Alipay
	Wepay  *Wepay
	Mysql  *mysql.Config
	Redis  *goredis.Options
	Donate *Donate
}

type Alipay struct {
	AppID           string
	AlipayPublicKey string
	PrivateKey      string
	IsProduction    bool
}

type Wepay struct {
	AppID                      string
	MchID                      string
	MchCertificateSerialNumber string
	MchPrivateKey              string
	MchAPIv3Key                string
}

type Donate struct {
	Target int
}

func Init() (c *Config) {
	c = &Config{
		&log.Config{},
		&Alipay{},
		&Wepay{},
		&mysql.Config{},
		&goredis.Options{},
		&Donate{},
	}
	conf.Load(c)
	return
}
