package biz

import (
	"context"
	"github.com/gomicroim/gomicroim/pkg/jwt"
	"user/internal/conf"
	"user/internal/data/cache"
)

type AuthUseCase struct {
	authRepo cache.AuthTokenRepo
	jwt      *jwt.TokenGenerate
}

func NewAuthUseCase(jwtConf *conf.Server_JWT, repo cache.AuthTokenRepo) *AuthUseCase {
	return &AuthUseCase{
		jwt:      jwt.NewTokenGenerate(jwtConf.AccessSecret, jwtConf.RefreshSecret),
		authRepo: repo,
	}
}

func (a *AuthUseCase) CreateToken(ctx context.Context, info jwt.ClientInfo, needPersistence bool) (*jwt.TokenDetails, error) {
	details, err := a.jwt.CreateToken(info)
	if err == nil && needPersistence {
		// persistence to redis
		if err = a.authRepo.CreateAuth(ctx, info.UserId, *details); err != nil {
			return nil, err
		}
	}
	return details, err
}

func (a *AuthUseCase) VerifyToken(ctx context.Context, token string, isRefreshToken bool) (*jwt.ClientInfo, *jwt.TokenDetails, error) {
	clientInfo, details, err := a.jwt.ParseToken(token, isRefreshToken)
	if err != nil {
		return nil, nil, err
	}
	// check token exist in redis
	if _, err = a.authRepo.FetchAuth(ctx, details.AccessUuid); err != nil {
		return nil, nil, err
	}
	return clientInfo, details, nil
}

func (a *AuthUseCase) GetClientInfo(token string) (*jwt.ClientInfo, error) {
	clientInfo, _, err := a.jwt.ParseToken(token, false)
	if err != nil {
		return nil, err
	}
	return clientInfo, nil
}

func (a *AuthUseCase) RefreshToken(ctx context.Context, refreshToken string) error {
	clientInfo, details, err := a.jwt.ParseToken(refreshToken, true)
	if err != nil {
		return err
	}
	//Delete the previous Refresh Token
	if _, err = a.authRepo.DeleteAuth(ctx, details.RefreshUuid); err != nil {
		return err
	}
	token, err := a.jwt.CreateToken(*clientInfo)
	if err != nil {
		return err
	}
	return a.authRepo.CreateAuth(ctx, clientInfo.UserId, *token)
}

func (a *AuthUseCase) DeleteToken(ctx context.Context, accessToken string) error {
	_, details, err := a.jwt.ParseToken(accessToken, false)
	if err != nil {
		return err
	}
	_, err = a.authRepo.DeleteAuth(ctx, details.AccessUuid)
	return err
}
