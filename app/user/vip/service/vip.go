package service

import (
	"backend/app/user/vip/api/grpc"
	"context"
)

func (s *Service) Info(ctx context.Context, req *grpc.InfoReq) (resp *grpc.InfoResp, err error) {
	return
}

func (s *Service) Grant(ctx context.Context, req *grpc.GrantReq) (resp *grpc.GrantResp, err error) {
	return
}

func (s *Service) AddPoint(ctx context.Context, req *grpc.AddPointReq) (resp *grpc.AddPointResp, err error) {
	return
}
