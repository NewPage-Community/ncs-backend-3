package service

import (
	pb "backend/app/game/stats/api/grpc/v1"
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
	resp.Total, err = s.dao.Get(info)
	if err != nil {
		return
	}
	resp.Uid = req.Uid
	resp.Score = float32(info.Score)
	resp.Rank = info.Rank
	return
}

func (s *Service) GetAll(ctx context.Context, req *pb.GetAllReq) (resp *pb.GetAllResp, err error) {
	resp = &pb.GetAllResp{}

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

	res, err := s.dao.GetAll(&model.Stats{
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

func (s *Service) Gets(ctx context.Context, req *pb.GetsReq) (resp *pb.GetsResp, err error) {
	resp = &pb.GetsResp{}

	for _, v := range req.Stats {
		if v.StatsName == "" || v.Version == "" {
			continue
		}
		if v.Range == "" {
			v.Range = StatsAllRange
		}

		info := &model.Stats{
			UID:       v.Uid,
			Range:     v.Range,
			StatsName: v.StatsName,
			Version:   v.Version,
		}
		resp.Total, err = s.dao.Get(info)
		if err != nil {
			return
		}
		resp.Stats = append(resp.Stats, &pb.StatsInfo{
			Range:     info.Range,
			StatsName: info.StatsName,
			Version:   info.Version,
			Uid:       info.UID,
			Score:     float32(info.Score),
			Rank:      info.Rank,
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

func (s *Service) Sets(ctx context.Context, req *pb.SetsReq) (resp *pb.SetResp, err error) {
	resp = &pb.SetResp{}

	for _, v := range req.Stats {
		if v.StatsName == "" || v.Version == "" {
			continue
		}
		if v.Range == "" {
			v.Range = StatsAllRange
		}

		err = s.dao.Set(&model.Stats{
			UID:       v.Uid,
			Range:     v.Range,
			StatsName: v.StatsName,
			Version:   v.Version,
			Score:     float64(v.Score),
		})
		if err != nil {
			return
		}
	}
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

func (s *Service) Incrs(ctx context.Context, req *pb.IncrsReq) (resp *pb.IncrResp, err error) {
	resp = &pb.IncrResp{}

	for _, v := range req.Stats {
		if v.StatsName == "" || v.Version == "" {
			continue
		}
		if v.Range == "" {
			v.Range = StatsAllRange
		}

		err = s.dao.Incr(&model.Stats{
			UID:       v.Uid,
			Range:     v.Range,
			StatsName: v.StatsName,
			Version:   v.Version,
			Score:     float64(v.Increment),
		})
		if err != nil {
			return
		}
	}
	return
}

func (s *Service) GetPartly(ctx context.Context, req *pb.GetPartlyReq) (resp *pb.GetPartlyResp, err error) {
	resp = &pb.GetPartlyResp{}

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
	if req.Start < 1 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid start")
		return
	}
	if req.End < 1 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid end")
		return
	}

	res, total, err := s.dao.GetPartly(&model.Stats{
		Range:     req.Range,
		StatsName: req.StatsName,
		Version:   req.Version,
	}, req.Start, req.End)
	if err != nil {
		return
	}

	resp.Total = total
	for _, v := range res {
		resp.Data = append(resp.Data, &pb.GetResp{
			Uid:   v.UID,
			Score: float32(v.Score),
			Rank:  v.Rank,
		})
	}

	return
}
