package model

import (
	"gorm.io/datatypes"
)

type UserModel struct {
	UID int64 `gorm:"primary_key" json:"uid"`
	// []Item{}
	Items datatypes.JSON `json:"items"`
}

// TableName return table name
func (*UserModel) TableName() string {
	return "np_backpack_users"
}

// IsValid .
func (i *UserModel) IsValid() bool {
	return i.UID > 0
}

func (i *UserModel) GetUser() (*User, error) {
	items, err := LoadItemsFromJSON(i.Items)
	return &User{
		UID:   i.UID,
		Items: items,
	}, err
}

type User struct {
	UID   int64  `json:"uid"`
	Items *Items `json:"items"`
}

// IsValid .
func (i *User) IsValid() bool {
	return i.UID > 0
}

func (i *User) AddItems(items *Items) {
	i.Items.AddItems(items)
}

func (i *User) RemoveItem(item Item, all bool) {
	i.Items.RemoveItem(item, all)
}

func (i *User) SearchItem(id int32) (item Item, found bool) {
	return i.Items.SearchItem(id)
}

func (i *User) CheckItem() {
	i.Items.Check()
}

func (i *User) GetModel() (*UserModel, error) {
	json, err := i.Items.JSON()
	return &UserModel{
		UID:   i.UID,
		Items: json,
	}, err
}
