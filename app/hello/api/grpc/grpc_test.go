package grpc

import (
	"context"
	"testing"
)

type Test struct {
}

func (*Test) Say(ctx context.Context, req *SayReq) (resp *SayResp, err error) {
	resp = &SayResp{}
	return
}

func TestGrpc(t *testing.T) {
	s := InitServer("tcp", "0.0.0.0:2333", &Test{})
	c := InitClient("0.0.0.0:2333")
	c.Say(context.Background(), &SayReq{Message: ""})
	CloseClient()
	s.Stop()
}
