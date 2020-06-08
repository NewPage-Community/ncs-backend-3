package grpc

import (
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

const ServiceName = "ncs-user-sign"

var conn *grpc.ClientConn

func InitClient(target string, opts ...grpc.CallOption) SignClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewSignClient(conn)
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}
