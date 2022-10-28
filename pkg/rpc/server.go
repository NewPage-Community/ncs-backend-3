package rpc

import (
	"backend/pkg/log"
	"context"
	"net"
	"net/http"
	"strconv"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	config      *ServerConfig
	grpcServer  *grpc.Server
	healthCheck *http.Server
	healthFunc  func() bool
}

type ServerConfig struct {
	Network               string
	GrpcPort              int
	HealthPort            int
	MaxConnectionIdle     time.Duration
	MaxConnectionAge      time.Duration
	MaxConnectionAgeGrace time.Duration
	KeepAliveInterval     time.Duration
	KeepAliveTimeout      time.Duration
}

var _defaultSerConf = &ServerConfig{
	Network:               "tcp",
	GrpcPort:              2333,
	HealthPort:            23333,
	MaxConnectionIdle:     60 * time.Second,
	MaxConnectionAge:      2 * time.Hour,
	MaxConnectionAgeGrace: 20 * time.Second,
	KeepAliveInterval:     60 * time.Second,
	KeepAliveTimeout:      20 * time.Second,
}

func NewServer(conf *ServerConfig) *Server {
	// Config
	if conf == nil {
		conf = _defaultSerConf
	}
	conf.Init()
	return &Server{config: conf}
}

func (conf *ServerConfig) Init() {
	if conf.MaxConnectionIdle <= 0 {
		conf.MaxConnectionIdle = _defaultSerConf.MaxConnectionIdle
	}
	if conf.MaxConnectionAge <= 0 {
		conf.MaxConnectionAge = _defaultSerConf.MaxConnectionAge
	}
	if conf.MaxConnectionAgeGrace <= 0 {
		conf.MaxConnectionAgeGrace = _defaultSerConf.MaxConnectionAgeGrace
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
}

func (s *Server) Grpc(reg func(s *grpc.Server)) {
	// Options
	keepParam := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     s.config.MaxConnectionIdle,
		MaxConnectionAge:      s.config.MaxConnectionAge,
		MaxConnectionAgeGrace: s.config.MaxConnectionAgeGrace,
		Time:                  s.config.KeepAliveInterval,
		Timeout:               s.config.KeepAliveTimeout,
	})
	keepPolicy := grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
		PermitWithoutStream: true,
	})
	tracer := grpc_opentracing.WithTracer(opentracing.GlobalTracer())
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(tracer),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(tracer),
		)),
		keepParam,
		keepPolicy,
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

	// Health
	grpc_health_v1.RegisterHealthServer(s.grpcServer, s)

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
	s.healthFunc = health
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", HttpHealthHandler(s.healthFunc))
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
	s.healthFunc = func() bool {
		return false
	}
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	if s.healthCheck != nil {
		if err := s.healthCheck.Shutdown(context.Background()); err != nil {
			log.Error(err)
		}
	}
}
