package model

type Item struct {
	Level  int32
	ID     int32
	Amount int32
	Length int64
}

type Reward struct {
	Season      int32 `json:"season"`
	MaxLevel    int32 `json:"max_level"`
	FreeRewards []Item
	AdvRewards  []Item
}
