package model

import (
	"backend/pkg/json"
	"gorm.io/datatypes"
)

type Item struct {
	ID          int32          `json:"id"`
	Name        string         `gorm:"not null" json:"name"`
	Description string         `gorm:"not null" json:"description"`
	Type        int32          `gorm:"not null" json:"type"`
	Price       int32          `gorm:"not null" json:"price"`
	Attributes  datatypes.JSON `json:"attributes"`
}

func (Item) TableName() string {
	return "np_backpack_items"
}

func (item *Item) SetAttributes(v *Attributes) (err error) {
	item.Attributes, err = json.Marshal(v)
	return
}

func (item *Item) GetAttributes(v *Attributes) (err error) {
	err = json.Unmarshal(item.Attributes, v)
	return
}
