package service

import (
	pb "backend/app/game/store/api/grpc"
	itemsService "backend/app/service/backpack/items/api/grpc"
	userService "backend/app/service/backpack/user/api/grpc"
	moneyService "backend/app/service/user/money/api/grpc"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

func (s *Service) BuyItem(ctx context.Context, req *pb.BuyItemReq) (resp *pb.BuyItemResp, err error) {
	resp = &pb.BuyItemResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	var item *itemsService.Item
	items, err := s.items.GetItems(ctx, &itemsService.GetItemsReq{})
	if err != nil {
		return
	}
	for _, v := range items.Items {
		if v.Id == req.ItemId {
			item = v
			break
		}
	}

	if item == nil {
		err = ecode.Errorf(codes.Unknown, "Can not found ItemID(%d)", req.ItemId)
		return
	}

	// Buy
	if req.Price > 0 {
		_, err = s.money.Pay(ctx, &moneyService.PayReq{
			Uid:    req.Uid,
			Price:  req.Price,
			Reason: "购买商品 " + item.Name,
		})
		// Payment failed
		if err != nil {
			return
		}
	}
	_, err = s.user.AddItems(ctx, &userService.AddItemsReq{
		Uid: req.Uid,
		Items: []*userService.Item{{
			Id:     item.Id,
			Amount: 0,
			Length: 0,
		}},
	})
	if err != nil && req.Price > 0 {
		// Add to backpack failed
		// Return money
		_, err = s.money.Give(ctx, &moneyService.GiveReq{
			Uid:    req.Uid,
			Money:  req.Price,
			Reason: "商品退款 " + item.Name,
		})
	}
	return
}

func (s *Service) HotSaleList(ctx context.Context, req *pb.HotSaleListReq) (resp *pb.HotSaleListResp, err error) {
	resp = &pb.HotSaleListResp{}

	for _, v := range s.hotSale.ItemIDs {
		resp.ItemsId = append(resp.ItemsId, v)
	}
	return
}

func (s *Service) SaleList(ctx context.Context, req *pb.SaleListReq) (resp *pb.SaleListResp, err error) {
	resp = &pb.SaleListResp{}

	var userItems *userService.GetItemsResp
	var userMoney *moneyService.GetResp
	var userItemsMap map[int32]bool
	if req.Uid > 0 {
		userItems, err = s.user.GetItems(ctx, &userService.GetItemsReq{Uid: req.Uid})
		if err != nil {
			return
		}
		userItemsMap = make(map[int32]bool)
		for _, v := range userItems.Items {
			// Unlimited item can not buy twice
			if v.ExprTime == 0 && v.Amount == 0 {
				userItemsMap[v.Id] = true
			}
		}
		userMoney, err = s.money.Get(ctx, &moneyService.GetReq{Uid: req.Uid})
		if err != nil {
			return
		}
		resp.Money = userMoney.Money
	}

	items, err := s.items.GetItems(ctx, &itemsService.GetItemsReq{})
	if err != nil {
		return
	}

	for _, v := range items.Items {
		have := false
		if req.Uid > 0 {
			_, have = userItemsMap[v.Id]
		}
		resp.Items = append(resp.Items, &pb.Item{
			ItemId:      v.Id,
			Name:        v.Name,
			Description: v.Description,
			Type:        v.Type,
			Price:       v.Price,
			Discount:    v.Discount,
			Available:   v.Available,
			AlreadyHave: have,
		})
	}
	return
}
