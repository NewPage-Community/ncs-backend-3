package service

import (
	pb "backend/app/service/user/title/api/grpc"
	"backend/app/service/user/title/model"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) GetTitle(ctx context.Context, req *pb.GetTitleReq) (resp *pb.GetTitleResp, err error) {
	resp = &pb.GetTitleResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	res, err := s.dao.Title(req.Uid)
	if err != nil {
		if ecode.GetError(err).Code != codes.NotFound {
			return
		}
		// not found -> return empty
		res = &model.Title{
			UID:         req.Uid,
			CustomTitle: "",
			TitleType:   0,
		}
		err = s.dao.Create(res)
	}
	resp.Info = &pb.Info{
		Uid:         res.UID,
		CustomTitle: res.CustomTitle,
		TitleType:   res.TitleType,
	}
	return
}

func (s *Service) SetTitle(ctx context.Context, req *pb.SetTitleReq) (resp *pb.SetTitleResp, err error) {
	resp = &pb.SetTitleResp{}

	if req.Info == nil {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Info")
		return
	}

	if req.Info.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Info.Uid)
		return
	}

	err = s.dao.Update(&model.Title{
		UID:         req.Info.Uid,
		CustomTitle: req.Info.CustomTitle,
		TitleType:   req.Info.TitleType,
	})
	return
}
