package dao

import (
	"backend/app/game/server/model"
	"backend/pkg/ecode"
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
