package mq

import (
	"context"
	"github.com/Shopify/sarama"
	"github.com/gomicroim/gomicroim/pkg/kafka"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/gomicroim/gomicroim/pkg/rescue"
	"go.uber.org/zap"
	"sync"
	"time"
	"wspush/internal/conf"
)

type ConsumerJob interface {
	// StartConsume start consumer mq message
	StartConsume(ctx context.Context)
}

type jobImpl struct {
	config   *conf.Kafka
	consumer sarama.Consumer
	logger   *log.Logger

	partitionLatestOffset map[int32]int64
	partitionMutex        sync.RWMutex
}

// MustNewJob new kafka consumer job
func MustNewJob(config *conf.Kafka) ConsumerJob {
	consumer, err := kafka.NewConsumer(config.Brokers, nil)
	if err != nil {
		log.L.Error("NewConsumer failed", zap.Error(err))
		panic(err)
	}

	partitions, err := consumer.Partitions(config.SendMsgTopic)
	if err != nil {
		panic(err)
	}

	log.L.Info("kafka", zap.Any("brokers", config.Brokers), zap.Any("partitions", partitions))

	return &jobImpl{
		config:                config,
		consumer:              consumer,
		logger:                log.L,
		partitionLatestOffset: make(map[int32]int64),
	}
}

func (j *jobImpl) StartConsume(ctx context.Context) {
	var handle = func(partition int32, partitionConsumer sarama.PartitionConsumer, msg *sarama.ConsumerMessage) {
		j.partitionMutex.Lock()
		j.partitionLatestOffset[partition] = msg.Offset
		j.partitionMutex.Unlock()

		j.onHandlePushMsg(ctx, partitionConsumer, msg)
	}

	// create kafka consumer go routine
	go rescue.WithRecover(func() {
		for {
			err := kafka.Consume(ctx, j.consumer, j.config.SendMsgTopic, handle)
			log.L.Error("dp kafka stop to consume, after 10s retry", zap.Error(err))
			time.Sleep(10 * time.Second)
		}
	})
}
