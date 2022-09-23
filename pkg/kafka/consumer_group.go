package kafka

import (
	"context"

	"github.com/Shopify/sarama"
)

// NewConsumerGroup A nil sarama.config use the
// Default config, refer to: https://github.com/Shopify/sarama/blob/v1.32.0/config.go#L483-L499
func NewConsumerGroup(addrs []string, groupID string, config *sarama.Config) (sarama.ConsumerGroup, error) {
	if config == nil {
		config = sarama.NewConfig()
		// Aliyun kafka version 2.2.0
		config.Version = sarama.V2_0_0_0
	}

	return sarama.NewConsumerGroup(addrs, groupID, config)
}

func ConsumeGroup(ctx context.Context, cg sarama.ConsumerGroup, topics []string, handler sarama.ConsumerGroupHandler) error {
	defer cg.Close()
	for {
		err := cg.Consume(ctx, topics, handler)
		if err != nil {
			return err
		}
	}
}
