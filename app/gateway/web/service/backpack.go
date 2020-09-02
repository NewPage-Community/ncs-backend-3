package service

import (
	itemsGW "backend/app/service/backpack/items/api/grpc"
	"backend/pkg/rpc"
)

func regBackpackService(gws *rpc.Gateways) {
	gws.AddGateway(
		itemsGW.RegisterItemsHandlerFromEndpoint,
		itemsGW.ServiceAddr,
	)
}
