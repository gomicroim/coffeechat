// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	log2 "github.com/gomicroim/gomicroim/v2/pkg/log"
	"wspush/internal/conf"
	"wspush/internal/service"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(server *conf.Server, discover *conf.Discover, discovery registry.Discovery, logger log.Logger, logLogger *log2.Logger) (*kratos.App, func(), error) {
	authClient := service.NewAuthClient(discover, discovery)
	wsServer := service.NewWsServer(server, authClient)
	app := newApp(logLogger, wsServer)
	return app, func() {
	}, nil
}
