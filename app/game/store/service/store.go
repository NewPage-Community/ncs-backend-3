package service

import (
	pb "backend/app/game/store/api/grpc"
	itemsService "backend/app/service/backpack/items/api/grpc"
	userService "backend/app/service/backpack/user/api/grpc"
	moneyService "backend/app/service/user/money/api/grpc"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
	"math"
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
	price := int32(math.Floor(float64(item.Price) * float64(item.Discount)))
	_, err = s.money.Pay(ctx, &moneyService.PayReq{
		Uid:    req.Uid,
		Price:  price,
		Reason: "购买商品 " + item.Name,
	})
	// Payment failed
	if err != nil {
		return
	}
	_, err = s.user.AddItems(ctx, &userService.AddItemsReq{
		Uid: req.Uid,
		Items: []*userService.Item{{
			Id:     item.Id,
			Amount: 0,
			Length: 0,
		}},
	})
	if err != nil {
		// Add to backpack failed
		// Return money
		_, err = s.money.Give(ctx, &moneyService.GiveReq{
			Uid:    req.Uid,
			Money:  price,
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
