package service

import (
	pb "backend/app/service/donate/api/grpc"
	"backend/app/service/donate/model"
	accountService "backend/app/service/user/account/api/grpc"
	moneyService "backend/app/service/user/money/api/grpc"
	"backend/pkg/log"
	"context"
	"fmt"
	"github.com/smartwalle/alipay/v3"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
	"time"
)

const (
	TradeSubject      = "捐助NewPage社区"
	TradeBody         = "SteamID账号 %d 捐助NewPage社区 %d 元"
	Timeout           = "5m"
	DonateRMBExchange = 100
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
	outTradeNo, err := s.dao.CreateDonate(acc.Uid, req.SteamId, req.Amount)
	if err != nil {
		return
	}

	// alipay
	res, err := s.alipay.TradePreCreate(alipay.TradePreCreate{
		Trade: alipay.Trade{
			Subject:        TradeSubject,
			OutTradeNo:     outTradeNo,
			TotalAmount:    strconv.FormatInt(int64(req.Amount), 10),
			Body:           fmt.Sprintf(TradeBody, req.SteamId, req.Amount),
			TimeoutExpress: Timeout,
		},
	})
	if err != nil {
		return
	}
	if !res.IsSuccess() {
		log.Error(res.Content.Msg, res.Content.SubMsg)
		err = fmt.Errorf("%s - %s", res.Content.Code, res.Content.SubCode)
	}

	// resp
	resp.QrCode = res.Content.QRCode
	resp.OutTradeNo = outTradeNo
	// Check trade loop
	s.CheckTrade(outTradeNo)

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

	res, err := s.dao.GetDonateList(req.StartTime, req.EndTime)
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
	outTradeNo, err := s.dao.CreateDonate(acc.Uid, req.SteamId, req.Amount)
	if err != nil {
		return
	}

	err = s.FinishDonate(outTradeNo)
	return
}
