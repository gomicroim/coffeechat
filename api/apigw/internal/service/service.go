package service

import (
	"apiuser/internal/conf"
	chat "chat/api"
	"context"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/google/wire"
	"go.uber.org/zap"
	"user/api/user"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewApiUserService, NewApiChatService, NewAuthClient, NewChatClient)

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

func NewChatClient(config *conf.Discover, discovery registry.Discovery) chat.ChatClient {
	rpcUserEndpoint := "discovery:///" + config.ServiceEndpoint.RpcChat
	log.L.Info("NewChatClient", zap.String("rpcChatEndpoint", rpcUserEndpoint))
	conn, err := grpc.DialInsecure(context.Background(),
		grpc.WithEndpoint(rpcUserEndpoint),
		grpc.WithDiscovery(discovery),
	)
	if err != nil {
		panic(err)
	}
	return chat.NewChatClient(conn)
}
