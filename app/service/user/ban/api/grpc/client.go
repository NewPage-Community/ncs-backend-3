package grpc

import (
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

const (
	ServiceName = "ncs-service-user-ban"
	ServiceAddr = ServiceName + ":2333"
)

var conn *grpc.ClientConn

func InitClient(target string, opts ...grpc.CallOption) BanClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewBanClient(conn)
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}