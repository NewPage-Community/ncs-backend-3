package grpc

import (
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

const (
	ServiceName = "ncs-user-vip"
	ServiceAddr = ServiceName + ":2333"
)

var conn *grpc.ClientConn

func InitClient(target string, opts ...grpc.CallOption) VIPClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewVIPClient(conn)
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}
