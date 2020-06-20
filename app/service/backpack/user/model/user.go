package model

import (
	"gorm.io/datatypes"
)

type User struct {
	UID int64 `gorm:"primary_key" json:"uid"`
	// []Item{}
	Items datatypes.JSON `json:"items"`
}

// TableName return table name
func (*User) TableName() string {
	return "np_backpack_users"
}

// IsValid .
func (i *User) IsValid() bool {
	return i.UID > 0
}

func (i *User) AddItem(item Item, repeat bool) (err error) {
	// Unmarshal -> Add -> Marshal
	items, err := LoadItemsFromJSON(i.Items)
	if err == nil {
		items.AddItem(item, repeat)
		i.Items, err = items.JSON()
	}
	return
}

func (i *User) RemoveItem(item Item, all bool) (err error) {
	// Unmarshal -> Remove -> Marshal
	items, err := LoadItemsFromJSON(i.Items)
	if err == nil {
		items.RemoveItem(item, all)
		i.Items, err = items.JSON()
	}
	return
}

func (i *User) SearchItem(id int32) (item Item, found bool) {
	// Unmarshal -> Search
	items, err := LoadItemsFromJSON(i.Items)
	if err == nil {
		return items.SearchItem(id)
	}
	return Item{}, false
}
