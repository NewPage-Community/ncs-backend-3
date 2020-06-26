package model

type SkinUser struct {
	UID      int64 `gorm:"primary_key" json:"uid"`
	UsedSkin int32 `json:"used_skin"`
}

// TableName return table name
func (*SkinUser) TableName() string {
	return "np_skin_users"
}

func (i *SkinUser) IsValid() bool {
	return i.UID > 0
}
