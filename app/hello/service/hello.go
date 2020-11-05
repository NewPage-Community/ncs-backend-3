package service

import(
	"context"
	pb "backend/app/hello/api/grpc/v1"
)

// Hello 实现 grpc 接口
func (s *Service) Hello(ctx context.Context, req *pb.HelloReq) (resp *pb.HelloResp, err error) {
	return &pb.HelloResp{
		Message: "Hello world!",
	}, nil
}