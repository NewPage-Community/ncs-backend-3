package main

import (
	"encoding/json"
	"gorm.io/datatypes"
)

type User struct {
	UID        int32 `gorm:"PRIMARY_KEY"`
	Username   string
	Steamid    int64
	Money      int32
	Firstjoin  int32
	Signtimes  int32
	Vippoint   int32
	Vipexpired int64
}

func (User) TableName() string {
	return "np_users"
}

type SkinStore struct {
	UID  int32 `gorm:"PRIMARY_KEY"`
	Skin datatypes.JSON
}

func (SkinStore) TableName() string {
	return "np_skins_store"
}

type SkinInfo struct {
	UID  string `json:"uid"`
	Time int64  `json:"time"`
}

type Skins []*SkinInfo

func (s *SkinStore) GetSkins() (res *Skins, err error) {
	res = &Skins{}
	err = json.Unmarshal(s.Skin, res)
	return
}

type Skin struct {
	UID   string `gorm:"PRIMARY_KEY"`
	Name  string
	Model string
	Arm   string
	ID    int32
}

func (Skin) TableName() string {
	return "np_skins"
}
