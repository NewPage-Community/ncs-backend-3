package steam

import (
	"backend/pkg/json"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

const (
	APIUrlDefault  = "https://api.steampowered.com"
	TimeoutDefault = 1
)

type APIClient interface {
	IsPlayingSharedGame(steamID uint64, appID int) (res PlayingSharedGame, err error)
}

type API struct {
	APIUrl string
	APIKey string
	// API request timeout (sec)
	Timeout int
	// Proxy
	HttpProxy string
}

type APIRequest struct {
	Service string
	Method  string
	Version int
}

func NewAPIClient(url, key string, timeout int) *API {
	if len(url) == 0 {
		url = APIUrlDefault
	}
	if timeout == 0 {
		timeout = TimeoutDefault
	}

	return &API{
		APIUrl:  url,
		APIKey:  key,
		Timeout: timeout,
	}
}

func (a *API) request(req APIRequest, values url.Values, v interface{}) error {
	if values == nil {
		return InvalidRequestValuesErr
	}
	values.Add("format", "json")
	values.Add("key", a.APIKey)

	apiURL := fmt.Sprintf("%s/%s/%s/v%d/?%s", a.APIUrl, req.Service, req.Method, req.Version, values.Encode())
	client := http.Client{Timeout: time.Second}
	if len(a.HttpProxy) > 0 {
		client.Transport = &http.Transport{
			Proxy: func(_ *http.Request) (*url.URL, error) {
				return url.Parse("http://" + a.HttpProxy)
			},
		}
	}
	resp, err := client.Get(apiURL)
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
