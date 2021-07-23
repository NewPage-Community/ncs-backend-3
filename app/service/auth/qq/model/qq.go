package model

import (
	"backend/pkg/jwt"
)

type QQConnect struct {
	UID    int64 `gorm:"primary_key"`
	OpenID string
}

// TableName 返回数据库表名
func (*QQConnect) TableName() string {
	return "np_qq_connect"
}

func (i *QQConnect) GetJWTPayload(parent *jwt.Payload) *jwt.Payload {
	parent.Set("uid", i.UID)
	parent.Set("qq_open_id", i.OpenID)
	return parent
}

func GetQQConnectFromJWTPayload(payload *jwt.Payload) *QQConnect {
	return &QQConnect{
		UID:    payload.GetInt64("uid"),
		OpenID: payload.GetString("qq_open_id"),
	}
}
