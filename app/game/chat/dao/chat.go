package dao

import (
	chatEvent "backend/app/game/chat/event"
	donateEvent "backend/app/service/donate/event"
	"context"
)

func (d *dao) CreateAllChatEvent(ctx context.Context, data *chatEvent.AllChatEventData) (err error) {
	return chatEvent.NewAllChatEvent(d.stream).Create(ctx, data)
}

func (d *dao) ListenAllChatEvent(cb chatEvent.AllChatCallback) error {
	return chatEvent.NewAllChatEvent(d.stream).Listen(cb)
}

func (d *dao) CreateDonateEvent(ctx context.Context, data *donateEvent.DonateEventData) (err error) {
	return donateEvent.NewDonateEvent(d.stream).Create(ctx, data)
}

func (d *dao) ListenDonateEvent(cb donateEvent.DonateCallback) error {
	return donateEvent.NewDonateEvent(d.stream).Listen(cb)
}
