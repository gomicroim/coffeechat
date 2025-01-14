//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"chat/internal/biz"
	"chat/internal/conf"
	"chat/internal/data"
	"chat/internal/server"
	"chat/internal/service"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	kratosLogs "github.com/go-kratos/kratos/v2/log"
	"github.com/gomicroim/gomicroim/pkg/log"

	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Bootstrap, *conf.Server, *conf.Data, kratosLogs.Logger, *log.Logger, *etcd.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
