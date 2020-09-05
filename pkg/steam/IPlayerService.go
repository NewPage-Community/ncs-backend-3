package steam

import (
	"net/url"
	"strconv"
)

const (
	SrvName = "IPlayerService"
)

type PlayingSharedGame struct {
	LenderSteamID uint64 `json:"lender_steamid,string"`
}

func (a *API) IsPlayingSharedGame(steamID uint64, appID int) (res PlayingSharedGame, err error) {
	v := url.Values{}
	v.Add("steamid", strconv.FormatInt(int64(steamID), 10))
	v.Add("appid_playing", strconv.Itoa(appID))

	data := &struct {
		Response PlayingSharedGame
	}{}
	err = a.request(APIRequest{
		Service: SrvName,
		Method:  "IsPlayingSharedGame",
		Version: 1,
	}, v, &data)
	res = data.Response
	return
}
