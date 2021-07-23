package jwt

import "encoding/json"

type Payload map[string]interface{}

func (i *Payload) Get(key string) interface{} {
	return (*i)[key]
}

func (i *Payload) Set(key string, value interface{}) {
	(*i)[key] = value
}

func (i *Payload) GetInt64(key string) int64 {
	number, _ := (*i)[key].(json.Number)
	d, _ := number.Int64()
	return d
}

func (i *Payload) GetFloat64(key string) float64 {
	number, _ := (*i)[key].(json.Number)
	d, _ := number.Float64()
	return d
}

func (i *Payload) GetString(key string) string {
	s, _ := (*i)[key].(string)
	return s
}
