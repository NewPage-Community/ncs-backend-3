package service

import (
	pb "backend/app/service/auth/steam/api/grpc/v1"
	"context"
)

func (s *Service) SignInWithSteam(ctx context.Context, req *pb.SignInWithSteamReq) (resp *pb.SignInWithSteamResp, err error) {
	resp = &pb.SignInWithSteamResp{}

	return
}

func (s *Service) GetSteamToken(ctx context.Context, req *pb.GetSteamTokenReq) (resp *pb.GetSteamTokenResp, err error) {
	resp = &pb.GetSteamTokenResp{}

	return
}

func (s *Service) GetUID(ctx context.Context, req *pb.GetUIDReq) (resp *pb.GetUIDResp, err error) {
	resp = &pb.GetUIDResp{}

	return
}
