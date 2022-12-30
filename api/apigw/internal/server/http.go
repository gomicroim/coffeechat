package server

import (
	"apiuser/internal/conf"
	"apiuser/internal/service"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/swagger-api/openapiv2"
	v1 "github.com/gomicroim/gomicroim/protos/api"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, user *service.ApiUserService, chat *service.ApiChatService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterApiUserHTTPServer(srv, user)
	v1.RegisterChatHTTPServer(srv, chat)

	// swagger-ui
	openApiHandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openApiHandler)
	return srv
}
