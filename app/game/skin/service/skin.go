package service

import (
	pb "backend/app/game/skin/api/grpc"
	itemsService "backend/app/service/backpack/items/api/grpc"
	"backend/pkg/ecode"
	"context"
	"google.golang.org/grpc/codes"
)

const (
	SkinType = 1
)

func (s *Service) GetSkins(ctx context.Context, req *pb.GetSkinsReq) (resp *pb.GetSkinsResp, err error) {
	resp = &pb.GetSkinsResp{}

	items, err := s.item.GetItems(ctx, &itemsService.GetItemsReq{
		Type: SkinType,
	})
	if err != nil {
		return
	}

	for _, v := range items.Items {
		resp.Info = append(resp.Info, &pb.SkinInfo{
			Id:       v.Id,
			Name:     v.Name,
			SkinPath: v.Attributes["skin_path"],
			ArmPath:  v.Attributes["arm_path"],
		})
	}
	return
}

func (s *Service) GetInfo(ctx context.Context, req *pb.GetInfoReq) (resp *pb.GetInfoResp, err error) {
	resp = &pb.GetInfoResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	info, err := s.dao.GetInfo(req.Uid)
	if err != nil {
		// not found
		if ecode.GetError(err).Code != codes.NotFound {
			return
		}
		info, err = s.dao.Create(req.Uid)
		if err != nil {
			return
		}
	}

	resp.Uid = info.UID
	resp.UsedSkin = info.UsedSkin
	return
}

func (s *Service) SetUsedSkin(ctx context.Context, req *pb.SetUsedSkinReq) (resp *pb.SetUsedSkinResp, err error) {
	resp = &pb.SetUsedSkinResp{}

	if req.Uid <= 0 {
		err = ecode.Errorf(codes.InvalidArgument, "Invalid UID")
		return
	}

	err = s.dao.SetUsedSkin(req.Uid, req.UsedSkin)
	return
}
