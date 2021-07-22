package rpc

import (
	"backend/pkg/conf"
	"backend/pkg/jwt"
	"backend/pkg/log"
	"context"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/opentracing/opentracing-go"
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
	config      *ServerConfig
	grpcServer  *grpc.Server
	healthCheck *http.Server
	healthFunc  func() bool
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
	JWT               jwt.JWT
}

var _defaultSerConf = ServerConfig{
	Network:           "tcp",
	GrpcPort:          2333,
	HealthPort:        23333,
	Timeout:           10 * time.Second,
	IdleTimeout:       60 * time.Second,
	MaxLifeTime:       2 * time.Hour,
	ForceCloseWait:    20 * time.Second,
	KeepAliveInterval: 60 * time.Second,
	KeepAliveTimeout:  20 * time.Second,
	JWT: jwt.JWT{
		ExpireTime: 86400,
		SecretKey:  "",
	},
}

func loadConf() *ServerConfig {
	c := &struct {
		Rpc *ServerConfig
	}{}
	*(c.Rpc) = _defaultSerConf
	conf.Load(c)
	return c.Rpc
}

func NewServer() *Server {
	// Config
	return &Server{config: loadConf()}
}

func (s *Server) Grpc(reg func(s *grpc.Server)) {
	// Options
	keepParam := grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle:     s.config.IdleTimeout,
		MaxConnectionAgeGrace: s.config.ForceCloseWait,
		Time:                  s.config.KeepAliveInterval,
		Timeout:               s.config.KeepAliveTimeout,
		MaxConnectionAge:      s.config.MaxLifeTime,
	})
	tracer := grpc_opentracing.WithTracer(opentracing.GlobalTracer())
	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			grpc_opentracing.UnaryServerInterceptor(tracer),
			jwt.UnaryServerInterceptor(&s.config.JWT),
		)),
		grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
			grpc_opentracing.StreamServerInterceptor(tracer),
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
