package dao

import (
	chatEvent "backend/app/game/chat/event"
	serverEvent "backend/app/game/server/event"
	donateEvent "backend/app/service/donate/event"
	"context"
)

func (d *dao) CreateAllChatEvent(ctx context.Context, data *chatEvent.AllChatEventData) (err error) {
	return chatEvent.NewAllChatEvent(d.stream).Create(ctx, data)
}

func (d *dao) ListenAllChatEvent(cb chatEvent.AllChatCallback) error {
	return chatEvent.NewAllChatEvent(d.stream).Listen(cb)
}

func (d *dao) CreateChangeMapEvent(ctx context.Context, data *serverEvent.ChangeMapEventData) (err error) {
	return serverEvent.NewChangeMapEvent(d.stream).Create(ctx, data)
}

func (d *dao) ListenChangeMapEvent(cb serverEvent.ChangeMapCallback) error {
	return serverEvent.NewChangeMapEvent(d.stream).Listen(cb)
}

func (d *dao) CreateDonateEvent(ctx context.Context, data *donateEvent.DonateEventData) (err error) {
	return donateEvent.NewDonateEvent(d.stream).Create(ctx, data)
}

func (d *dao) ListenDonateEvent(cb donateEvent.DonateCallback) error {
	return donateEvent.NewDonateEvent(d.stream).Listen(cb)
}
