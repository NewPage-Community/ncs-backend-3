package model

import (
	"backend/pkg/json"
	"fmt"
	"strconv"
)

const (
	KeyPrefix = "ncs:game:"
)

type CookieModel struct {
	Value       string `json:"value"`
	LastUpdated int64  `json:"last_updated"`
}

type Cookie struct {
	UID    int64                  `json:"uid"`
	Cookie map[string]CookieModel `json:"cookie"`
}

type CookieOld struct {
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

func (c *Cookie) MergeCookieData(data string) error {
	var dataFormat struct {
		UID    int64                  `json:"uid"`
		Cookie map[string]interface{} `json:"cookie"`
	}
	err := json.Unmarshal([]byte(data), &dataFormat)
	if err != nil {
		return err
	}

	var newCookie = make(map[string]CookieModel)
	for k, v := range dataFormat.Cookie {
		// check type
		if _, ok := v.(string); ok {
			// old format need to convert
			newCookie[k] = CookieModel{
				Value:       fmt.Sprintf("%v", v),
				LastUpdated: 0,
			}
		} else {
			// ingore invalid data
			value, ok := v.(map[string]interface{})["value"].(string)
			if !ok {
				continue
			}
			lastUpdated, ok := v.(map[string]interface{})["last_updated"].(float64)
			if !ok {
				continue
			}
			newCookie[k] = CookieModel{
				Value:       value,
				LastUpdated: int64(lastUpdated),
			}
		}
	}
	c.UID = dataFormat.UID
	c.Cookie = newCookie
	return nil
}
