package v1

import (
	"backend/pkg/log"
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

const (
	// ServiceName 这里修改服务名字
	ServiceName = "ncs-hello"
	ServiceAddr = ServiceName + ":2333"
)

var conn *grpc.ClientConn

// InitClient 初始化客户端
func InitClient(target string, opts ...grpc.CallOption) HelloClient {
	conn = rpc.Dial(context.Background(), target, nil)
	// 这里修改 grpc 接口
	return NewHelloClient(conn)
}

// Close 优雅关闭客户端
func Close() {
	if conn != nil {
		if err := conn.Close(); err != nil {
			log.Error(err)
		}
	}
}
