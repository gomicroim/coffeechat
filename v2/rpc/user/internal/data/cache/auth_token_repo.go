package cache

import (
	"CoffeeChat/jwt"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

const (
	kAuthCachePrefix = "auth"
)

type AuthTokenRepo interface {
	CreateAuth(ctx context.Context, userid int64, details jwt.TokenDetails) error
	DeleteAuth(ctx context.Context, givenUuid string) (userId int64, err error)
	FetchAuth(ctx context.Context, tokenUuid string) (userId int64, err error)
}

type authTokenRepo struct {
	client *redis.Client
}

func NewAuthTokenRepo(client *redis.Client) AuthTokenRepo {
	return &authTokenRepo{
		client: client,
	}
}

func (a *authTokenRepo) buildTokenUuidKey(uuid string) string {
	return fmt.Sprintf("%s:token:%s", kAuthCachePrefix, uuid)
}

// CreateAuth create a new token and auto delete when token expires
func (a *authTokenRepo) CreateAuth(ctx context.Context, userid int64, details jwt.TokenDetails) error {
	at := time.Unix(details.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(details.RtExpires, 0)
	now := time.Now()

	key := a.buildTokenUuidKey(details.AccessUuid)
	value := strconv.FormatInt(userid, 10)

	// token auto expires
	if err := a.client.Set(ctx, key, value, at.Sub(now)).Err(); err != nil {
		return err
	}
	key = a.buildTokenUuidKey(details.RefreshUuid)
	return a.client.Set(ctx, key, strconv.Itoa(int(userid)), rt.Sub(now)).Err()
}

// DeleteAuth delete token immediately
func (a *authTokenRepo) DeleteAuth(ctx context.Context, givenUuid string) (userId int64, err error) {
	key := a.buildTokenUuidKey(givenUuid)
	userId, err = a.client.Del(ctx, key).Result()
	return
}

// FetchAuth if token not expires, return the userId
func (a *authTokenRepo) FetchAuth(ctx context.Context, tokenUuid string) (userId int64, err error) {
	key := a.buildTokenUuidKey(tokenUuid)
	userid, err := a.client.Get(ctx, key).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseInt(userid, 10, 64)
	return userID, nil
}
