package service

import (
	pb "backend/app/service/backpack/items/api/grpc/v1"
	"backend/app/service/backpack/items/model"
	"backend/pkg/log"
	"context"

	"google.golang.org/protobuf/types/known/structpb"
)

func (s *Service) GetItems(ctx context.Context, req *pb.GetItemsReq) (resp *pb.GetItemsResp, err error) {
	resp = &pb.GetItemsResp{}

	var items []*pb.Item
	res, err := s.dao.GetItems()
	if err != nil {
		return
	}

	for _, v := range res {
		if req.Type > 0 && v.Type != req.Type {
			continue
		}
		var attr *model.Attributes
		if attr, err = v.GetAttributes(); err != nil {
			log.Error(err)
			continue
		}
		var any *structpb.Struct
		if any, err = structpb.NewStruct(*attr); err != nil {
			log.Error(err)
			continue
		}
		items = append(items, &pb.Item{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Type:        v.Type,
			Price:       v.Price,
			Attributes:  any,
			Discount:    v.Discount,
			Available:   v.Available,
		})
	}

	resp.Items = items
	return
}
