package biz

import (
	pb "chat/api"
	"chat/internal/mq"
	"context"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/gomicroim/gomicroim/protos/chat"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestMessageSyncUseCase_WriteMsg(t *testing.T) {
	bc, _ := setupBiz(t)
	producer := mq.NewMsgProducer(bc.Data.Kafka)
	consumer := mq.NewChatConsumerGroup(bc.Data.Kafka)

	biz := NewMessageSyncUseCase(bc, log.L, producer, consumer)
	time.Sleep(time.Second)

	err := biz.WriteMsg(context.Background(), &pb.IMBaseMsg{
		FromUserId:  952700,
		To:          "22",
		SessionType: chat.IMSessionType_SessionTypeSingle,
		ClientMsgId: "333.",
		MsgType:     chat.IMMsgType_MsgTypeMix,
		MsgData:     "hello world",
	})
	time.Sleep(time.Second * 5)
	require.NoError(t, err)
}

func TestMessageSyncUseCase_GetSyncMessage(t *testing.T) {
	bc, _ := setupBiz(t)
	producer := mq.NewMsgProducer(bc.Data.Kafka)
	consumer := mq.NewChatConsumerGroup(bc.Data.Kafka)

	biz := NewMessageSyncUseCase(bc, log.L, producer, consumer)
	time.Sleep(time.Second)

	reply, err := biz.GetSyncMessage(context.Background(), &pb.SyncMessageRequest{
		Member:   "22",
		LastRead: 0,
		Count:    10,
	})
	require.NoError(t, err)
	t.Log(reply)
}
