package steam

import (
	"fmt"
	. "net/http"
)

func APICodeError(code int) error {
	switch code {
	case StatusOK:
		return nil
	case StatusBadRequest:
		return InvalidRequestValuesErr
	case StatusUnauthorized, StatusForbidden:
		return InvalidAPIKeyErr
	case StatusTooManyRequests:
		return TooManyRequestsErr
	case StatusInternalServerError:
		return InternalServerErr
	case StatusServiceUnavailable:
		return ServiceUnavailableErr
	default:
		return fmt.Errorf("api unknow status code: %d", code)
	}
}
