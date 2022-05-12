package steam

import (
	"net/url"
	"strconv"
)

const (
	ISteamUserSrv = "ISteamUser"
)

type UserGroupList struct {
	Group []struct {
		GID string `json:"gid"`
	} `json:"group"`
}

func (a *API) GetUserGroupList(steamID uint64) (res UserGroupList, err error) {
	v := url.Values{}
	v.Add("steamid", strconv.FormatInt(int64(steamID), 10))

	data := &struct {
		Response UserGroupList
	}{}
	err = a.request(APIRequest{
		Service: ISteamUserSrv,
		Method:  "GetUserGroupList",
		Version: 1,
	}, v, &data)
	res = data.Response
	return
}
