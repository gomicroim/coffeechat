package service

import (
	"apiuser/internal/conf"
	"context"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	"users/api/user"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewApiUserService, NewAuthClient)

func NewAuthClient(config *conf.Discover, discovery registry.Discovery) user.AuthClient {
	rpcUserEndpoint := "discovery:///" + config.ServiceEndpoint.RpcUser
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint(rpcUserEndpoint),
		grpc.WithDiscovery(discovery),
	)
	if err != nil {
		panic(err)
	}
	return user.NewAuthClient(conn)
}
