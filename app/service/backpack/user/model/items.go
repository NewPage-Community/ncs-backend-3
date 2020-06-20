package model

import (
	"backend/pkg/json"
	"sort"
	"time"
)

type Item struct {
	ID     int32 `json:"id"`
	Amount int32 `json:"amount"`
	// 0 = unlimited time
	ExprTime int64 `json:"expr_time"`
}

func (item *Item) IsExpired() bool {
	return item.ExprTime != 0 && item.ExprTime < time.Now().Unix()
}

type Items []Item

func LoadItemsFromJSON(data []byte) (items *Items, err error) {
	items = &Items{}
	err = json.Unmarshal(data, items)
	return
}

func (items *Items) AddItem(item Item, repeat bool) {
	index, found := items.search(item.ID)
	if item.Amount <= 0 {
		item.Amount = 1
	}

	if found {
		// found the same item and +1
		// check it allow repeat
		if repeat {
			(*items)[index].Amount += item.Amount
		}
	} else {
		// not found
		*items = append(*items, item)
		sort.Sort(items)
	}
}

func (items *Items) RemoveItem(item Item, all bool) {
	index, found := items.search(item.ID)

	if found {
		if all || (*items)[index].Amount <= 1 {
			*items = append((*items)[:index], (*items)[index+1:]...)
		} else {
			(*items)[index].Amount--
		}
	}
}

func (items *Items) SearchItem(id int32) (item Item, found bool) {
	index, found := items.search(id)

	if found {
		item = (*items)[index]
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
	for index := range *items {
		if (*items)[index].IsExpired() || (*items)[index].Amount <= 0 {
			*items = append((*items)[:index], (*items)[index+1:]...)
		}
	}
}
