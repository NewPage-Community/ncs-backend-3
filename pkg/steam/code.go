package steam

import (
	"fmt"
	"net/http"
)

func APICodeError(code int) error {
	switch code {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return ErrInvalidRequestValues
	case http.StatusUnauthorized, http.StatusForbidden:
		return ErrInvalidAPIKey
	case http.StatusTooManyRequests:
		return ErrTooManyRequests
	case http.StatusInternalServerError:
		return ErrInternalServer
	case http.StatusServiceUnavailable:
		return ErrServiceUnavailable
	default:
		return fmt.Errorf("api unknow status code: %d", code)
	}
}
