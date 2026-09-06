package admin

import (
	"google.golang.org/grpc"
	channelzservice "google.golang.org/grpc/channelz/service"
	internaladmin "google.golang.org/grpc/internal/admin"
)

func init() {

	internaladmin.AddService(func(registrar grpc.ServiceRegistrar) (func(), error) {
		channelzservice.RegisterChannelzServiceToServer(registrar)
		return nil, nil
	})
}

func Register(s grpc.ServiceRegistrar) (cleanup func(), _ error) {
	_ = "STUB: not implemented"
	return nil, nil
}
