package dao

import (
	"backend/app/game/server/event"
	"backend/app/game/server/model"
	"backend/pkg/ecode"
	"context"

	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

func (d *dao) Info(address string) (res *model.Info, err error) {
	res = &model.Info{}

	// DB
	err = d.db.Where(&model.Info{Address: address}).First(res).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found server (%s)", address)
	}
	return
}

func (d *dao) UpdateRcon(server *model.Info) (err error) {
	// DB
	err = d.db.Model(server).Update("rcon", server.Rcon).Error
	return
}

func (d *dao) InfoWithID(id int32) (res *model.Info, err error) {
	res = &model.Info{}

	// DB
	err = d.db.First(res, id).Error
	if err == gorm.ErrRecordNotFound {
		err = ecode.Errorf(codes.NotFound, "Can not found serverID (%d)", id)
	}
	return
}

func (d *dao) AllInfo() (res []*model.Info, err error) {
	// DB
	err = d.db.Find(&res).Error
	return
}

func (d *dao) CreateChangeMapEvent(ctx context.Context, data *event.ChangeMapEventData) error {
	return event.NewChangeMapEvent(d.stream).Create(ctx, data)
}
