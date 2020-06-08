package rpc

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	ot "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"time"
)

type ClientConfig struct {
	Dial      time.Duration
	Timeout   time.Duration
	MaxRetry  uint
	RetryCode []codes.Code
}

var _defaultCliConf = &ClientConfig{
	Dial:     time.Duration(time.Second * 10),
	Timeout:  time.Duration(time.Millisecond * 250),
	MaxRetry: 3,
	RetryCode: []codes.Code{
		codes.Canceled,
		codes.DataLoss,
		codes.Unavailable,
	},
}

func Dial(ctx context.Context, target string, conf *ClientConfig) *grpc.ClientConn {
	// Config
	conf = setCliConf(conf)

	// Initialize
	if conf.Dial > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(conf.Dial))
		defer cancel()
	}
	conn, err := grpc.DialContext(ctx, target, getDialOption(conf)...)
	if err != nil {
		panic(err)
	}
	return conn
}

func getDialOption(conf *ClientConfig) []grpc.DialOption {
	// Options
	return []grpc.DialOption{
		grpc.WithUnaryInterceptor(
			grpc_opentracing.UnaryClientInterceptor(
				grpc_opentracing.WithTracer(ot.GlobalTracer()))),
		grpc.WithStreamInterceptor(
			grpc_opentracing.StreamClientInterceptor(
				grpc_opentracing.WithTracer(ot.GlobalTracer()))),
		grpc.WithUnaryInterceptor(
			grpc_retry.UnaryClientInterceptor(
				grpc_retry.WithMax(conf.MaxRetry),
				grpc_retry.WithCodes(conf.RetryCode...),
			)),
		grpc.WithInsecure(),
	}
}

func setCliConf(conf *ClientConfig) *ClientConfig {
	if conf == nil {
		conf = _defaultCliConf
	}
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
	return conf
}
