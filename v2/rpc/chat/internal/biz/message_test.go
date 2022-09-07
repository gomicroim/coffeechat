package biz

import (
	"CoffeeChat/log"
	"chat/internal/conf"
	"chat/internal/data"
	"chat/internal/data/cache"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
	"testing"
)

func setupBiz() (*data.Data, *redis.Client) {
	bc := conf.MustLoad("../../configs/config.yaml")

	entClient, err := data.NewEntClient(bc.Data)
	if err != nil {
		panic(err)
	}

	redisClient, err := data.NewRedis(bc.Data, log.L)
	if err != nil {
		panic(err)
	}

	dat, _, err := data.NewData(log.L, entClient, redisClient)
	return dat, redisClient
}

func TestMessageUseCase_Send(t *testing.T) {
	dat, redisClient := setupBiz()
	uc := NewMessageUseCase(data.NewMessageRepo(dat, log.L), cache.NewMsgSeq(redisClient), data.NewSessionRepo(dat, log.L))
	msg, err := uc.send(context.Background(), 1, "2", "ddddd-ffff-eeee-cccc", 1, "hello")
	assert.NoError(t, err)
	t.Log(msg)
}

func TestMessageUseCase_SendGroup(t *testing.T) {
	dat, redisClient := setupBiz()
	uc := NewMessageUseCase(data.NewMessageRepo(dat, log.L), cache.NewMsgSeq(redisClient), data.NewSessionRepo(dat, log.L))
	msg, err := uc.sendGroup(context.Background(), 1, "22", "ddddd-ffff-eeee-cccc", 1, "hello group msg")
	assert.NoError(t, err)
	t.Log(msg)
}
