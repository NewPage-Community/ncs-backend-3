package grpc

import (
	"backend/pkg/log"
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

const (
	ServiceName = "ncs-service-pass-user"
	ServiceAddr = ServiceName + ":2333"
)

var conn *grpc.ClientConn

func InitClient(target string, opts ...grpc.CallOption) UserClient {
	conn = rpc.Dial(context.Background(), target, nil)
	return NewUserClient(conn)
}

func Close() {
	if conn != nil {
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}
}
