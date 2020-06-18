package model

type User struct {
	UID   int64 `gorm:"primary_key" json:"uid"`
	Type  int32 `gorm:"not null;INDEX" json:"type"`
	Point int32 `gorm:"not null;INDEX" json:"point"`
}

// TableName return table name
func (*User) TableName() string {
	return "np_pass_users"
}
