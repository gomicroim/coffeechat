package biz

import (
	"CoffeeChat/log"
	"chat/internal/conf"
	"chat/internal/data"
	"chat/internal/data/cache"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func setupMessageUseCase() *MessageUseCase {
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
	return NewMessageUseCase(data.NewMessageRepo(dat, log.L), cache.NewMsgSeq(redisClient), data.NewSessionRepo(dat, log.L))
}

func TestMessageUseCase_Send(t *testing.T) {
	uc := setupMessageUseCase()
	msg, err := uc.send(context.Background(), 1, "2", "ddddd-ffff-eeee-cccc", 1, "hello",
		time.Now().Unix())
	assert.NoError(t, err)
	t.Log(msg)
}

func TestMessageUseCase_SendGroup(t *testing.T) {
	uc := setupMessageUseCase()
	msg, err := uc.sendGroup(context.Background(), 1, "22", "ddddd-ffff-eeee-cccc", 1, "hello group msg",
		time.Now().Unix())
	assert.NoError(t, err)
	t.Log(msg)
}
