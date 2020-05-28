package model

type Info struct {
	ServerID int64  `gorm:"primary_key;unique;not null" json:"server_id"`
	ModID    int64  `gorm:"not null" json:"mod_id"`
	GameID   int64  `gorm:"not null" json:"game_id"`
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

//TODO: random
func (i *Info) GenerateRcon() {
	i.Rcon = ""
}
