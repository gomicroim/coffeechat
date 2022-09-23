package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/gomicroim/gomicroim/v2/pkg/jwt"
	"user/api/user"
	"user/internal/biz"
	"user/internal/data"
)

var (
	kErrLoginTypeNotSupport = errors.Forbidden(user.ErrorReason_USER_LOGIN_TYPE_NOT_SUPPORT.String(), "login type not support")
)

const (
	kDomain = "im"
)

type UserService struct {
	user.UnimplementedAuthServer

	dc *biz.DeviceUseCase
	uc *biz.UserUseCase
	ac *biz.AuthUseCase
}

func NewAuthService(dc *biz.DeviceUseCase, uc *biz.UserUseCase, ac *biz.AuthUseCase) *UserService {
	return &UserService{
		dc: dc,
		uc: uc,
		ac: ac,
	}
}

func (s *UserService) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterReply, error) {
	_, err := s.dc.Register(ctx, &data.Device{
		DeviceID:   req.DeviceId,
		AppVersion: req.AppVersion,
		OsVersion:  req.OsVersion,
	})
	if err != nil {
		return nil, err
	}
	token, err := s.ac.CreateToken(ctx, jwt.ClientInfo{
		DeviceId: req.DeviceId,
	}, false)
	if err != nil {
		return nil, err
	}
	return &user.RegisterReply{
		AccessToken: token.AccessToken,
		AtExpires:   token.AtExpires,
	}, nil
}

func (s *UserService) Auth(ctx context.Context, req *user.AuthRequest) (*user.AuthReply, error) {
	if req.LoginType != user.AuthRequest_loginTypeMobile {
		return nil, kErrLoginTypeNotSupport
	}

	info, err := s.ac.GetClientInfo(req.AccessToken)
	if err != nil {
		return nil, err
	}

	// auto register
	result, err := s.uc.RegisterByPhone(ctx, req.ByMobile.Phone)
	if err != nil {
		return nil, err
	}

	tokenDetail, err := s.ac.CreateToken(ctx, jwt.ClientInfo{
		UserId:     result.ID,
		DeviceId:   info.DeviceId,
		ClientType: req.ClientType.String(),
		Domain:     kDomain,
	}, true)
	if err != nil {
		return nil, err
	}

	return &user.AuthReply{
		AccessToken:  tokenDetail.AccessToken,
		RefreshToken: tokenDetail.RefreshToken,
		AtExpires:    tokenDetail.AtExpires,
		RtExpires:    tokenDetail.RtExpires,
		UserId:       result.ID,
	}, nil
}

func (s *UserService) TokenValid(ctx context.Context, req *user.TokenValidRequest) (*user.TokenValidReply, error) {
	clientInfo, details, err := s.ac.VerifyToken(ctx, req.Token, req.IsRefreshToken)
	if err != nil {
		return nil, err
	}
	return &user.TokenValidReply{
		Expires: details.AtExpires,
		ClientInfo: &user.TokenValidReply_ClientInfo{
			UserId:     clientInfo.UserId,
			DeviceId:   clientInfo.DeviceId,
			ClientType: clientInfo.ClientType,
			Domain:     clientInfo.Domain,
			AppVersion: clientInfo.AppVersion,
		},
	}, nil
}