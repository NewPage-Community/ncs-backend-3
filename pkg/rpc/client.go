package rpc

import (
	"context"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/keepalive"
)

type ClientConfig struct {
	Dial                time.Duration
	KeepaliveInterval   time.Duration
	KeepaliveTimeout    time.Duration
	PermitWithoutStream bool

	// Request retry
	RetryMaxTime uint
	RetryCode    []codes.Code
}

var _defaultCliConf = &ClientConfig{
	Dial:                10 * time.Second,
	KeepaliveInterval:   10 * time.Second,
	KeepaliveTimeout:    1 * time.Second,
	PermitWithoutStream: true,
	RetryMaxTime:        3,
	RetryCode: []codes.Code{
		codes.DataLoss,
		codes.Unavailable,
	},
}

func Dial(ctx context.Context, target string, conf *ClientConfig) *grpc.ClientConn {
	// Config
	if conf == nil {
		conf = _defaultCliConf
	}
	conf.Init()

	// Options
	keepParam := grpc.WithKeepaliveParams(keepalive.ClientParameters{
		Time:                conf.KeepaliveInterval,
		Timeout:             conf.KeepaliveTimeout,
		PermitWithoutStream: conf.PermitWithoutStream,
	})
	retry := []grpc_retry.CallOption{
		grpc_retry.WithMax(conf.RetryMaxTime),
		grpc_retry.WithCodes(conf.RetryCode...),
	}
	tracer := grpc_opentracing.WithTracer(ot.GlobalTracer())
	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				grpc_opentracing.UnaryClientInterceptor(tracer),
				grpc_retry.UnaryClientInterceptor(retry...),
			)),
		grpc.WithStreamInterceptor(
			grpc_middleware.ChainStreamClient(
				grpc_opentracing.StreamClientInterceptor(tracer),
				grpc_retry.StreamClientInterceptor(retry...),
			),
		),
		grpc.WithInsecure(),
		keepParam,
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
	}

	// Initialize
	if conf.Dial > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, conf.Dial)
		defer cancel()
	}
	// Use dns with balancer
	dnsTarget := "dns:///" + target
	conn, err := grpc.DialContext(ctx, dnsTarget, opts...)
	if err != nil {
		panic(err)
	}
	return conn
}

func (conf *ClientConfig) Init() {
	if conf.Dial <= 0 {
		conf.Dial = _defaultCliConf.Dial
	}
	if conf.RetryMaxTime <= 0 {
		conf.RetryMaxTime = _defaultCliConf.RetryMaxTime
	}
	if len(conf.RetryCode) <= 0 {
		conf.RetryCode = _defaultCliConf.RetryCode
	}
}
