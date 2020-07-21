package model

import (
	"backend/pkg/json"
	"sort"
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

func (item *Item) IsExpired() bool {
	return item.ExprTime != 0 && item.ExprTime < time.Now().Unix()
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

func LoadItemsFromJSON(data []byte) (items *Items, err error) {
	items = &Items{}
	err = json.Unmarshal(data, items)
	return
}

func (items *Items) AddItems(_items *Items) {
	for i := range *_items {
		items.addItem(*(*_items)[i])
	}
	sort.Sort(items)
}

func (items *Items) addItem(item Item) {
	index, found := items.search(item.ID)
	now := time.Now().Unix()

	if found {
		if item.CalType() == ItemCalTypeUnlimited {
			(*items)[index].Amount = 0
			(*items)[index].ExprTime = 0
		} else if (*items)[index].CalType() == ItemCalTypeAmount {
			(*items)[index].Amount += item.Amount
		} else if (*items)[index].CalType() == ItemCalTypeTime {
			(*items)[index].ExprTime += item.ExprTime
		}
	} else {
		// not found
		// fix expr time
		if item.CalType() == ItemCalTypeTime {
			item.ExprTime = now + item.ExprTime
		}
		*items = append(*items, &item)
	}
}

func (items *Items) RemoveItem(item Item, all bool) {
	index, found := items.search(item.ID)
	length := len(*items)

	if found {
		if all || (*items)[index].Amount <= 1 {
			// Last one
			if index+1 == length {
				*items = (*items)[:index]
			} else {
				*items = append((*items)[:index], (*items)[index+1:]...)
			}
		} else {
			(*items)[index].Amount--
		}
	}
}

func (items *Items) SearchItem(id int32) (item Item, found bool) {
	index, found := items.search(id)

	if found {
		item = *(*items)[index]
	}
	return
}

//search items Binary Search
func (items *Items) search(id int32) (index int, found bool) {
	length := len(*items)
	low := 0
	high := length - 1
	for low <= high {
		mid := (low + high) / 2
		if (*items)[mid].ID > id {
			high = mid - 1
		} else if (*items)[mid].ID < id {
			low = mid + 1
		} else {
			return mid, true
		}
	}
	return length, false
}

func (items *Items) Len() int {
	return len(*items)
}

func (items *Items) Less(i, j int) bool {
	return (*items)[i].ID < (*items)[j].ID
}

func (items *Items) Swap(i, j int) {
	(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
}

func (items *Items) JSON() ([]byte, error) {
	return json.Marshal(items)
}

func (items *Items) Check() {
	for index := 0; index < len(*items); index++ {
		if (*items)[index].IsExpired() {
			// Last one
			if index+1 == len(*items) {
				*items = (*items)[:index]
			} else {
				*items = append((*items)[:index], (*items)[index+1:]...)
			}
		}
	}
}
