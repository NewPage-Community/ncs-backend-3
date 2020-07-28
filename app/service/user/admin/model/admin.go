package model

var gFlagBits [26]int

func init() {
	// a - a
	gFlagBits[0] = 1 << 0
	gFlagBits['b'-'a'] = 1 << 1
	gFlagBits['c'-'a'] = 1 << 2
	gFlagBits['d'-'a'] = 1 << 3
	gFlagBits['e'-'a'] = 1 << 4
	gFlagBits['f'-'a'] = 1 << 5
	gFlagBits['g'-'a'] = 1 << 6
	gFlagBits['h'-'a'] = 1 << 7
	gFlagBits['i'-'a'] = 1 << 8
	gFlagBits['j'-'a'] = 1 << 9
	gFlagBits['k'-'a'] = 1 << 10
	gFlagBits['l'-'a'] = 1 << 11
	gFlagBits['m'-'a'] = 1 << 12
	gFlagBits['n'-'a'] = 1 << 13
	gFlagBits['z'-'a'] = 1 << 14
	gFlagBits['o'-'a'] = 1 << 15
	gFlagBits['p'-'a'] = 1 << 16
	gFlagBits['q'-'a'] = 1 << 17
	gFlagBits['r'-'a'] = 1 << 18
	gFlagBits['s'-'a'] = 1 << 19
	gFlagBits['t'-'a'] = 1 << 20
}

type Admin struct {
	UID      int64  `gorm:"primary_key;unique;not null" json:"uid"`
	Flag     string `json:"flag"`
	Immunity int32  `gorm:"not null" json:"immunity"`
}

// TableName return table name
func (*Admin) TableName() string {
	return "np_admins"
}

// GetFlagBits .
func (admin *Admin) GetFlagBits() int {
	flag := 0
	tmp := []byte(admin.Flag)
	for index := 0; index < len(tmp); index++ {
		flag |= gFlagBits[tmp[index]-'a']
	}
	return flag
}
