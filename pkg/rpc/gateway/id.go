package gateway

import (
	"context"
	"google.golang.org/grpc/metadata"
)

const (
	IDKey = "x-gateway-id"
)

func GetID(ctx context.Context) string {
	if md, ok := metadata.FromIncomingContext(ctx); ok {
		if id, ok := md[IDKey]; ok {
			return id[0]
		}
	}
	return ""
}

func InjectID(ctx context.Context, id string) context.Context {
	return metadata.NewOutgoingContext(ctx, metadata.Pairs(IDKey, id))
}
