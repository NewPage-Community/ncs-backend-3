package service

import "backend/pkg/rpc"

func RegService(gws *rpc.Gateways) {
	regServerService(gws)
	regUserService(gws)
}
