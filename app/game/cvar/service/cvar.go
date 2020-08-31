package service

import (
	pb "backend/app/game/cvar/api/grpc"
	"backend/app/game/cvar/model"
	serverService "backend/app/game/server/api/grpc"
	"backend/pkg/log"
	"context"
	"fmt"
	"github.com/opentracing/opentracing-go"
)

const (
	CronSchedule = "@every 1m"
)

func (s *Service) regCron() {
	_, err := s.cron.AddFunc(CronSchedule, s.CheckCVars)
	if err != nil {
		log.Error(err)
	}
}

func (s *Service) CheckCVars() {
	span, ctx := opentracing.StartSpanFromContext(
		context.Background(),
		"CheckCVars")
	defer span.Finish()

	res, err := s.dao.GetCVars()
	if err != nil {
		log.Error(err)
		return
	}

	for _, v := range res {
		if v.NeedUpdate {
			if err := s.NotifyCVarToServer(ctx, v); err != nil {
				log.Error(err)
			}
			if err := s.dao.UpdatedCVar(v.ID); err != nil {
				log.Error(err)
			}
		}
	}
}

func (s *Service) NotifyCVarToServer(ctx context.Context, cvar *model.CVar) (err error) {
	res, err := s.server.AllInfo(ctx, &serverService.AllInfoReq{})
	if err != nil {
		return
	}

	for _, v := range res.Info {
		if cvar.IsType(v.GameId, v.ModId, v.ServerId) {
			go func() {
				_, err := s.server.Rcon(ctx, &serverService.RconReq{
					ServerId: v.ServerId,
					Cmd:      fmt.Sprintf("%s \"%s\"", cvar.Key, cvar.Value),
				})
				if err != nil {
					log.Warn(err)
				}
			}()
		}
	}
	return
}

func (s *Service) UpdateCVars(ctx context.Context, req *pb.UpdateCVarsReq) (resp *pb.UpdateCVarsResp, err error) {
	resp = &pb.UpdateCVarsResp{}

	res, err := s.dao.GetCVars()
	if err != nil {
		return
	}

	for _, v := range res {
		if v.IsType(req.GameId, req.ModId, req.ServerId) {
			resp.Cvars = append(resp.Cvars, &pb.CVarInfo{
				Key:   v.Key,
				Value: v.Value,
			})
		}
	}
	return
}
