package model

import "time"

const (
	BanTypeAll = iota
	BanTypeMod
	BanTypeGame
	BanTypeServer
)

type Ban struct {
	ID         int64  `gorm:"primary_key;unique;not null" json:"id"`
	SteamID    int64  `gorm:"not null" json:"steam_id"`
	CreateTime int64  `gorm:"not null;INDEX" json:"create_time"'`
	Length     int64  `gorm:"not null;INDEX" json:"length"`
	Type       int8   `gorm:"not null" json:"type"`
	ServerID   int64  `gorm:"not null" json:"server_id"`
	ModID      int64  `gorm:"not null" json:"mod_id"`
	GameID     int64  `gorm:"not null" json:"game_id"`
	OpUID      int64  `gorm:"not null" json:"op_uid"`
	RemoveUID  int64  `gorm:"not null" json:"remove_uid"`
	Reason     string `gorm:"not null" json:"reason"`
}

// TableName return table name
func (*Ban) TableName() string {
	return "np_bans"
}

// IsValid .
func (i *Ban) IsValid() bool {
	return i.ID > 0
}

// IsExpired check ban record is expired?
func (i *Ban) IsExpired() bool {
	return i.CreateTime+i.Length < time.Now().Unix() && i.Length != 0
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
