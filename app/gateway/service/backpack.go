package service

import (
	itemsGW "backend/app/service/backpack/items/api/grpc"
	userGW "backend/app/service/backpack/user/api/grpc"
	"backend/pkg/rpc"
)

func regBackpackService(gws *rpc.Gateways) {
	gws.AddGateway(
		userGW.RegisterUserHandlerFromEndpoint,
		userGW.ServiceAddr,
	)
	gws.AddGateway(
		itemsGW.RegisterItemsHandlerFromEndpoint,
		itemsGW.ServiceAddr,
	)
}
