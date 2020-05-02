package rpc

import (
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	ot "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"net"
	"time"
)

type ServerConfig struct {
	Network           string
	Addr              string
	Timeout           time.Duration
	IdleTimeout       time.Duration
	MaxLifeTime       time.Duration
	ForceCloseWait    time.Duration
	KeepAliveInterval time.Duration
	KeepAliveTimeout  time.Duration
	RegFunc           func(s *grpc.Server)
}

var _defaultSerConf = &ServerConfig{
	Network:           "tcp",
	Addr:              "0.0.0.0:2333",
	Timeout:           time.Duration(time.Second),
	IdleTimeout:       time.Duration(time.Second * 60),
	MaxLifeTime:       time.Duration(time.Hour * 2),
	ForceCloseWait:    time.Duration(time.Second * 20),
	KeepAliveInterval: time.Duration(time.Second * 60),
	KeepAliveTimeout:  time.Duration(time.Second * 20),
}

func NewServer(conf *ServerConfig) *grpc.Server {
	// Config
	conf = setSerConf(conf)

	// Options
	keepParam := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(conf.IdleTimeout),
		MaxConnectionAgeGrace: time.Duration(conf.ForceCloseWait),
		Time:                  time.Duration(conf.KeepAliveInterval),
		Timeout:               time.Duration(conf.KeepAliveTimeout),
		MaxConnectionAge:      time.Duration(conf.MaxLifeTime),
	})
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(
			otgrpc.OpenTracingServerInterceptor(ot.GlobalTracer())),
		grpc.StreamInterceptor(
			otgrpc.OpenTracingStreamServerInterceptor(ot.GlobalTracer())),
		keepParam,
	}

	// Network
	lis, err := net.Listen(conf.Network, conf.Addr)
	if err != nil {
		panic(err)
	}

	// Initialize
	s := grpc.NewServer(opts...)

	// Register
	if conf.RegFunc != nil {
		conf.RegFunc(s)
	}

	// Serve
	go func() {
		if err := s.Serve(lis); err != nil {
			panic(err)
		}
	}()
	return s
}

func setSerConf(conf *ServerConfig) *ServerConfig {
	if conf == nil {
		conf = _defaultSerConf
	}
	if conf.Timeout <= 0 {
		conf.Timeout = _defaultSerConf.Timeout
	}
	if conf.IdleTimeout <= 0 {
		conf.IdleTimeout = _defaultSerConf.IdleTimeout
	}
	if conf.MaxLifeTime <= 0 {
		conf.MaxLifeTime = _defaultSerConf.MaxLifeTime
	}
	if conf.ForceCloseWait <= 0 {
		conf.ForceCloseWait = _defaultSerConf.ForceCloseWait
	}
	if conf.KeepAliveInterval <= 0 {
		conf.KeepAliveInterval = _defaultSerConf.KeepAliveInterval
	}
	if conf.KeepAliveTimeout <= 0 {
		conf.KeepAliveTimeout = _defaultSerConf.KeepAliveTimeout
	}
	if conf.Addr == "" {
		conf.Addr = _defaultSerConf.Addr
	}
	if conf.Network == "" {
		conf.Network = _defaultSerConf.Network
	}
	return conf
}
