package server

import (
	v1 "apichat/api/chat/v1"
	"apichat/internal/conf"
	"apichat/internal/service"
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gomicroim/gomicroim/v2/pkg/jwt"
	"strings"
)

func extractToken(tr transport.Transporter) (string, error) {
	// Authorization: Bearer ${token}
	token := tr.RequestHeader().Get("Authorization")
	if token == "" {
		return "", errors.New("authorization required")
	}
	if strings.Index(token, "Bearer") == -1 {
		return "", errors.New("miss Authorization token")
	}
	arr := strings.Split(token, " ")
	if len(arr) != 2 {
		return "", errors.New("invalid Bearer value")
	}
	return arr[1], nil
}

func authToken(generate *jwt.TokenGenerate) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				token, err := extractToken(tr)
				if err != nil {
					return nil, err
				}
				clientInfo, _, err := generate.ParseToken(token, false)
				if err != nil {
					return nil, err
				}
				// save jwt info to context
				ctx = jwt.WithToken(ctx, clientInfo)
			}
			return handler(ctx, req)
		}
	}
}

// NewHTTPServer new a HTTP server.
func NewHTTPServer(c *conf.Server, greeter *service.ChatService, logger log.Logger) *http.Server {
	gen := jwt.NewTokenGenerate(c.Jwt.AccessSecret, c.Jwt.RefreshSecret)
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			authToken(gen),
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
	v1.RegisterChatHTTPServer(srv, greeter)
	return srv
}
