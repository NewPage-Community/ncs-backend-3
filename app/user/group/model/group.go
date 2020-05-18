package model

import (
	"backend/pkg/json"
	"sort"
)

type Group struct {
	GID       int64  `gorm:"primary_key;unique;not null" json:"GID"`
	Name      string `gorm:"not null" json:"Name"`
	ShortName string `gorm:"not null;INDEX" json:"ShortName"`
}

// TableName return table name
func (*Group) TableName() string {
	return "np_groups"
}

// IsValid .
func (g *Group) IsValid() bool {
	return g.GID > 0
}

type User struct {
	UID    int64  `gorm:"primary_key;unique;not null" json:"UID"`
	Groups string `gorm:"not null" json:"Groups"`
}

// TableName return table name
func (*User) TableName() string {
	return "np_groups_user"
}

// IsValid .
func (u *User) IsValid() bool {
	return u.UID > 0
}

type Groups struct {
	gID []int64
}

func NewGroups(groups string) (g *Groups, err error) {
	g = &Groups{gID: make([]int64, 0)}
	err = json.Unmarshal([]byte(groups), &g.gID)
	g.Sort()
	return
}

func (g *Groups) InGroup(gid int64) bool {
	for _, v := range g.gID {
		if v == gid {
			return true
		}
	}
	return false
}

func (g *Groups) ToJSON() (s string, err error) {
	t, err := json.Marshal(g.gID)
	s = string(t)
	return
}

func (g *Groups) Count() int {
	return len(g.gID)
}

func (g *Groups) Add(gid int64) {
	if !g.InGroup(gid) && gid != 0 {
		g.gID = append(g.gID, gid)
		g.Sort()
	}
}

func (g *Groups) Sort() {
	if len(g.gID) > 0 {
		sort.Slice(g.gID, func(i, j int) bool {
			return g.gID[i] < g.gID[j]
		})
	}
}
