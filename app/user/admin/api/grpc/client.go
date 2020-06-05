package grpc

import (
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

func InitClient(target string, opts ...grpc.CallOption) AdminClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewAdminClient(conn)
}

func Close() {
	if conn != nil {
		conn.Close()
	}
}
