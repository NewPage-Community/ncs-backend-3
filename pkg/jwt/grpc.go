package jwt

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"strings"
)

const (
	authorizationHeader = "authorization"
	bearerPrefix        = "Bearer "
)

type jwtPayload struct{}

func (c *JWT) GetPayload(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	// Get token
	token := TokenFromContext(ctx)
	// Skip when do not have token
	if len(token) == 0 {
		return handler(ctx, req)
	}

	// Valid token and get payload
	payload, err := c.GetTokenPayload(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "token not valid: %v", err)
	}

	// Call handler
	return handler(ContextWithPayload(ctx, payload), req)
}

// UnaryServerInterceptor ...
func UnaryServerInterceptor(c *JWT) grpc.UnaryServerInterceptor {
	return c.GetPayload
}

// ContextWithPayload ...
func ContextWithPayload(ctx context.Context, payload interface{}) context.Context {
	return context.WithValue(ctx, jwtPayload{}, payload)
}

// PayloadFormContext ...
func PayloadFormContext(ctx context.Context) *Payload {
	payload, ok := ctx.Value(jwtPayload{}).(Payload)
	if !ok {
		payload = make(Payload)
	}
	return &payload
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
