package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
)

// Unmarshal ...
func Unmarshal(data []byte, v interface{}) (err error) {
	err = jsoniter.Unmarshal(data, v)
	if err != nil {
		err = fmt.Errorf("json unmarshal error: %s (%s)", err.Error(), string(data))
	}
	return
}

// Marshal ...
func Marshal(v interface{}) (s []byte, err error) {
	s, err = jsoniter.Marshal(v)
	if err != nil {
		err = fmt.Errorf("json marshal error: %s (%v)", err.Error(), v)
	}
	return
}
