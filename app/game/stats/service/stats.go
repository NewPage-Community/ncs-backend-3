package service

import (
	pb "backend/app/game/stats/api/grpc"
	"backend/app/game/stats/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

const (
	StatsAllRange = "all"
)

func (s *Service) Get(ctx context.Context, req *pb.GetReq) (resp *pb.GetResp, err error) {
	resp = &pb.GetResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if req.StatsName == "" {
		err = ecode.Errorf(codes.InvalidArgument, "Empty stats name")
		return
	}
	if req.Version == "" {
		err = ecode.Errorf(codes.InvalidArgument, "Empty version")
		return
	}
	if req.Range == "" {
		req.Range = StatsAllRange
	}

	info := &model.Stats{
		UID:       req.Uid,
		Range:     req.Range,
		StatsName: req.StatsName,
		Version:   req.Version,
	}
	err = s.dao.Get(info)
	if err != nil {
		return
	}
	resp.Uid = req.Uid
	resp.Score = float32(info.Score)
	resp.Rank = info.Rank
	return
}

func (s *Service) Gets(ctx context.Context, req *pb.GetsReq) (resp *pb.GetsResp, err error) {
	resp = &pb.GetsResp{}

	if req.StatsName == "" {
		err = ecode.Errorf(codes.InvalidArgument, "Empty stats name")
		return
	}
	if req.Version == "" {
		err = ecode.Errorf(codes.InvalidArgument, "Empty version")
		return
	}
	if req.Range == "" {
		req.Range = StatsAllRange
	}

	res, err := s.dao.Gets(&model.Stats{
		Range:     req.Range,
		StatsName: req.StatsName,
		Version:   req.Version,
	})
	if err != nil {
		return
	}
	for _, v := range res {
		resp.Stats = append(resp.Stats, &pb.GetResp{
			Uid:   v.UID,
			Score: float32(v.Score),
			Rank:  v.Rank,
		})
	}
	return
}

func (s *Service) Set(ctx context.Context, req *pb.SetReq) (resp *pb.SetResp, err error) {
	resp = &pb.SetResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if req.StatsName == "" {
		err = ecode.Errorf(codes.InvalidArgument, "Empty stats name")
		return
	}
	if req.Version == "" {
		err = ecode.Errorf(codes.InvalidArgument, "Empty version")
		return
	}
	if req.Range == "" {
		req.Range = StatsAllRange
	}

	err = s.dao.Set(&model.Stats{
		UID:       req.Uid,
		Range:     req.Range,
		StatsName: req.StatsName,
		Version:   req.Version,
		Score:     float64(req.Score),
	})
	return
}

func (s *Service) Incr(ctx context.Context, req *pb.IncrReq) (resp *pb.IncrResp, err error) {
	resp = &pb.IncrResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if req.StatsName == "" {
		err = ecode.Errorf(codes.InvalidArgument, "Empty stats name")
		return
	}
	if req.Version == "" {
		err = ecode.Errorf(codes.InvalidArgument, "Empty version")
		return
	}
	if req.Range == "" {
		req.Range = StatsAllRange
	}

	err = s.dao.Incr(&model.Stats{
		UID:       req.Uid,
		Range:     req.Range,
		StatsName: req.StatsName,
		Version:   req.Version,
		Score:     float64(req.Increment),
	})
	return
}
