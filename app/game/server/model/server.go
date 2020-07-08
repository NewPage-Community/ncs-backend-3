package model

type Info struct {
	ServerID int32  `gorm:"primary_key;unique;not null" json:"server_id"`
	ModID    int32  `gorm:"not null" json:"mod_id"`
	GameID   int32  `gorm:"not null" json:"game_id"`
	Rcon     string `gorm:"not null" json:"rcon"`
	Hostname string `gorm:"not null" json:"hostname"`
	Address  string `gorm:"not null;unique" json:"address"`
}

// TableName return table name
func (*Info) TableName() string {
	return "np_servers"
}

// IsValid .
func (i *Info) IsValid() bool {
	return i.ServerID > 0
}

func (i *Info) GenerateRcon() {
	i.Rcon = getRandomString(24)
}
