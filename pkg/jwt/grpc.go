package jwt

import (
	"context"
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer "
)

// PayloadFormContext ...
func (c *JWT) PayloadFormContext(ctx context.Context) *Payload {
	token := TokenFromContext(ctx)
	if len(token) > 0 {
		// Verify token and get payload
		payload, err := c.GetTokenPayload(token)
		if err == nil {
			return &payload
		}
	}
	return new(Payload)
}

// TokenFromContext ...
func TokenFromContext(ctx context.Context) (token string) {
	data, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return
	}

	for _, v := range data[authorizationHeader] {
		if strings.HasPrefix(v, bearerPrefix) {
			token = v[len(bearerPrefix):]
		}
	}
	return
}
