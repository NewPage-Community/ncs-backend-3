package service

import "backend/pkg/rpc"

func RegService(gws *rpc.Gateways) {
	regGameService(gws)
	regUserService(gws)
}
