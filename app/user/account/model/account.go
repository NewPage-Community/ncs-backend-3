package model

type Info struct {
	UID       int64  `gorm:"primary_key" json:"UID"`
	SteamID   int64  `gorm:"column:steamid;;unique;not null;UNIQUE_INDEX" json:"SteamID"`
	Username  string `gorm:"column:username" json:"Username"`
	FirstJoin int64  `gorm:"column:firstjoin;not null" json:"FirstJoin"`
}

// TableName return table name
func (*Info) TableName() string {
	return "np_accounts"
}

// IsValid .
func (i *Info) IsValid() bool {
	return i.UID > 0
}
