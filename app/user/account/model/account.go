package model

type Info struct {
	UID       int64  `gorm:"primary_key;unique;not null" json:"UID"`
	SteamID   int64  `gorm:"unique;not null;INDEX" json:"SteamID"`
	Username  string `json:"Username"`
	FirstJoin int64  `gorm:"not null;INDEX" json:"FirstJoin"`
}

// TableName return table name
func (*Info) TableName() string {
	return "np_accounts"
}

// IsValid .
func (i *Info) IsValid() bool {
	return i.UID > 0
}
