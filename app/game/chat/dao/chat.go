package dao

import (
	chatEvent "backend/app/game/chat/event"
	"context"
)

func (d *dao) CreateChatEvent(ctx context.Context, data *chatEvent.AllChatEventData) (err error) {
	return chatEvent.NewAllChatEvent(d.stream).Create(ctx, data)
}

func (d *dao) ListenChatEvent(cb chatEvent.AllChatCallback) error {
	return chatEvent.NewAllChatEvent(d.stream).Listen(cb)
}
