package biz

import (
	"context"
	"github.com/gomicroim/gomicroim/v2/pkg/log"
	"user/internal/data"
	"user/internal/data/ent"
)

type UserUseCase struct {
	repo data.UserRepo
	log  *log.Logger
}

func NewUserUseCase(repo data.UserRepo, logger *log.Logger) *UserUseCase {
	return &UserUseCase{
		repo: repo,
		log:  logger,
	}
}

func (u *UserUseCase) RegisterByPhone(ctx context.Context, phone string) (*data.User, error) {
	user, err := u.repo.FindByPhone(ctx, phone)
	if err != nil {
		if ent.IsNotFound(err) {
			user = &data.User{
				NickName: phone,
				Phone:    phone,
			}
			return u.repo.Save(ctx, user)
		}
		return nil, err
	}
	return user, nil
}
