package event

import (
	pb "backend/app/game/server/api/grpc/v1"
	"backend/pkg/database/redis"
	"backend/pkg/json"
	"backend/pkg/log"
	"context"
)

const eventKey = "ncs:server:changemap"

type ChangeMapEvent struct {
	stream *redis.Stream
}

type ChangeMapEventData pb.ChangeMapNotifyReq
type ChangeMapCallback func(context.Context, *ChangeMapEventData)

func NewChangeMapEvent(stream *redis.Stream) *ChangeMapEvent {
	return &ChangeMapEvent{stream}
}

func (e *ChangeMapEvent) Create(ctx context.Context, data *ChangeMapEventData) (err error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return
	}
	return e.stream.Publish(ctx, eventKey, string(raw))
}

func (e *ChangeMapEvent) Listen(cb ChangeMapCallback) error {
	return e.stream.Subscribe(eventKey, func(c context.Context, s string) {
		data := &ChangeMapEventData{}
		err := json.Unmarshal([]byte(s), data)
		if err != nil {
			log.Error(err)
		} else {
			cb(c, data)
		}
	})
}
