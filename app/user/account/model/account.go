package model

type Info struct {
	UID       int64  `gorm:"primary_key;unique;not null" json:"uid"`
	SteamID   int64  `gorm:"unique;not null;INDEX" json:"steam_id"`
	Username  string `json:"username"`
	FirstJoin int64  `gorm:"not null;INDEX" json:"first_join"`
}

// TableName return table name
func (*Info) TableName() string {
	return "np_accounts"
}

// IsValid .
func (i *Info) IsValid() bool {
	return i.UID > 0
}
