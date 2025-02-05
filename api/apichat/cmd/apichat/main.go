package main

import (
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"github.com/gomicroim/gomicroim/pkg/log"
	"go.uber.org/zap"
	"os"

	"apichat/internal/conf"
	kratoslog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/http"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     = "api-chat"
	Version  string
	flagConf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs/config.example.yaml", "config path, eg: -conf config.yaml")
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
	log.L.Info("connect etcd:", zap.Strings("endpoints", bc.Discover.Etcd.Endpoints))

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
