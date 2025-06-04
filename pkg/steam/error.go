package steam

import (
	"errors"
)

var (
	ErrInvalidAPIKey        = errors.New("invalid api key")
	ErrInvalidRequestValues = errors.New("invalid request values")
	ErrTooManyRequests      = errors.New("too many request to api")
	ErrInternalServer       = errors.New("api internal error")
	ErrServiceUnavailable   = errors.New("api service unavailable")
)
