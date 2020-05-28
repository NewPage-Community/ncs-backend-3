# NCS 微服务后端

## 生成相关文件

```shell script
# grpc
protoc --gofast_out=plugins=grpc:. *.proto

# mock
mockgen -source=xxx.go -destination=xxx.mock.go -package xxx
```

## 编译

本仓库采用Bazel来进行构建，请自行安装环境。

```shell script
bazel build //...
bazel test //...
```

## 提交

请本地编译测试完毕后再提交，以免阻塞master的构建

## 规范

### Golang

1. 变量+函数采用驼峰命名方式(ServerID）

### 数据库

1. ORM模型命名采用驼峰(ServerID）
2. 字段命名采用(server_id)

> 使用ORM模型的命名在sql语句中将会自动转换为规则2的命名规则

### JSON

1. 字段命名采用(server_id)

### Protocol Buffers

1. 字段命名采用(server_id)
2. 请求消息命名(xxxxReq)
3. 响应消息命名(xxxxResp)