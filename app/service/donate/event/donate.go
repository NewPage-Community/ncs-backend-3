package event

import (
	pb "backend/app/service/donate/api/grpc/v1"
	"backend/pkg/database/redis"
	"backend/pkg/json"
	"backend/pkg/log"
	"context"
)

const eventKey = "ncs:donate:new"

type DonateEvent struct {
	stream *redis.Stream
}

type DonateEventData pb.AddDonateReq
type DonateCallback func(context.Context, *DonateEventData)

func NewDonateEvent(stream *redis.Stream) *DonateEvent {
	return &DonateEvent{stream}
}

func (e *DonateEvent) Create(ctx context.Context, data *DonateEventData) (err error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return
	}
	return e.stream.Publish(ctx, eventKey, string(raw))
}

func (e *DonateEvent) Listen(cb DonateCallback) error {
	return e.stream.Subscribe(eventKey, func(c context.Context, s string) {
		data := &DonateEventData{}
		err := json.Unmarshal([]byte(s), data)
		if err != nil {
			log.Error(err)
		} else {
			cb(c, data)
		}
	})
}
