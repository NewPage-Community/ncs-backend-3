package model

import "time"

const (
	BanTypeNone = iota
	BanTypeAll
	BanTypeGame
	BanTypeMod
	BanTypeServer
)

type Ban struct {
	ID         uint32 `gorm:"primary_key;not null" json:"id"`
	UID        uint64 `gorm:"not null" json:"uid"`
	IP         string `json:"ip"`
	CreateTime int64  `gorm:"not null;INDEX" json:"create_time"`
	ExpireTime int64  `gorm:"not null;INDEX" json:"expire_time"`
	Type       int32  `gorm:"not null" json:"type"`
	ServerID   int64  `gorm:"not null" json:"server_id"`
	ModID      int64  `gorm:"not null" json:"mod_id"`
	GameID     int64  `gorm:"not null" json:"game_id"`
	Reason     string `gorm:"not null" json:"reason"`
}

// TableName return table name
func (*Ban) TableName() string {
	return "np_bans"
}

// IsValid .
func (i *Ban) IsValid() bool {
	return i.UID > 0
}

// IsExpired check ban record is expired?
func (i *Ban) IsExpired() bool {
	return i.ExpireTime < time.Now().UTC().Unix()
}

// IsBanned .
func (i *Ban) IsBanned(serverID int64, modID int64, gameID int64) bool {
	banned := !i.IsExpired()
	switch i.Type {
	case BanTypeGame:
		banned = banned && gameID == i.GameID
	case BanTypeMod:
		banned = banned && modID == i.ModID
	case BanTypeServer:
		banned = banned && serverID == i.ServerID
	}
	return banned
}
