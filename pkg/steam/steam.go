package steam

import (
	"backend/pkg/json"
	"fmt"
	"net/http"
	"net/url"
)

const (
	APIUrlDefault = "https://api.steampowered.com"
)

type APIClient interface {
	IsPlayingSharedGame(steamID uint64, appID int) (res PlayingSharedGame, err error)
}

type API struct {
	APIUrl string
	APIKey string
}

type APIRequest struct {
	Service string
	Method  string
	Version int
}

func NewAPIClient(url, key string) *API {
	if len(url) == 0 {
		url = APIUrlDefault
	}

	return &API{
		APIUrl: url,
		APIKey: key,
	}
}

func (a *API) request(req APIRequest, values url.Values, v interface{}) error {
	if values == nil {
		return InvalidRequestValuesErr
	}
	values.Add("format", "json")
	values.Add("key", a.APIKey)

	url := fmt.Sprintf("%s/%s/%s/v%d/?%s", a.APIUrl, req.Service, req.Method, req.Version, values.Encode())
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := APICodeError(resp.StatusCode); err != nil {
		return err
	}

	d := json.NewDecoder(resp.Body)
	return d.Decode(v)
}
