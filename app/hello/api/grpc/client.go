package grpc

import (
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

var conn *grpc.ClientConn
var client HelloClient

func InitClient(target string) {
	conn = rpc.Dial(target)
	client = NewHelloClient(conn)
}

func CloseClient() {
	if conn != nil {
		conn.Close()
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