package data

import (
	"context"
	"github.com/gomicroim/gomicroim/pkg/log"
	"time"
	"user/internal/data/ent"
	"user/internal/data/ent/user"
	"user/internal/data/pojo"
)

type userRepo struct {
	data *Data
	log  *log.Logger
}

type User struct {
	ID       int64          `json:"id,omitempty"`
	Created  time.Time      `json:"created,omitempty"`
	Updated  time.Time      `json:"updated,omitempty"`
	NickName string         `json:"nick_name,omitempty"`
	Sex      int8           `json:"sex,omitempty"`
	Phone    string         `json:"phone,omitempty"`
	Email    string         `json:"email,omitempty"`
	Extra    pojo.UserExtra `json:"extra,omitempty"`
}

type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	Update(context.Context, *User) error
	FindByID(context.Context, int64) (*User, error)
	FindByPhone(ctx context.Context, phone string) (*User, error)
	ListAll(context.Context) ([]*User, error)
}

func NewUserRepo(data *Data, logger *log.Logger) UserRepo {
	return &userRepo{
		data: data,
		log:  logger,
	}
}

func (u *userRepo) ent2Model(user *ent.User, err error) (*User, error) {
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       user.ID,
		Created:  user.Created,
		Updated:  user.Updated,
		NickName: user.NickName,
		Sex:      user.Sex,
		Phone:    user.Phone,
		Email:    user.Email,
		Extra:    user.Extra,
	}, nil
}

func (u *userRepo) Save(ctx context.Context, user *User) (*User, error) {
	return u.ent2Model(u.data.User.Create().
		SetNickName(user.NickName).
		SetEmail(user.Email).
		SetPhone(user.Phone).
		SetSex(user.Sex).
		SetExtra(user.Extra).Save(ctx))
}
func (u *userRepo) Update(ctx context.Context, newUser *User) error {
	_, err := u.data.User.UpdateOneID(newUser.ID).
		SetNickName(newUser.NickName).
		SetSex(newUser.Sex).
		SetPhone(newUser.Phone).
		SetEmail(newUser.Email).
		Save(ctx)
	return err
}
func (u *userRepo) FindByID(ctx context.Context, id int64) (*User, error) {
	return u.ent2Model(u.data.User.Query().Where(user.ID(id)).Only(ctx))
}
func (u *userRepo) FindByPhone(ctx context.Context, phone string) (*User, error) {
	return u.ent2Model(u.data.User.Query().Where(user.Phone(phone)).Only(ctx))
}
func (u *userRepo) ListAll(ctx context.Context) ([]*User, error) {
	entUsers, err := u.data.User.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	users := make([]*User, len(entUsers))
	for k, v := range entUsers {
		item, _ := u.ent2Model(v, nil)
		users[k] = item
	}
	return users, nil
}
