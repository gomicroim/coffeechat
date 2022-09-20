package service

import (
	"apichat/internal/conf"
	"chat/api/chat"
	"context"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/gomicroim/gomicroim/v2/pkg/log"
	"github.com/google/wire"
	"go.uber.org/zap"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewChatClient, NewChatService)

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
