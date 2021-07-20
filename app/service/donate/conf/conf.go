package conf

import (
	"backend/pkg/conf"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
)

type Config struct {
	Log    *log.Config
	Alipay *Alipay
	Wepay  *Wepay
	Mysql  *mysql.Config
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
		&Donate{},
	}
	conf.Load(c)
	return
}
