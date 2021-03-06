package model

const (
	PassLevelPoint = int32(10800)
	PassPointBoxID = int32(1000)
)

type User struct {
	UID      int64 `gorm:"primary_key" json:"uid"`
	PassType int32 `gorm:"not null;INDEX" json:"pass_type"`
	Point    int32 `gorm:"not null;INDEX" json:"point"`
}

// TableName return table name
func (*User) TableName() string {
	return "np_pass_users"
}

func (u *User) Level() int32 {
	if u.Point == 0 {
		return 0
	}
	return u.Point/PassLevelPoint + 1
}
