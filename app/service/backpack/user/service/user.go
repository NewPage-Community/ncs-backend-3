package service

import (
	pb "backend/app/service/backpack/user/api/grpc"
	"backend/app/service/backpack/user/model"
	"backend/pkg/ecode"
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
	res, err = s.dao.Get(req.Uid)
	if err != nil {
		// not found -> create
		if ecode.GetError(err).Code != codes.NotFound {
			return
		}
		res, err = s.dao.Create(req.Uid)
		if err != nil {
			return
		}
	}

	for _, v := range *res.Items {
		resp.Items = append(resp.Items, &pb.Item{
			Id:       v.ID,
			Amount:   v.Amount,
			ExprTime: v.ExprTime,
		})
	}

	resp.Uid = res.UID
	return
}

func (s *Service) AddItems(ctx context.Context, req *pb.AddItemsReq) (resp *pb.AddItemsResp, err error) {
	resp = &pb.AddItemsResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}
	if len(req.Items) == 0 {
		return
	}

	var items model.Items
	for i := range req.Items {
		item := &model.Item{
			ID:       req.Items[i].Id,
			Amount:   req.Items[i].Amount,
			ExprTime: req.Items[i].Length,
		}
		if item.IsValid() {
			items = append(items, item)
		}
	}

	err = s.dao.AddItems(req.Uid, &items)
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
