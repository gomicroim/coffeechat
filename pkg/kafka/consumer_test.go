package kafka

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
	"time"
)

func consumeLoop(ctx context.Context, c sarama.Consumer, t *testing.T) {
	t.Log("ConsumeAllPartitions")
	for {
		err := ConsumeAllPartitions(ctx, c, topic, func(partition int32, partitionConsumer sarama.PartitionConsumer, message *sarama.ConsumerMessage) {
			t.Log("consumer a msg", "key:", message.Key, ",value:", string(message.Value))
		})
		if err != ErrConsumerCanceled {
			t.Log("consumer error, retry after 2s:", err)
			time.Sleep(time.Second * 2)
		} else {
			t.Log("consumer success closed")
			return
		}
	}
}

func TestConsumeAllPartitions(t *testing.T) {
	c, err := NewConsumer(kafkaAddr, nil)
	defer c.Close()
	require.NoError(t, err)

	ctx, cancel := context.WithCancel(context.Background())
	go consumeLoop(ctx, c, t)

	// cancel consume
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(200 * time.Millisecond)

	// restart consume
	ctx, cancel = context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	consumeLoop(ctx, c, t)
}

func TestConsumeAllPartitionsWithPanic(t *testing.T) {
	c, err := NewConsumer(kafkaAddr, nil)
	defer c.Close()
	require.NoError(t, err)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	defer cancel()
	err = ConsumeAllPartitions(ctx, c, topic, func(partition int32, partitionConsumer sarama.PartitionConsumer, message *sarama.ConsumerMessage) {
		t.Log("consumer a msg", "key:", message.Key, ",value:", string(message.Value))
		panicHere := strings.Split(string(message.Key), ".")[2]
		t.Log(panicHere)
	})
	require.Equal(t, err, ErrConsumerCanceled)
}
