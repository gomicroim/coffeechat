//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	klog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/google/wire"
	"wspush/internal/conf"
	"wspush/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Discover, registry.Discovery, klog.Logger, *log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(newApp, service.ProviderSet))
}
