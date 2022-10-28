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
	Time                time.Duration
	Timeout             time.Duration
	PermitWithoutStream bool
	MaxRetry            uint
	RetryCode           []codes.Code
}

var _defaultCliConf = &ClientConfig{
	Dial:                10 * time.Second,
	Time:                10 * time.Second,
	Timeout:             1 * time.Second,
	PermitWithoutStream: true,
	MaxRetry:            3,
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
		Time:                conf.Time,
		Timeout:             conf.Timeout,
		PermitWithoutStream: conf.PermitWithoutStream,
	})
	retry := []grpc_retry.CallOption{
		grpc_retry.WithMax(conf.MaxRetry),
		grpc_retry.WithCodes(conf.RetryCode...),
		grpc_retry.WithPerRetryTimeout(conf.Timeout),
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
	if conf.Timeout <= 0 {
		conf.Timeout = _defaultCliConf.Timeout
	}
	if conf.MaxRetry <= 0 {
		conf.MaxRetry = _defaultCliConf.MaxRetry
	}
	if len(conf.RetryCode) <= 0 {
		conf.RetryCode = _defaultCliConf.RetryCode
	}
}
