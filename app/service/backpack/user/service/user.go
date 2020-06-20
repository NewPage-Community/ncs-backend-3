package service

import (
	pb "backend/app/service/backpack/user/api/grpc"
	"backend/app/service/backpack/user/model"
	"backend/pkg/ecode"
	"backend/pkg/json"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) GetItems(ctx context.Context, req *pb.GetItemsReq) (resp *pb.GetItemsResp, err error) {
	resp = &pb.GetItemsResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	var res *model.User
	var items []*pb.Item

	res, err = s.dao.Get(req.Uid)
	if err != nil {
		if ecode.GetError(err).Code != codes.NotFound {
			return
		}
		res, err = s.dao.Create(req.Uid)
		if err != nil {
			return
		}
	}
	err = json.Unmarshal(res.Items, &items)
	if err != nil {
		return
	}

	resp.Info = &pb.Info{
		Uid:   res.UID,
		Items: items,
	}
	return
}

func (s *Service) AddItem(ctx context.Context, req *pb.AddItemReq) (resp *pb.AddItemResp, err error) {
	resp = &pb.AddItemResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}
	if req.Item == nil {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Item")
		return
	}
	if req.Item.Id <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Item ID")
		return
	}

	err = s.dao.AddItem(req.Uid, model.Item{
		ID:       req.Item.Id,
		Amount:   req.Item.Amount,
		ExprTime: req.Item.ExprTime,
	}, req.Repeat)
	return
}

func (s *Service) RemoveItem(ctx context.Context, req *pb.RemoveItemReq) (resp *pb.RemoveItemResp, err error) {
	resp = &pb.RemoveItemResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}
	if req.Item == nil {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Item")
		return
	}
	if req.Item.Id <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Item ID")
		return
	}

	err = s.dao.RemoveItem(req.Uid, model.Item{
		ID:       req.Item.Id,
		Amount:   req.Item.Amount,
		ExprTime: req.Item.ExprTime,
	}, req.All)
	return
}
