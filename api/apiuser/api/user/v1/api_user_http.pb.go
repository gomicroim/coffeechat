// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// protoc-gen-go-http v2.3.1

package v1

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationApiUserAuth = "/user.v1.ApiUser/Auth"
const OperationApiUserRefreshToken = "/user.v1.ApiUser/RefreshToken"
const OperationApiUserRegister = "/user.v1.ApiUser/Register"

type ApiUserHTTPServer interface {
	Auth(context.Context, *AuthRequest) (*AuthReply, error)
	RefreshToken(context.Context, *RefreshTokenRequest) (*RefreshTokenReply, error)
	Register(context.Context, *RegisterRequest) (*RegisterReply, error)
}

func RegisterApiUserHTTPServer(s *http.Server, srv ApiUserHTTPServer) {
	r := s.Route("/")
	r.POST("/auth/device/register", _ApiUser_Register0_HTTP_Handler(srv))
	r.POST("/auth/login", _ApiUser_Auth0_HTTP_Handler(srv))
	r.POST("/auth/token/refresh", _ApiUser_RefreshToken0_HTTP_Handler(srv))
}

func _ApiUser_Register0_HTTP_Handler(srv ApiUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiUserRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*RegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RegisterReply)
		return ctx.Result(200, reply)
	}
}

func _ApiUser_Auth0_HTTP_Handler(srv ApiUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in AuthRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiUserAuth)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Auth(ctx, req.(*AuthRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*AuthReply)
		return ctx.Result(200, reply)
	}
}

func _ApiUser_RefreshToken0_HTTP_Handler(srv ApiUserHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RefreshTokenRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationApiUserRefreshToken)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.RefreshToken(ctx, req.(*RefreshTokenRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RefreshTokenReply)
		return ctx.Result(200, reply)
	}
}

type ApiUserHTTPClient interface {
	Auth(ctx context.Context, req *AuthRequest, opts ...http.CallOption) (rsp *AuthReply, err error)
	RefreshToken(ctx context.Context, req *RefreshTokenRequest, opts ...http.CallOption) (rsp *RefreshTokenReply, err error)
	Register(ctx context.Context, req *RegisterRequest, opts ...http.CallOption) (rsp *RegisterReply, err error)
}

type ApiUserHTTPClientImpl struct {
	cc *http.Client
}

func NewApiUserHTTPClient(client *http.Client) ApiUserHTTPClient {
	return &ApiUserHTTPClientImpl{client}
}

func (c *ApiUserHTTPClientImpl) Auth(ctx context.Context, in *AuthRequest, opts ...http.CallOption) (*AuthReply, error) {
	var out AuthReply
	pattern := "/auth/login"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApiUserAuth))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiUserHTTPClientImpl) RefreshToken(ctx context.Context, in *RefreshTokenRequest, opts ...http.CallOption) (*RefreshTokenReply, error) {
	var out RefreshTokenReply
	pattern := "/auth/token/refresh"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApiUserRefreshToken))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ApiUserHTTPClientImpl) Register(ctx context.Context, in *RegisterRequest, opts ...http.CallOption) (*RegisterReply, error) {
	var out RegisterReply
	pattern := "/auth/device/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationApiUserRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
