//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/google/wire"
	"userinternal/biz"
	"userinternal/conf"
	"userinternal/data"
	"userinternal/server"
	"userinternal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Server_JWT, *conf.Data_Redis, klog.Logger, *log.Logger, *etcd.Registry) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
