package rpc

import (
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"net/http"
	"strconv"
	"time"
)

type Server struct {
	HealthCheck HealthFunc
	config      *ServerConfig
	grpcServer  *grpc.Server
}

type ServerConfig struct {
	Network           string
	Port              int
	Timeout           time.Duration
	IdleTimeout       time.Duration
	MaxLifeTime       time.Duration
	ForceCloseWait    time.Duration
	KeepAliveInterval time.Duration
	KeepAliveTimeout  time.Duration
}

var _defaultSerConf = &ServerConfig{
	Network:           "tcp",
	Port:              2333,
	Timeout:           time.Second,
	IdleTimeout:       time.Second * 60,
	MaxLifeTime:       time.Hour * 2,
	ForceCloseWait:    time.Second * 20,
	KeepAliveInterval: time.Second * 60,
	KeepAliveTimeout:  time.Second * 20,
}

func NewServer(conf *ServerConfig) (s *Server) {
	// Config
	if conf == nil {
		conf = _defaultSerConf
	}
	conf.Init()
	s = &Server{config: conf}
	s.initGrpc()
	return
}

func (conf *ServerConfig) Init() {
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
	if conf.Port == 0 {
		conf.Port = _defaultSerConf.Port
	}
	if conf.Network == "" {
		conf.Network = _defaultSerConf.Network
	}
}

func (s *Server) initGrpc() {
	// Options
	keepParam := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     s.config.IdleTimeout,
		MaxConnectionAgeGrace: s.config.ForceCloseWait,
		Time:                  s.config.KeepAliveInterval,
		Timeout:               s.config.KeepAliveTimeout,
		MaxConnectionAge:      s.config.MaxLifeTime,
	})
	tracer := grpc_opentracing.WithTracer(ot.GlobalTracer())
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(tracer),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(tracer),
		)),
		keepParam,
	}

	// Initialize
	s.grpcServer = grpc.NewServer(opts...)

	// Health
	if s.HealthCheck == nil {
		s.HealthCheck = func() bool {
			return true
		}
	}
	grpc_health_v1.RegisterHealthServer(s.grpcServer, s)
}

func (s *Server) GrpcServer() *grpc.Server {
	return s.grpcServer
}

func (s *Server) Serve() {
	// Register reflection service on gRPC server.
	reflection.Register(s.grpcServer)

	lis, err := net.Listen(s.config.Network, "0.0.0.0:"+strconv.Itoa(s.config.Port))
	if err != nil {
		panic(err)
	}

	// Serve
	go func() {
		err := s.grpcServer.Serve(lis)
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}
	}()
}

func (s *Server) Stop() {
	s.HealthCheck = func() bool {
		return false
	}
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
}
