package grpc

import (
	"backend/pkg/log"
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn

func InitClient(target string, opts ...grpc.CallOption) HelloClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewHelloClient(conn)
}

func CloseClient() {
	if conn != nil {
		err := conn.Close()
		if err != nil {
			log.Error(err)
		}
	}
}
