package service

import (
	serverGW "backend/app/server/api/grpc"
)

func regServerService() []gateway {
	return []gateway{
		{
			handlder: serverGW.RegisterServerHandlerFromEndpoint,
			endpoint: serverGW.ServiceAddr,
		},
	}
}
