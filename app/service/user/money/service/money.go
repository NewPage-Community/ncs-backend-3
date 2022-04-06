package service

import (
	pb "backend/app/service/user/money/api/grpc/v1"
	"backend/pkg/ecode"
	"context"

	"google.golang.org/grpc/codes"
)

func (s *Service) Get(ctx context.Context, req *pb.GetReq) (resp *pb.GetResp, err error) {
	resp = &pb.GetResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	res, err := s.dao.Get(req.Uid)
	if err != nil {
		return
	}
	resp.Uid = req.Uid
	resp.Money = res.RMB
	return
}

func (s *Service) Pay(ctx context.Context, req *pb.PayReq) (resp *pb.PayResp, err error) {
	resp = &pb.PayResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if req.Price <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Price(%d)", req.Price)
		return
	}

	err = s.dao.Pay(req.Uid, uint32(req.Price), req.Reason)
	if err != nil {
		return
	}

	return
}

func (s *Service) Give(ctx context.Context, req *pb.GiveReq) (resp *pb.GiveResp, err error) {
	resp = &pb.GiveResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if req.Money <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Money(%d)", req.Money)
		return
	}

	err = s.dao.Give(req.Uid, uint32(req.Money), req.Reason)
	if err != nil {
		return
	}

	return
}

func (s *Service) Records(ctx context.Context, req *pb.RecordsReq) (resp *pb.RecordsResp, err error) {
	resp = &pb.RecordsResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	res, err := s.dao.GetRecords(req.Uid, req.Days)
	if err != nil {
		return
	}

	for _, v := range *res {
		resp.Records = append(resp.Records, &pb.Record{
			Timestamp: v.CreatedAt.Unix(),
			Amount:    v.Amount,
			Reason:    v.Reason,
		})
	}
	return
}

func (s *Service) Gift(ctx context.Context, req *pb.GiftReq) (resp *pb.GiftResp, err error) {
	resp = &pb.GiftResp{}

	if req.Uid <= 0 || req.Target <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID or Target(%d, %d)", req.Uid, req.Target)
		return
	}
	if req.Money == 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Money(%d)", req.Money)
		return
	}

	err = s.dao.Gift(req.Uid, req.Target, req.Money)
	return
}
