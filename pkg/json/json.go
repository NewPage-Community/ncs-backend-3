package json

import (
	"encoding/json"
	"fmt"
	"io"
)

// Unmarshal ...
func Unmarshal(data []byte, v interface{}) (err error) {
	err = json.Unmarshal(data, v)
	if err != nil {
		err = fmt.Errorf("json unmarshal error: %s (%s)", err.Error(), string(data))
	}
	return
}

// Marshal ...
func Marshal(v interface{}) (s []byte, err error) {
	s, err = json.Marshal(v)
	if err != nil {
		err = fmt.Errorf("json marshal error: %s (%v)", err.Error(), v)
	}
	return
}

func NewDecoder(reader io.Reader) *json.Decoder {
	return json.NewDecoder(reader)
}
