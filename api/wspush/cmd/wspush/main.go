package main

import (
	"context"
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	kratoslog "github.com/go-kratos/kratos/v2/log"
	"github.com/gomicroim/gomicroim/pkg/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"os"
	"wspush/internal/conf"
	"wspush/internal/mq"
	"wspush/internal/service"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     = "ws-push"
	Version  string
	flagConf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs/config.example.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger *log.Logger, httpSrv *service.WsServer) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Logger(logger),
		kratos.Server(
			httpSrv,
		),
	)
}

func main() {
	flag.Parse()
	kratoslog.SetLogger(log.MustNewLogger(id, Name, Version, true, 2))
	log.SetGlobalLogger(log.MustNewLogger(id, Name, Version, true, 0))

	bc := conf.MustLoad(flagConf)

	// etcd conn
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints: bc.Discover.Etcd.Endpoints,
	})
	if err != nil {
		panic(err)
	}
	dis := etcd.New(etcdClient)
	log.L.Info("connect etcd", zap.Strings("etcd", bc.Discover.Etcd.Endpoints))

	// start kafka consumer (manual manager partition)
	consumerJob := mq.MustNewJob(bc.Kafka)
	consumerJob.StartConsume(context.Background())

	app, cleanup, err := wireApp(bc.Server, bc.Discover, dis,
		log.MustNewLogger(id, Name, Version, true, 4), // fix kratos caller stack
		log.L)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
