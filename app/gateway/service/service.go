package service

import (
	"context"
	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type regHandler func(ctx context.Context, mux *runtime.ServeMux, endpoint string, opts []grpc.DialOption) (err error)

type gateway struct {
	handlder regHandler
	endpoint string
	cancel   context.CancelFunc
}

type Service struct {
	opts     []grpc.DialOption
	mux      *runtime.ServeMux
	gateways []gateway
}

func (gw *gateway) regHandler(mux *runtime.ServeMux, opts []grpc.DialOption) error {
	ctx, cancel := context.WithCancel(context.Background())
	gw.cancel = cancel
	err := gw.handlder(ctx, mux, gw.endpoint, opts)
	return err
}

func NewService(mux *runtime.ServeMux, opts []grpc.DialOption) *Service {
	return &Service{
		opts:     opts,
		mux:      mux,
		gateways: make([]gateway, 0),
	}
}

func (s *Service) addService() {
	s.gateways = append(s.gateways, regServerService()...)
	s.gateways = append(s.gateways, regUserService()...)
}

func (s *Service) Register() {
	s.addService()
	for _, gw := range s.gateways {
		err := gw.regHandler(s.mux, s.opts)
		if err != nil {
			glog.Fatal(err)
		}
	}
}

func (s *Service) Close() {
	for _, gw := range s.gateways {
		if gw.cancel != nil {
			gw.cancel()
		}
	}
}
