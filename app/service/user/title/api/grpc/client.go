package grpc

import (
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

const (
	ServiceName = "ncs-service-user-title"
	ServiceAddr = ServiceName + ":2333"
)

var conn *grpc.ClientConn

func InitClient(target string, opts ...grpc.CallOption) TitleClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewTitleClient(conn)
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}
