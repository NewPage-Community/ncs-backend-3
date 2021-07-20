package service

import (
	itemsGW "backend/app/service/backpack/items/api/grpc/v1"
	"backend/pkg/rpc"
)

func regBackpackService(gws *rpc.Gateways) {
	gws.AddGateway(
		itemsGW.RegisterItemsHandlerFromEndpoint,
		itemsGW.ServiceAddr,
	)
}
