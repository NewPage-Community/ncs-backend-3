module backend

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.37.1

require (
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/NewPage-Community/go-source-server-query v1.1.3
	github.com/alicebob/miniredis/v2 v2.15.1
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v8 v8.11.3
	github.com/golang/glog v1.0.0
	github.com/golang/mock v1.6.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.6.0
	github.com/lonelyevil/khl v0.0.14
	github.com/miRemid/amy v0.1.3
	github.com/opentracing/opentracing-go v1.2.0
	github.com/pkg/errors v0.9.1 // indirect
	github.com/robfig/cron/v3 v3.0.1
	github.com/smartwalle/alipay/v3 v3.1.6
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.9.0
	github.com/uber/jaeger-client-go v2.29.1+incompatible
	github.com/uber/jaeger-lib v2.4.1+incompatible
	github.com/wechatpay-apiv3/wechatpay-go v0.2.7
	github.com/yohcop/openid-go v1.0.0
	github.com/yuin/gopher-lua v0.0.0-20210529063254-f4c35e4016d9 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519 // indirect
	golang.org/x/net v0.0.0-20210924151903-3ad01bbaa167 // indirect
	golang.org/x/oauth2 v0.0.0-20210819190943-2bc19b11175f
	golang.org/x/sys v0.0.0-20210925032602-92d5a993a665 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20210924002016-3dee208752a0
	google.golang.org/grpc v1.41.0
	google.golang.org/protobuf v1.27.1
	gorm.io/datatypes v1.0.2
	gorm.io/driver/mysql v1.1.2
	gorm.io/gorm v1.21.15
)
