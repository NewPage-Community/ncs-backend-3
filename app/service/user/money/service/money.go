package service

import (
	pb "backend/app/service/user/money/api/grpc"
	"backend/pkg/ecode"
	"backend/pkg/log"
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

	err = s.dao.Pay(req.Uid, req.Price)
	if err != nil {
		return
	}

	recordErr := s.dao.AddRecord(req.Uid, -req.Price, req.Reason)
	if recordErr != nil {
		log.Warn(recordErr)
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

	err = s.dao.Give(req.Uid, req.Money)
	if err != nil {
		return
	}

	recordErr := s.dao.AddRecord(req.Uid, req.Money, req.Reason)
	if recordErr != nil {
		log.Warn(recordErr)
	}
	return
}

func (s *Service) Records(ctx context.Context, req *pb.RecordsReq) (resp *pb.RecordsResp, err error) {
	resp = &pb.RecordsResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	res, err := s.dao.GetRecords(req.Uid)
	if err != nil {
		return
	}

	for _, v := range *res {
		resp.Records = append(resp.Records, &pb.Record{
			Timestamp: v.Timestamp,
			Amount:    v.Amount,
			Reason:    v.Reason,
		})
	}
	return
}
