package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/transport/http"
	"strings"
	"user/api/user"

	pb "github.com/gomicroim/gomicroim/protos/api"
)

type ApiUserService struct {
	pb.UnimplementedApiUserServer

	client user.AuthClient
}

func NewApiUserService(client user.AuthClient) *ApiUserService {
	return &ApiUserService{client: client}
}

func (s *ApiUserService) DeviceRegister(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterReply, error) {
	result, err := s.client.Register(ctx, &user.RegisterRequest{
		DeviceId:   req.DeviceId,
		AppVersion: req.AppVersion,
		OsVersion:  req.OsVersion,
	})
	if err != nil {
		return nil, err
	}
	return &pb.RegisterReply{
		AccessToken: result.AccessToken,
		AtExpires:   result.AtExpires,
	}, nil
}

func (s *ApiUserService) Auth(ctx context.Context, req *pb.AuthRequest) (*pb.AuthReply, error) {
	httpCtx := ctx.(http.Context)
	token, err := extractToken(httpCtx)
	if err != nil {
		return nil, err
	}

	if req.LoginType != pb.AuthRequest_loginTypeMobile {
		return nil, errors.BadRequest(pb.ApiErrorReason_AUTH_NOT_SUPPORT_LOGIN_TYPE.String(), "not support login type")
	}
	if req.ByMobile == nil || req.ByMobile.Phone == "" || req.ByMobile.Code == "" {
		return nil, errors.BadRequest(pb.ApiErrorReason_AUTH_LOGIN_MISS_PHONE_OR_CODE.String(), "miss phone or code")
	}

	result, err := s.client.Auth(ctx, &user.AuthRequest{
		LoginType: user.AuthRequest_LoginType(req.LoginType),
		ByMobile: &user.AuthRequest_MobileAuth{
			Phone: req.ByMobile.Phone,
			Code:  req.ByMobile.Code,
		},
		AccessToken: token,
		ClientType:  user.AuthRequest_ClientType(req.ClientType),
	})
	if err != nil {
		return nil, err
	}

	return &pb.AuthReply{
		Token: &pb.TokenInfo{
			AccessToken:  result.TokenInfo.AccessToken,
			RefreshToken: result.TokenInfo.RefreshToken,
			AtExpires:    result.TokenInfo.AtExpires,
			RtExpires:    result.TokenInfo.RtExpires,
		},
		UserId: result.UserId,
	}, nil
}

func (s *ApiUserService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.RefreshTokenReply, error) {
	httpCtx := ctx.(http.Context)
	refreshToken, err := extractToken(httpCtx)
	if err != nil {
		return nil, err
	}
	result, err := s.client.RefreshToken(ctx, &user.RefreshTokenRequest{RefreshToken: refreshToken})
	if err != nil {
		return nil, err
	}
	return &pb.RefreshTokenReply{Token: &pb.TokenInfo{
		AccessToken:  result.TokenInfo.AccessToken,
		RefreshToken: result.TokenInfo.RefreshToken,
		AtExpires:    result.TokenInfo.AtExpires,
		RtExpires:    result.TokenInfo.RtExpires,
	}}, nil
}

func extractToken(ctx http.Context) (string, error) {
	token := ctx.Header().Get("Authorization")
	if token == "" {
		return "", errors.BadRequest(pb.ApiErrorReason_TOKEN_MISS.String(), "miss token")
	}

	arr := strings.Split(token, " ")
	if len(arr) != 2 {
		return "", errors.BadRequest(pb.ApiErrorReason_TOKEN_INVALID.String(), "miss token")
	} else {
		return arr[1], nil
	}
}
