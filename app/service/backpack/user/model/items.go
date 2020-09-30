package model

import (
	"backend/pkg/json"
	"time"
)

const (
	ItemCalTypeUnlimited = iota
	ItemCalTypeAmount
	ItemCalTypeTime
)

type Item struct {
	ID     int32 `json:"id"`
	Amount int32 `json:"amount"`
	// 0 = unlimited time
	ExprTime int64 `json:"expr_time"`
}

func (item *Item) IsValid() bool {
	return item.ID > 0
}

func (item *Item) IsExpired(now int64) bool {
	// Unlimited time
	if item.ExprTime == 0 {
		return false
	}
	return item.ExprTime < now
}

func (item *Item) CalType() int {
	if item.Amount == 0 && item.ExprTime == 0 {
		return ItemCalTypeUnlimited
	}
	if item.ExprTime == 0 {
		return ItemCalTypeAmount
	}
	return ItemCalTypeTime
}

type Items []*Item
type ItemsMap map[int32]*Item

func LoadItemsFromJSON(data []byte) (items *ItemsMap, err error) {
	items = &ItemsMap{}
	err = json.Unmarshal(data, items)
	return
}

func (items *ItemsMap) AddItems(_items *Items) {
	for i := range *_items {
		items.addItem(*(*_items)[i])
	}
}

func (items *ItemsMap) addItem(item Item) {
	id := item.ID
	_, found := (*items)[id]

	if found {
		if item.CalType() == ItemCalTypeUnlimited {
			(*items)[id].Amount = 0
			(*items)[id].ExprTime = 0
		} else if (*items)[id].CalType() == ItemCalTypeAmount {
			(*items)[id].Amount += item.Amount
		} else if (*items)[id].CalType() == ItemCalTypeTime {
			(*items)[id].ExprTime += item.ExprTime
		}
	} else {
		// not found
		// fix expr time
		if item.CalType() == ItemCalTypeTime {
			item.ExprTime = time.Now().Unix() + item.ExprTime
		}
		(*items)[id] = &item
	}
}

func (items *ItemsMap) RemoveItem(item Item, all bool) {
	id := item.ID
	_, found := (*items)[id]
	if found {
		// delete item when amount = 1 and need to remove one item
		// Do not happens amount = 0 , because item type is wrong
		// amount = 0 && time = 0 -> ItemCalTypeUnlimited
		if all || (*items)[id].Amount == 1 {
			delete(*items, id)
		} else {
			(*items)[id].Amount--
		}
	}
}

func (items *ItemsMap) SearchItem(id int32) (item Item, found bool) {
	tmp, found := (*items)[id]
	if found {
		item = *tmp
	}
	return
}

func (items *ItemsMap) JSON() ([]byte, error) {
	return json.Marshal(items)
}

func (items *ItemsMap) Check() {
	now := time.Now().Unix()
	for k := range *items {
		if (*items)[k].IsExpired(now) {
			delete(*items, k)
		}
	}
}
