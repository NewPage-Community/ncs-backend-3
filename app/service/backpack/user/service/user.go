package service

import (
	itemsSrv "backend/app/service/backpack/items/api/grpc/v1"
	pb "backend/app/service/backpack/user/api/grpc/v1"
	"backend/app/service/backpack/user/model"
	"backend/pkg/ecode"
	"backend/pkg/rpc/gateway"
	"context"
	"google.golang.org/grpc/codes"
	"time"
)

func (s *Service) GetItems(ctx context.Context, req *pb.GetItemsReq) (resp *pb.GetItemsResp, err error) {
	resp = &pb.GetItemsResp{}

	// Web gateway force cover UID
	if id := gateway.GetID(ctx); id == "gateway-web" {
		req.Uid = s.jwt.PayloadFormContext(ctx).GetInt64("uid")
	}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	var res *model.User
	res, err = s.dao.Get(req.Uid)
	if err != nil {
		return
	}

	now := time.Now().Unix()
	for _, v := range *res.Items {
		if !v.IsExpired(now) {
			resp.Items = append(resp.Items, &pb.Item{
				Id:       v.ID,
				Amount:   v.Amount,
				ExprTime: v.ExprTime,
			})
		}
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

func (s *Service) Init(ctx context.Context, req *pb.InitReq) (resp *pb.InitResp, err error) {
	resp = &pb.InitResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	_, err = s.dao.Create(req.Uid)
	return
}

func (s *Service) HaveItem(ctx context.Context, req *pb.HaveItemReq) (resp *pb.HaveItemResp, err error) {
	resp = &pb.HaveItemResp{
		Have: false,
	}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}
	if req.Id == 0 {
		return
	}

	res, err := s.dao.Get(req.Uid)
	if err != nil {
		return
	}

	if v, ok := (*res.Items)[req.Id]; ok {
		if !v.IsExpired(time.Now().Unix()) {
			resp.Have = true
		}
	}
	return
}

func (s *Service) AddAllItems(ctx context.Context, req *pb.AddAllItemsReq) (resp *pb.AddAllItemsResp, err error) {
	resp = &pb.AddAllItemsResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	res, err := s.items.GetItems(ctx, &itemsSrv.GetItemsReq{})
	if err != nil {
		return
	}

	var items model.Items
	for _, v := range res.Items {
		items = append(items, &model.Item{
			ID:       v.Id,
			Amount:   0,
			ExprTime: 0,
		})
	}

	err = s.dao.AddItems(req.Uid, &items)
	return
}
