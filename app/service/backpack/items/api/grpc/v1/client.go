package v1

import (
	"backend/pkg/log"
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

const (
	ServiceName = "ncs-service-backpack-items"
	ServiceAddr = ServiceName + ":2333"
)

var conn *grpc.ClientConn

func InitClient(target string, opts ...grpc.CallOption) ItemsClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewItemsClient(conn)
}

func Close() {
	if conn != nil {
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}
}
