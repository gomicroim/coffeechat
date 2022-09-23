package mq

import (
	"context"
	"github.com/Shopify/sarama"
)

func (j *jobImpl) onHandlePushMsg(ctx context.Context, partitionConsumer sarama.PartitionConsumer, msg *sarama.ConsumerMessage) {

}
