package service

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Service struct {
	opts   []grpc.DialOption
	mux    *runtime.ServeMux
	ctx    context.Context
	cancel context.CancelFunc
}

func NewService(mux *runtime.ServeMux, opts []grpc.DialOption) *Service {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	return &Service{
		opts:   opts,
		mux:    mux,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (s *Service) Register() (err error) {
	err = s.regUserService()
	err = s.regServerService()
	return
}
