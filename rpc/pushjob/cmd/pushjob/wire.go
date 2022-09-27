//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/google/wire"
	"pushjob/internal/conf"
	"pushjob/internal/mq"
	"pushjob/internal/server"
	"pushjob/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, klog.Logger, *log.Logger, *etcd.Registry, mq.PushMsgProducer) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, service.ProviderSet, newApp))
}
