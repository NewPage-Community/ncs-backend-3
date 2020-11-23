package model

type QQConnect struct {
	UID    int64 `gorm:"primary_key"`
	OpenID string
}

// TableName 返回数据库表名
func (*QQConnect) TableName() string {
	return "np_qq_connect"
}
