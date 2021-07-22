package dao

import (
	"backend/app/service/auth/qq/conf"
	"backend/app/service/auth/qq/model"
	"backend/pkg/database/mysql"
	"backend/pkg/log"
	"gorm.io/gorm"
)

// Dao 结构定义
type Dao interface {
	// 自定义接口
	GetUID(openID string) (*model.QQConnect, error)
	GetStatus(uid int64) (*model.QQConnect, error)
	BindQQ(info model.QQConnect) error
	UnbindQQ(info model.QQConnect) error
	// 以下必须包含
	Healthy() bool
	Close()
}

type dao struct {
	db *gorm.DB
}

// Init Dao初始化
func Init(config *conf.Config) (d *dao) {
	// Mysql 例子
	d = &dao{
		db: mysql.Init(),
	}
	// 数据库自动融合
	if err := d.db.AutoMigrate(&model.QQConnect{}); err != nil {
		log.Error(err)
	}
	return
}

// Healthy Dao层健康检查
func (d *dao) Healthy() bool {
	return mysql.Healthy(d.db)
}

// Close Dao层连接关闭
func (d *dao) Close() {
}
