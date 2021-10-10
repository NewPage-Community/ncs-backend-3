package service

import (
	pb "backend/app/game/store/api/grpc/v1"
	itemsService "backend/app/service/backpack/items/api/grpc/v1"
	userService "backend/app/service/backpack/user/api/grpc/v1"
	passService "backend/app/service/pass/user/api/grpc/v1"
	passModel "backend/app/service/pass/user/model"
	moneyService "backend/app/service/user/money/api/grpc/v1"
	vipService "backend/app/service/user/vip/api/grpc/v1"
	"backend/pkg/ecode"
	"context"
	"fmt"

	"google.golang.org/grpc/codes"
)

const (
	PassBoxID          = passModel.PassPointBoxID
	Pass1Price         = int32(6888)
	Pass2Price         = int32(9888)
	VIPMonthPrice      = int32(2000)
	VIPSeasonPrice     = int32(5400)
	VIPSemiannualPrice = int32(10560)
	VIPAnnualPrice     = int32(20000)
	Month              = int64(2592000)
	BuyVIPPoint        = int32(50)
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
	if !item.Available {
		err = ecode.Errorf(codes.Unknown, "item(%d) is not available", req.ItemId)
		return
	}

	amount := int32(0)
	if item.Id == PassBoxID {
		amount = 1
	}

	// Buy
	if req.Price > 0 {
		_, err = s.money.Pay(ctx, &moneyService.PayReq{
			Uid:    req.Uid,
			Price:  req.Price,
			Reason: "购买 " + item.Name,
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
			Amount: amount,
			Length: 0,
		}},
	})
	if err != nil && req.Price > 0 {
		// Add to backpack failed
		// Return money
		_, err = s.money.Give(ctx, &moneyService.GiveReq{
			Uid:    req.Uid,
			Money:  req.Price,
			Reason: "退款 " + item.Name,
		})
	}
	return
}

func (s *Service) HotSaleList(ctx context.Context, req *pb.HotSaleListReq) (resp *pb.HotSaleListResp, err error) {
	resp = &pb.HotSaleListResp{}
	resp.ItemsId = append(resp.ItemsId, s.hotSale.ItemIDs...)
	return
}

func (s *Service) SaleList(ctx context.Context, req *pb.SaleListReq) (resp *pb.SaleListResp, err error) {
	resp = &pb.SaleListResp{}

	var userItems *userService.GetItemsResp
	var userMoney *moneyService.GetResp
	var userPass *passService.InfoResp
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
		userPass, err = s.pass.Info(ctx, &passService.InfoReq{Uid: req.Uid})
		if err != nil {
			return
		}
		resp.PassType = userPass.Info.PassType
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

func (s *Service) BuyPass(ctx context.Context, req *pb.BuyPassReq) (resp *pb.BuyPassResp, err error) {
	resp = &pb.BuyPassResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}

	var price int32
	if req.Type == 1 {
		price = Pass1Price
	} else if req.Type == 2 {
		price = Pass2Price
	} else {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid Type(%d)", req.Type)
		return
	}

	// Buy
	_, err = s.money.Pay(ctx, &moneyService.PayReq{
		Uid:    req.Uid,
		Price:  price,
		Reason: "购买通行证高级版",
	})
	// Payment failed
	if err != nil {
		return
	}
	_, err = s.pass.UpgradePass(ctx, &passService.UpgradePassReq{
		Uid:  req.Uid,
		Type: req.Type,
	})
	if err != nil {
		// Return money
		_, err = s.money.Give(ctx, &moneyService.GiveReq{
			Uid:    req.Uid,
			Money:  price,
			Reason: "通行证高级版退款",
		})
	}
	return
}

func (s *Service) BuyVIP(ctx context.Context, req *pb.BuyVIPReq) (resp *pb.BuyVIPResp, err error) {
	resp = &pb.BuyVIPResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID(%d)", req.Uid)
		return
	}
	if req.Month <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Month should be larger than 0")
		return
	}

	var price int32
	switch req.Month {
	case 1:
		price = VIPMonthPrice
	case 3:
		price = VIPSeasonPrice
	case 6:
		price = VIPSemiannualPrice
	case 12:
		price = VIPAnnualPrice
	case 24:
		price = VIPAnnualPrice * 2
	case 36:
		price = VIPAnnualPrice * 3
	default:
		price = VIPMonthPrice * req.Month
	}

	// Buy
	_, err = s.money.Pay(ctx, &moneyService.PayReq{
		Uid:    req.Uid,
		Price:  price,
		Reason: fmt.Sprintf("续费 VIP %d个月", req.Month),
	})
	// Payment failed
	if err != nil {
		return
	}
	_, err = s.vip.Renewal(ctx, &vipService.RenewalReq{
		Uid:    req.Uid,
		Length: Month * int64(req.Month),
	})
	if err != nil {
		// Return money
		_, err = s.money.Give(ctx, &moneyService.GiveReq{
			Uid:    req.Uid,
			Money:  price,
			Reason: "VIP退款",
		})
	} else {
		_, err = s.vip.AddPoint(ctx, &vipService.AddPointReq{
			Uid:      req.Uid,
			AddPoint: BuyVIPPoint * req.Month,
		})
	}
	return
}
