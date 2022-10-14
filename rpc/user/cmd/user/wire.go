//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/google/wire"
	"user/internal/biz"
	"user/internal/conf"
	"user/internal/data"
	"user/internal/server"
	"user/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Bootstrap, *conf.Data, *conf.Server_JWT, *conf.Data_Redis, klog.Logger, *log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
