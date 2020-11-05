package model

// Hello 数据库模型
type Hello struct {
	Hello string
}

// TableName 返回数据库表名
func (*Hello) TableName() string {
	return "np_hello"
}
