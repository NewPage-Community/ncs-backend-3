package model

import (
	"github.com/NewPage-Community/go-steam"
	"net"
	"time"
)

type Info struct {
	ServerID  int32                     `gorm:"primary_key;unique;not null" json:"server_id"`
	ModID     int32                     `gorm:"not null" json:"mod_id"`
	GameID    int32                     `gorm:"not null" json:"game_id"`
	Rcon      string                    `gorm:"not null" json:"rcon"`
	Hostname  string                    `gorm:"not null" json:"hostname"`
	Address   string                    `gorm:"not null;unique" json:"address"`
	ShortName string                    `gorm:"not null" json:"short_name"`
	A2SInfo   steam.InfoResponse        `gorm:"-"`
	A2SPlayer steam.PlayersInfoResponse `gorm:"-"`
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

func (i *Info) RCON(cmd string) (resp string, err error) {
	server, err := steam.Connect(i.Address, &steam.ConnectOptions{
		Dial: (&net.Dialer{
			Timeout: 500 * time.Millisecond,
		}).Dial,
		RCONPassword: i.Rcon,
	})
	if err != nil {
		return
	}

	resp, err = server.Send(cmd)
	server.Close()
	return
}

func (i *Info) RequestA2S() (err error) {
	server, err := steam.Connect(i.Address, &steam.ConnectOptions{
		Dial: (&net.Dialer{
			Timeout: 100 * time.Millisecond,
		}).Dial,
	})
	if err != nil {
		return
	}

	a2sInfo, err := server.Info()
	if err == nil && a2sInfo != nil {
		i.A2SInfo = *a2sInfo
	}
	a2sPlayer, err := server.PlayersInfo()
	if err == nil && a2sPlayer != nil {
		i.A2SPlayer = *a2sPlayer
	}
	server.Close()
	return
}
