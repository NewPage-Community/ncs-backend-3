package grpc

import (
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

const ServiceName = "ncs-server"

var conn *grpc.ClientConn

func InitClient(target string, opts ...grpc.CallOption) ServerClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewServerClient(conn)
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}
