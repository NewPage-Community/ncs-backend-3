# NCS 微服务后端

使用简易自封装的微服务框架，能快速开发相关业务🎯

使用k8s进行服务发现，搭配⛵️Istio进行治理

使用CI/CD全自动化构建和部署，拯救懒人🤤

## 目录

| 目录 | 描述 |
| ----  | ----- |
| app   | 服务 |
| pkg   | 框架 |
| build | 构建文件  |
| chart | Helm chart仓库 |
| google | 谷歌GRPC文件 |

## 编译

本仓库采用Bazel来进行构建，已使用Docker提前部署所有环境，开箱即用。

```bash
# all in one
make all
```

## 提交

请本地编译测试完毕后再提交，以免阻塞master的构建

## 规范

### Golang

1. 变量+函数采用驼峰命名方式(ServerID）

### 数据库表

1. ORM模型命名采用驼峰(ServerID）
2. 字段命名采用(server_id)

> 使用ORM模型的命名在sql语句中将会自动转换为规则2的命名规则

### JSON

1. 字段命名采用(server_id)

### Protocol Buffers

1. 字段命名采用(server_id)
2. 请求消息命名(xxxxReq)
3. 响应消息命名(xxxxResp)