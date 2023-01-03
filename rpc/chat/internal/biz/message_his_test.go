package biz

import (
	"chat/internal/conf"
	"chat/internal/data"
	"chat/internal/data/cache"
	"context"
	pb "github.com/gomicroim/gomicroim/protos/wspush"
	"github.com/stretchr/testify/require"
	"math"
	"testing"

	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/stretchr/testify/assert"
)

func setupBiz(t *testing.T) (*conf.Bootstrap, *data.Data) {
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
	require.NoError(t, err)
	return bc, dat
}

func TestMessageUseCase_SendSingle(t *testing.T) {
	_, dat := setupBiz(t)
	uc := NewMessageHistoryUseCase(data.NewMessageRepo(dat, log.L), cache.NewMsgSeq(dat.RedisClient), data.NewSessionRepo(dat, log.L))
	msg, err := uc.Send(context.Background(), 1, "2", pb.IMSessionType_SessionTypeSingle, "ddddd-ffff-eeee-cccc", 1, "hello")
	assert.NoError(t, err)
	t.Log(msg)
}

func TestMessageUseCase_SendGroup(t *testing.T) {
	_, dat := setupBiz(t)
	uc := NewMessageHistoryUseCase(data.NewMessageRepo(dat, log.L), cache.NewMsgSeq(dat.RedisClient), data.NewSessionRepo(dat, log.L))
	msg, err := uc.Send(context.Background(), 1, "22", pb.IMSessionType_SessionTypeNormalGroup, "ddddd-ffff-eeee-cccc", 1, "hello group msg")
	assert.NoError(t, err)
	t.Log(msg)
}

func TestMessageUseCase_GetMessageList(t *testing.T) {
	_, dat := setupBiz(t)
	uc := NewMessageHistoryUseCase(data.NewMessageRepo(dat, log.L), cache.NewMsgSeq(dat.RedisClient), data.NewSessionRepo(dat, log.L))
	msg, err := uc.GetMessageList(context.Background(), 1, "22",
		pb.IMSessionType_SessionTypeNormalGroup, true,
		int64(math.MaxInt64), 10)
	assert.NoError(t, err)
	t.Log(msg)
}
