package model

type Title struct {
	UID         int64  `gorm:"unique" json:"uid"`
	CustomTitle string `json:"custom_title"`
	TitleType   int32  `json:"title_type"`
}

// TableName return table name
func (*Title) TableName() string {
	return "np_title"
}
