package grpc

import (
	"backend/pkg/log"
	"backend/pkg/rpc"
	"context"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	client HelloClient
	opts   []grpc.CallOption
}

func InitClient(target string, opts ...grpc.CallOption) (c *Client) {
	c = &Client{}
	c.conn = rpc.Dial(context.Background(), target, nil)
	c.client = NewHelloClient(c.conn)
	c.opts = opts
	return
}

func (c *Client) Close() {
	if c.conn != nil {
		err := c.conn.Close()
		if err != nil {
			log.Error(err)
		}
	}
}

// --- API --- //

func (c *Client) Say(ctx context.Context, message string) (string, error) {
	res, err := c.client.Say(ctx, &SayReq{Message: message}, c.opts...)
	if err != nil {
		return "", err
	}
	return res.Message, nil
}
