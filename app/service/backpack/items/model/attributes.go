package model

type Attributes map[string]string

func (attr Attributes) Set(key string, value string) {
	attr[key] = value
}

func (attr Attributes) Get(key string) (string, bool) {
	value, ok := attr[key]
	return value, ok
}
