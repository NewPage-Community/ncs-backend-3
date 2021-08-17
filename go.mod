module backend

go 1.14

replace google.golang.org/grpc => google.golang.org/grpc v1.37.1

require (
	github.com/NewPage-Community/go-source-server-query v1.1.2
	github.com/alicebob/miniredis/v2 v2.11.4
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/go-redis/redis/v7 v7.4.0
	github.com/golang/glog v0.0.0-20210429001901-424d2337a529
	github.com/golang/mock v1.6.0
	github.com/gorilla/websocket v1.4.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/gunslinger23/kaiheila v0.0.0-20210128124554-d0443c147652
	github.com/json-iterator/go v1.1.10
	github.com/miRemid/amy v0.1.3
	github.com/opentracing/opentracing-go v1.1.0
	github.com/robfig/cron/v3 v3.0.0
	github.com/smartwalle/alipay/v3 v3.1.5
	github.com/smartystreets/goconvey v1.6.4
	github.com/spf13/viper v1.7.0
	github.com/uber/jaeger-client-go v2.23.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible
	github.com/wechatpay-apiv3/wechatpay-go v0.2.3
	github.com/yohcop/openid-go v1.0.0
	go.uber.org/zap v1.15.0
	golang.org/x/net v0.0.0-20210716203947-853a461950ff // indirect
	golang.org/x/oauth2 v0.0.0-20210615190721-d04028783cf1
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	golang.org/x/tools v0.1.5 // indirect
	google.golang.org/genproto v0.0.0-20210617175327-b9e0b3197ced
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gorm.io/datatypes v0.0.0-20200709131824-976937c55e2d
	gorm.io/driver/mysql v0.3.0
	gorm.io/gorm v1.20.0
)
