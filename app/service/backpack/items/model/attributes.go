package model

type Attributes map[string]interface{}

func (attr Attributes) Set(key string, value interface{}) {
	attr[key] = value
}

func (attr Attributes) Get(key string) (interface{}, bool) {
	value, ok := attr[key]
	return value, ok
}
