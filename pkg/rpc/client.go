package rpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-opentracing/go/otgrpc"
	ot "github.com/opentracing/opentracing-go"
	"google.golang.org/grpc"
	"time"
)

type ClientConfig struct {
	Dial    time.Duration
	Timeout time.Duration
}

var _defaultCliConf = &ClientConfig{
	Dial:    time.Duration(time.Second * 10),
	Timeout: time.Duration(time.Millisecond * 250),
}

func Dial(ctx context.Context, target string, conf *ClientConfig) *grpc.ClientConn {
	// Config
	conf = setCliConf(conf)

	// Options
	opts := []grpc.DialOption{
		grpc.WithUnaryInterceptor(
			otgrpc.OpenTracingClientInterceptor(ot.GlobalTracer())),
		grpc.WithStreamInterceptor(
			otgrpc.OpenTracingStreamClientInterceptor(ot.GlobalTracer())),
		grpc.WithInsecure(),
	}

	// Initialize
	if conf.Dial > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(conf.Dial))
		defer cancel()
	}
	conn, err := grpc.DialContext(ctx, target, opts...)
	if err != nil {
		panic(err)
	}
	return conn
}

func setCliConf(conf *ClientConfig) *ClientConfig {
	if conf == nil {
		conf = _defaultCliConf
	}
	if conf.Dial <= 0 {
		conf.Timeout = _defaultCliConf.Dial
	}
	if conf.Timeout <= 0 {
		conf.Timeout = _defaultCliConf.Timeout
	}
	return conf
}
