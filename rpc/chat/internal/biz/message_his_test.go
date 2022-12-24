package biz

import (
	"chat/internal/conf"
	"chat/internal/data"
	"chat/internal/data/cache"
	"context"
	"github.com/go-redis/redis/v8"
	pb "github.com/gomicroim/gomicroim/protos/chat"
	"math"
	"testing"

	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/stretchr/testify/assert"
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
	uc := NewMessageHistoryUseCase(data.NewMessageRepo(dat, log.L), cache.NewMsgSeq(redisClient), data.NewSessionRepo(dat, log.L))
	msg, err := uc.Send(context.Background(), 1, "2", pb.IMSessionType_SessionTypeNormalGroup, "ddddd-ffff-eeee-cccc", 1, "hello")
	assert.NoError(t, err)
	t.Log(msg)
}

func TestMessageUseCase_SendGroup(t *testing.T) {
	dat, redisClient := setupBiz()
	uc := NewMessageHistoryUseCase(data.NewMessageRepo(dat, log.L), cache.NewMsgSeq(redisClient), data.NewSessionRepo(dat, log.L))
	msg, err := uc.sendGroup(context.Background(), 1, "22", "ddddd-ffff-eeee-cccc", 1, "hello group msg")
	assert.NoError(t, err)
	t.Log(msg)
}

func TestMessageUseCase_GetMessageList(t *testing.T) {
	dat, redisClient := setupBiz()
	uc := NewMessageHistoryUseCase(data.NewMessageRepo(dat, log.L), cache.NewMsgSeq(redisClient), data.NewSessionRepo(dat, log.L))
	msg, err := uc.GetMessageList(context.Background(), 1, "22",
		pb.IMSessionType_SessionTypeNormalGroup, true,
		int64(math.MaxInt64), 10)
	assert.NoError(t, err)
	t.Log(msg)
}
