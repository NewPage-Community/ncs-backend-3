package model

import (
	"backend/pkg/json"
	"time"
)

type Record struct {
	Timestamp int64  `json:"timestamp"`
	Amount    int32  `json:"amount"`
	Reason    string `json:"reason"`
}

type Records []*Record

func GetRecordsFromJSON(data []byte) (r *Records, err error) {
	r = &Records{}
	err = json.Unmarshal(data, r)
	return
}

func (r *Records) Add(amount int32, reason string) {
	*r = append(*r, &Record{
		Timestamp: time.Now().Unix(),
		Amount:    amount,
		Reason:    reason,
	})
}

func (r *Records) RemoveExpr() {
	lastMonth := time.Now().AddDate(0, -1, 0).Unix()
	if len(*r) > 0 {
		for (*r)[0].Timestamp < lastMonth {
			*r = (*r)[1:]
			if len(*r) == 0 {
				break
			}
		}
	}
}

func (r *Records) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}
