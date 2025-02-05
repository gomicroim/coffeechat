package conf

import (
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"os"
)

func MustLoad(flagconf string) *Bootstrap {
	c := config.New(
		config.WithSource(
			file.NewSource(flagconf),
		),
	)
	defer c.Close()

	if err := c.Load(); err != nil {
		panic(err)
	}

	var bc Bootstrap
	if err := c.Scan(&bc); err != nil {
		panic(err)
	}

	if bc.Server.ServerId == "auto" {
		name, err := os.Hostname()
		if err != nil {
			panic(err)
		}
		bc.Server.ServerId = name
	}
	return &bc
}
