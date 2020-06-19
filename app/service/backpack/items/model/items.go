package model

import (
	"backend/pkg/json"
	"gorm.io/datatypes"
	"sort"
	"time"
)

type Item struct {
	// Model
	ID         int32          `json:"id"`
	Type       int32          `gorm:"not null" json:"-"`
	Price      int32          `gorm:"not null" json:"-"`
	Attributes datatypes.JSON `json:"-"`

	// User backpack
	Amount int32 `gorm:"-" json:"amount"`
	// 0 = unlimited time
	ExprTime int64 `gorm:"-" json:"expr_time"`
}

func (item *Item) IsExpired() bool {
	return item.ExprTime != 0 && item.ExprTime < time.Now().Unix()
}

func (item *Item) SetAttributes(v interface{}) (err error) {
	item.Attributes, err = json.Marshal(v)
	return
}

func (item *Item) GetAttributes(v interface{}) (err error) {
	err = json.Unmarshal(item.Attributes, v)
	return
}

type Items struct {
	items []Item
}

func LoadItemsFromJSON(data []byte) (items *Items, err error) {
	items = &Items{}
	err = json.Unmarshal(data, items.items)
	return
}

func (items *Items) AddItem(item Item, repeat bool) {
	index, found := items.search(item.ID)

	if found {
		// found the same item and +1
		// check it allow repeat
		if repeat {
			items.items[index].Amount++
		}
	} else {
		// not found
		items.items = append(items.items, item)
		sort.Sort(items)
	}
}

func (items *Items) RemoveItem(item Item, all bool) {
	index, found := items.search(item.ID)

	if found {
		if all || items.items[index].Amount <= 1 {
			items.items = append(items.items[:index], items.items[index+1:]...)
		} else {
			items.items[index].Amount--
		}
	}
}

func (items *Items) SearchItem(id int32) (item Item, found bool) {
	index, found := items.search(id)

	if found {
		item = items.items[index]
	}
	return
}

func (items *Items) search(id int32) (index int, found bool) {
	itemsNum := len(items.items)
	index = sort.Search(itemsNum, func(j int) bool {
		return items.items[j].ID == id
	})
	found = index != itemsNum
	return
}

func (items *Items) Len() int {
	return len(items.items)
}

func (items *Items) Less(i, j int) bool {
	return items.items[i].ID < items.items[j].ID
}

func (items *Items) Swap(i, j int) {
	items.items[i], items.items[j] = items.items[j], items.items[i]
}

func (items *Items) JSON() ([]byte, error) {
	return json.Marshal(items.items)
}
