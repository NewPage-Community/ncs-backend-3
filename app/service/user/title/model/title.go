package model

type Title struct {
	UID         int64  `gorm:"primary_key" json:"uid"`
	CustomTitle string `json:"custom_title"`
	TitleType   int32  `gorm:"not null" json:"title_type"`
}

// TableName return table name
func (*Title) TableName() string {
	return "np_title"
}

func (t *Title) IsValid() bool {
	return t.UID > 0
}
