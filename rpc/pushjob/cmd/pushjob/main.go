package main

import (
	"flag"
	"github.com/go-kratos/kratos/contrib/registry/etcd/v2"
	"github.com/go-kratos/kratos/v2"
	"os"
	"pushjob/internal/mq"

	kratoslog "github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/gomicroim/gomicroim/pkg/log"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.uber.org/zap"
	"pushjob/internal/conf"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	Name     = "push-job"
	Version  string
	flagConf string

	id, _ = os.Hostname()
)

func init() {
	flag.StringVar(&flagConf, "conf", "../../configs/config.example.yaml", "config path, eg: -conf config.yaml")
}

func newApp(logger *log.Logger, gs *grpc.Server, registry *etcd.Registry) *kratos.App {
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
	)
}

func main() {
	flag.Parse()
	kratoslog.SetLogger(log.MustNewLogger(id, Name, Version, true, 2))
	log.SetGlobalLogger(log.MustNewLogger(id, Name, Version, true, 0))

	bc := conf.MustLoad(flagConf)

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

	// create kafka producer
	producer, err := mq.NewMsgProducer(bc.Kafka.Brokers, bc.Kafka.WsPushTopic)
	if err != nil {
		panic(err)
	}

	app, cleanup, err := wireApp(bc.Server, bc.Data,
		log.MustNewLogger(id, Name, Version, true, 4), // fix kratos caller stack
		log.L, reg, producer)
	if err != nil {
		panic(err)
	}
	defer cleanup()

	// start and wait for stop signal
	if err := app.Run(); err != nil {
		panic(err)
	}
}
