package model

import (
	"backend/pkg/json"
	"strconv"
)

const (
	KeyPrefix = "ncs:game:"
)

type Cookie struct {
	UID    int64             `json:"uid"`
	Cookie map[string]string `json:"cookie"`
}

func (Cookie) Key() string {
	return KeyPrefix + "cookie"
}

func (c *Cookie) GetUID() string {
	return strconv.FormatInt(c.UID, 10)
}

func (c *Cookie) IsValid() bool {
	return c.UID > 0
}

func (c *Cookie) CookieJSON() (string, error) {
	j, err := json.Marshal(c.Cookie)
	return string(j), err
}

func (c *Cookie) GetCookieFromJSON(data string) error {
	return json.Unmarshal([]byte(data), &c.Cookie)
}
