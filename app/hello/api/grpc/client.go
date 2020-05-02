package grpc

import (
	"backend/pkg/log"
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var client HelloClient

func InitClient(target string) {
	conn = rpc.Dial(context.Background(), target, nil)
	client = NewHelloClient(conn)
}

func CloseClient() {
	if conn != nil {
		err := conn.Close()
		if err != nil {
			log.Error(err)
		}
	}
}

// -- API -- //

func CallSay(ctx context.Context, message string) (string, error) {
	resp, err := client.Say(ctx, &SayReq{Message: message})
	if err != nil {
		return "", err
	}
	return resp.Message, nil
}
