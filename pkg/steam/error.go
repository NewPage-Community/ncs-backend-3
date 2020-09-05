package steam

import (
	"errors"
)

var (
	InvalidAPIKeyErr        = errors.New("invalid api key")
	InvalidRequestValuesErr = errors.New("invalid request values")
	TooManyRequestsErr      = errors.New("too many request to api")
	InternalServerErr       = errors.New("api internal error")
	ServiceUnavailableErr   = errors.New("api service unavailable")
)
