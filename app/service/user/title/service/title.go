package service

import (
	pb "backend/app/service/user/title/api/grpc"
	"backend/app/service/user/title/model"
	"context"
)

func (s *Service) GetTitle(ctx context.Context, req *pb.GetTitleReq) (resp *pb.GetTitleResp, err error) {
	resp = &pb.GetTitleResp{}

	res, err := s.dao.Title(req.Uid)
	resp.Info = &pb.Info{
		Uid:         res.UID,
		CustomTitle: res.CustomTitle,
		TitleType:   res.TitleType,
	}
	return
}

func (s *Service) SetTitle(ctx context.Context, req *pb.SetTitleReq) (resp *pb.SetTitleResp, err error) {
	resp = &pb.SetTitleResp{}

	err = s.dao.Update(&model.Title{
		UID:         req.Info.Uid,
		CustomTitle: req.Info.CustomTitle,
		TitleType:   req.Info.TitleType,
	})
	return
}
