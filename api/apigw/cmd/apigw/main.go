package main

import (
	"apiuser/internal/conf"
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	kratoslog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gomicroim/gomicroim/pkg/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"os"
)

var (
	Name             = "apigw"
	Version          string
	flagConf         string
	EtcdApiNameSpace = "/api"

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs/config.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger *log.Logger, hs *http.Server) *kratos.App {
	return kratos.New(
		kratos.ID(id),
		kratos.Name(Name),
		kratos.Version(Version),
		kratos.Metadata(map[string]string{}),
		kratos.Logger(logger),
		kratos.Server(
			hs,
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

	// wire depends
	app, cleanup, err := wireApp(bc.Server, bc.Discover,
		log.MustNewLogger(id, Name, Version, true, 4),
		log.L, dis)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
