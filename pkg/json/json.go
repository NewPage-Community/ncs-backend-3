package json

import (
	"backend/pkg/ecode"
	jsoniter "github.com/json-iterator/go"
	"google.golang.org/grpc/codes"
)

// Unmarshal ...
func Unmarshal(data []byte, v interface{}) (err error) {
	err = jsoniter.Unmarshal(data, v)
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "JSON Unmarshal error: %v", err)
	}
	return
}

// Marshal ...
func Marshal(v interface{}) (s []byte, err error) {
	s, err = jsoniter.Marshal(v)
	if err != nil {
		err = ecode.Errorf(codes.Unknown, "JSON Marshal error: %v", err)
	}
	return
}
