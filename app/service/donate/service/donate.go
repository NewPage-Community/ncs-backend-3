package service

import (
	pb "backend/app/service/donate/api/grpc/v1"
	"backend/app/service/donate/event"
	"backend/app/service/donate/model"
	accountService "backend/app/service/user/account/api/grpc/v1"
	moneyService "backend/app/service/user/money/api/grpc/v1"
	"backend/pkg/log"
	"context"
	"time"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	TradeSubject        = "捐助NewPage社区"
	TradeBody           = "捐助NewPage社区-%d"
	Timeout             = "5m"
	DonateRMBExchange   = 100
	DonateCheckSchedule = "@every 5s"
)

func (s *Service) CreateDonate(ctx context.Context, req *pb.CreateDonateReq) (resp *pb.CreateDonateResp, err error) {
	resp = &pb.CreateDonateResp{}

	if req.SteamId == 0 || req.Amount == 0 {
		err = status.Error(codes.InvalidArgument, "Invalid args")
		return
	}

	// Get UID
	acc, err := s.account.UID(context.Background(), &accountService.UIDReq{
		SteamId: req.SteamId,
	})
	if err != nil {
		return
	}

	// Get trade no
	outTradeNo, err := s.dao.CreateDonate(acc.Uid, req.SteamId, req.Amount, model.DonatePayment(req.Payment))
	if err != nil {
		return
	}

	// Payment
	var payment Payment
	switch model.DonatePayment(req.Payment) {
	case model.Alipay:
		payment = s.alipay
	case model.Wepay:
		payment = s.wepay
	default:
		return
	}

	resp.OutTradeNo = outTradeNo
	resp.QrCode, err = payment.CreateTrade(outTradeNo, req.SteamId, req.Amount)
	if err != nil {
		return
	}

	return
}

func (s *Service) FinishDonate(outTradeNo string) (err error) {
	ctx := context.Background()
	db, err := s.dao.GetTradeInfo(outTradeNo)
	if err != nil {
		return
	}

	// Do noting when donate already payed
	if db.Status == model.DonatePayed {
		return
	}

	// Give rmb
	_, err = s.money.Give(ctx, &moneyService.GiveReq{
		Uid:    db.UID,
		Money:  db.Amount * DonateRMBExchange,
		Reason: "捐助奖励",
	})
	if err != nil {
		return
	}

	// Send to MQ
	info, err := s.account.Info(ctx, &accountService.InfoReq{Uid: db.UID})
	if err == nil {
		log.Error(s.dao.CreateDonateEvent(ctx, &event.DonateEventData{
			Uid:      db.UID,
			Amount:   db.Amount,
			Username: info.Info.Username,
		}))
	}

	err = s.dao.FinishTrade(outTradeNo)
	return
}

func (s *Service) QueryDonate(ctx context.Context, req *pb.QueryDonateReq) (resp *pb.QueryDonateResp, err error) {
	resp = &pb.QueryDonateResp{}

	if len(req.OutTradeNo) == 0 {
		err = status.Error(codes.InvalidArgument, "empty trade no")
		return
	}

	res, err := s.dao.GetTradeInfo(req.OutTradeNo)
	if err != nil {
		return
	}

	resp.Uid = res.UID
	resp.Amount = res.Amount
	resp.Status = int32(res.Status)
	resp.CreateAt = res.CreatedAt
	return
}

func (s *Service) GetDonateList(ctx context.Context, req *pb.GetDonateListReq) (resp *pb.GetDonateListResp, err error) {
	resp = &pb.GetDonateListResp{}

	if req.StartTime == 0 {
		err = status.Error(codes.InvalidArgument, "invalid start time")
		return
	}
	if req.EndTime == 0 {
		req.EndTime = time.Now().Unix()
	}

	var uid int64
	if req.SteamId > 0 {
		res, err := s.account.UID(ctx, &accountService.UIDReq{SteamId: req.SteamId})
		if err != nil {
			return resp, err
		}
		uid = res.Uid
	}

	res, err := s.dao.GetDonateList(req.StartTime, req.EndTime, uid)
	if err != nil {
		return
	}

	for _, v := range res {
		resp.List = append(resp.List, &pb.QueryDonateResp{
			Uid:      v.UID,
			Amount:   v.Amount,
			Status:   int32(v.Status),
			CreateAt: v.CreatedAt,
		})
	}
	return
}

func (s *Service) AddDonate(ctx context.Context, req *pb.AddDonateReq) (resp *pb.AddDonateResp, err error) {
	resp = &pb.AddDonateResp{}

	if req.SteamId == 0 || req.Amount == 0 {
		err = status.Error(codes.InvalidArgument, "Invalid args")
		return
	}

	// Get UID
	acc, err := s.account.UID(ctx, &accountService.UIDReq{
		SteamId: req.SteamId,
	})
	if err != nil {
		return
	}

	// Get trade no
	outTradeNo, err := s.dao.CreateDonate(acc.Uid, req.SteamId, req.Amount, 0)
	if err != nil {
		return
	}

	err = s.FinishDonate(outTradeNo)
	return
}

func (s *Service) CheckTrade() {
	list, err := s.dao.GetCheckTradeList()
	if err != nil {
		log.Error(err)
		return
	}

	for _, trade := range list {
		var payment Payment
		switch trade.Payment {
		case model.Alipay:
			payment = s.alipay
		case model.Wepay:
			payment = s.wepay
		default:
			// Skip invalid payment
			continue
		}

		// Query trade
		success, err := payment.CheckTrade(trade.OutTradeNo)
		// Finish trade when it is success
		if err == nil && success {
			if err := s.FinishDonate(trade.OutTradeNo); err != nil {
				log.Error(err)
			}
			continue
		}

		// Timeout to cancel trade
		if trade.CreatedAt < time.Now().Add(-5*time.Minute).Unix() {
			if err := payment.CancelTrade(trade.OutTradeNo); err != nil {
				log.Error(err)
			}
			if err := s.dao.CancelTrade(trade.OutTradeNo); err != nil {
				log.Error(err)
			}
		}
	}
}

func (s *Service) regCron() {
	_, err := s.cron.AddFunc(DonateCheckSchedule, s.CheckTrade)
	if err != nil {
		log.Error(err)
	}
}
