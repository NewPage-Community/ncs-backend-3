package event

import (
	pb "backend/app/game/chat/api/grpc/v1"
	"backend/pkg/database/redis"
	"backend/pkg/json"
	"backend/pkg/log"
	"context"
)

const eventKey = "ncs:chat:allchat"

type AllChatEvent struct {
	stream *redis.Stream
}

type AllChatEventData pb.AllChatReq
type AllChatCallback func(context.Context, *AllChatEventData)

func NewAllChatEvent(stream *redis.Stream) *AllChatEvent {
	return &AllChatEvent{stream}
}

func (e *AllChatEvent) Create(ctx context.Context, data *AllChatEventData) (err error) {
	raw, err := json.Marshal(data)
	if err != nil {
		return
	}
	return e.stream.Publish(ctx, eventKey, string(raw))
}

func (e *AllChatEvent) Listen(cb AllChatCallback) error {
	return e.stream.Subscribe(eventKey, func(c context.Context, s string) {
		data := &AllChatEventData{}
		err := json.Unmarshal([]byte(s), data)
		if err != nil {
			log.Error(err)
		} else {
			cb(c, data)
		}
	})
}
