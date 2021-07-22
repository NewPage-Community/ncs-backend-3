package gateway

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

const (
	IDKey = "x-gateway-id"
)

func GetID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return ""
	}
	id := md.Get(IDKey)
	if len(id) == 0 {
		return ""
	}
	return id[0]
}

func InjectID(ctx context.Context, id string) {
	_ = grpc.SetHeader(ctx, metadata.Pairs(IDKey, id))
}
