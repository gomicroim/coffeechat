package main

import (
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/gomicroim/gomicroim/pkg/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"net/url"
	"os"

	"chat/internal/conf"
	kratoslog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     = "rpc-chat"
	Version  string
	flagConf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs/config.example.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger kratoslog.Logger, config *conf.Bootstrap, gs *grpc.Server, registry *etcd.Registry) *kratos.App {
	var endpoint = kratos.Endpoint([]*url.URL{}...)
	if config.Registry.Etcd.RegisterEndPoint != "" {
		endpoint = kratos.Endpoint(&url.URL{Host: config.Registry.Etcd.RegisterEndPoint, Scheme: "grpc"})
	}
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			gs,
		),
		kratos.Registrar(registry),
		endpoint,
	)
}

func main() {
	flag.Parse()
	kratoslog.SetLogger(log.MustNewLogger(id, Name, Version, true, 2))
	logger := log.MustNewLogger(id, Name, Version, true, 0)
	log.SetGlobalLogger(logger)

	bc := conf.MustLoad(flagConf)
	Name = bc.Registry.Etcd.RegisterServerName

	// register etcd
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: bc.Registry.Etcd.Endpoints,
	})
	if err != nil {
		panic(err)
	}
	log.L.Info("register etcd",
		zap.Strings("endpoints", bc.Registry.Etcd.Endpoints))
	reg := etcd.New(etcdClient)

	app, cleanup, err := wireApp(bc, bc.Server, bc.Data, logger, logger, reg)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
