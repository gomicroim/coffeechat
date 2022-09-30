package kafka

import (
	"context"
	"sync"
	"time"

	"github.com/Shopify/sarama"
)

type ConsumerHandler func(partition int32, partitionConsumer sarama.PartitionConsumer, message *sarama.ConsumerMessage)

// NewConsumer A nil sarama.config use the
// Default config, refer to: https://github.com/Shopify/sarama/blob/v1.32.0/config.go#L483-L499
// Example:
//	consumer, err := kafka.NewConsumer(kafkaAddr, nil)
//	if err != nil {
//		panic(err)
//	}
// 	partitions, err := consumer.Partitions(topic)
//	if err != nil {
//		panic(err)
//	}
//	go func() {
//      // will block
//		if err = kafka.Consume(ctx, consumer, topic, msg); err != nil {
//			t.Log(err)
//		} else {
//			t.Log("consume exit")
//		}
//	}()
// Note:
// - 使用 assign 模式，如非必要，请尽量使用ConsumerGroup方式消费。
// - 如下场景建议使用：要实现类广播效果，但是是同一个应用(服务)。比如多个gateway都要能同时消费到msg-dp的消息以解决跨网关消息转发的问题，
//   如果动态创建consumeGroup实现广播目的，成本比较大，具体参考阿里云的限制：https://help.aliyun.com/document_detail/120676.html
func NewConsumer(addrs []string, config *sarama.Config) (sarama.Consumer, error) {
	if config == nil {
		config = sarama.NewConfig()
		config.Version = sarama.V2_0_0_0
	}
	return sarama.NewConsumer(addrs, config)
}

// NewConsumerClient create consumer use default config and return client
func NewConsumerClient(addrs []string, config *sarama.Config) (sarama.Client, sarama.Consumer, error) {
	if config == nil {
		config = sarama.NewConfig()
		config.Version = sarama.V2_0_0_0
	}
	client, err := sarama.NewClient(addrs, config)
	if err != nil {
		return nil, nil, err
	}
	consumer, err := sarama.NewConsumerFromClient(client)
	return client, consumer, err
}

// Consume start consume, will block until exit, call in `goroutine`
// note: `handle` called in `goroutine`
func Consume(ctx context.Context, consumer sarama.Consumer, topic string, handle ConsumerHandler) error {
	defer consumer.Close()

	partitions, err := consumer.Partitions(topic)
	if err != nil {
		return err
	}

	waitGroup := sync.WaitGroup{}
	for k, part := range partitions {
		p, err := consumer.ConsumePartition(topic, part, sarama.OffsetNewest)
		if err != nil {
			return err
		}

		waitGroup.Add(1)
		go func(partition int32, partitionConsumer sarama.PartitionConsumer) {
			defer waitGroup.Done()
			defer partitionConsumer.AsyncClose()

			for {
				select {
				case <-ctx.Done():
					return
				case m := <-partitionConsumer.Messages():
					handle(partition, partitionConsumer, m)
				default:
					time.Sleep(time.Millisecond)
				}
			}
		}(int32(k), p)
	}
	waitGroup.Wait()
	return nil
}
