package model

import (
	"backend/pkg/json"
)

type Record struct {
	Timestamp int64  `json:"timestamp"`
	Amount    int32  `json:"amount"`
	Reason    string `json:"reason"`
}

func (r *Record) JSON() ([]byte, error) {
	return json.Marshal(r)
}

type Records []*Record
