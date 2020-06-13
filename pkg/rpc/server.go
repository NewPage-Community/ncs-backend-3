package rpc

import (
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	config      *ServerConfig
	grpcServer  *grpc.Server
	healthCheck *http.Server
}

type ServerConfig struct {
	Network           string
	GrpcPort          int
	HealthPort        int
	Timeout           time.Duration
	IdleTimeout       time.Duration
	MaxLifeTime       time.Duration
	ForceCloseWait    time.Duration
	KeepAliveInterval time.Duration
	KeepAliveTimeout  time.Duration
}

var _defaultSerConf = &ServerConfig{
	Network:           "tcp",
	GrpcPort:          2333,
	HealthPort:        23333,
	Timeout:           time.Duration(time.Second),
	IdleTimeout:       time.Duration(time.Second * 60),
	MaxLifeTime:       time.Duration(time.Hour * 2),
	ForceCloseWait:    time.Duration(time.Second * 20),
	KeepAliveInterval: time.Duration(time.Second * 60),
	KeepAliveTimeout:  time.Duration(time.Second * 20),
}

func NewServer(conf *ServerConfig) *Server {
	// Config
	conf = setSerConf(conf)
	return &Server{config: conf}
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
	if conf.GrpcPort == 0 {
		conf.GrpcPort = _defaultSerConf.GrpcPort
	}
	if conf.Network == "" {
		conf.Network = _defaultSerConf.Network
	}
	if conf.HealthPort == 0 {
		conf.HealthPort = _defaultSerConf.HealthPort
	}
	return conf
}

func (s *Server) Grpc(reg func(s *grpc.Server)) {
	// Options
	keepParam := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     time.Duration(s.config.IdleTimeout),
		MaxConnectionAgeGrace: time.Duration(s.config.ForceCloseWait),
		Time:                  time.Duration(s.config.KeepAliveInterval),
		Timeout:               time.Duration(s.config.KeepAliveTimeout),
		MaxConnectionAge:      time.Duration(s.config.MaxLifeTime),
	})
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(),
		)),
		keepParam,
	}

	// Network
	lis, err := net.Listen(s.config.Network, "0.0.0.0:"+strconv.Itoa(s.config.GrpcPort))
	if err != nil {
		panic(err)
	}

	// Initialize
	s.grpcServer = grpc.NewServer(opts...)

	// Register
	if reg != nil {
		reg(s.grpcServer)
	}

	// Register reflection service on gRPC server.
	reflection.Register(s.grpcServer)

	// grpc server
	go func() {
		if err := s.grpcServer.Serve(lis); err != nil {
			panic(err)
		}
	}()
}

func (s *Server) HealthCheck(health func() bool) {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", HealthCheck(health))
	s.healthCheck = &http.Server{
		Addr:    "0.0.0.0:" + strconv.Itoa(s.config.HealthPort),
		Handler: mux,
	}
	// health check
	go func() {
		err := s.healthCheck.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *Server) Stop() {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	if s.healthCheck != nil {
		ctx, _ := context.WithTimeout(context.Background(), time.Second*10)
		s.healthCheck.Shutdown(ctx)
	}
}
