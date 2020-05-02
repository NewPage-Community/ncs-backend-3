package service

import (
	"backend/app/hello/api/grpc"
	"context"
)

type Hello struct {
}

func (*Hello) Say(ctx context.Context, req *grpc.SayReq) (resp *grpc.SayResp, err error) {
	resp = &grpc.SayResp{Message: "Hello " + req.Message + "!"}
	return
}
