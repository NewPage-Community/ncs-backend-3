package service

import (
	pb "backend/app/service/backpack/items/api/grpc"
	"backend/app/service/backpack/items/model"
	"backend/pkg/log"
	"context"
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
		var attr model.Attributes
		if err = v.GetAttributes(&attr); err != nil {
			log.Error(err)
			continue
		}
		items = append(items, &pb.Item{
			Id:          v.ID,
			Name:        v.Name,
			Description: v.Description,
			Type:        v.Type,
			Price:       v.Price,
			Attributes:  attr,
		})
	}

	resp.Items = items
	return
}
