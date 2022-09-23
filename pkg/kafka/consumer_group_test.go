package kafka

import (
	"context"
	"log"
	"testing"

	"github.com/Shopify/sarama"
)

type MessageHandlerAutoCommit struct{}

func (h *MessageHandlerAutoCommit) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *MessageHandlerAutoCommit) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h *MessageHandlerAutoCommit) ConsumeClaim(_ sarama.ConsumerGroupSession, claims sarama.ConsumerGroupClaim) error {
	for msg := range claims.Messages() {
		log.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
	}
	return nil
}

type MessageHandlerManualCommit struct{}

func (h *MessageHandlerManualCommit) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *MessageHandlerManualCommit) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h *MessageHandlerManualCommit) ConsumeClaim(session sarama.ConsumerGroupSession, claims sarama.ConsumerGroupClaim) error {
	for msg := range claims.Messages() {
		log.Printf("Message topic:%q partition:%d offset:%d  value:%s\n", msg.Topic, msg.Partition, msg.Offset, string(msg.Value))
		session.MarkMessage(msg, "")
		session.Commit()
	}
	return nil
}

// 自动提交
func TestConsumerGroupAutoCommit(t *testing.T) {

	cg, err := NewConsumerGroup(kafkaAddr, consumerGroupId, nil)
	if err != nil {
		t.Fatal(err)
	}

	handler := &MessageHandlerAutoCommit{}
	err = ConsumeGroup(context.Background(), cg, []string{topic}, handler)
	if err != nil {
		t.Fatal(err)
	}

}

// 手动提交
func TestConsumerGroupManualCommit(t *testing.T) {
	config := sarama.NewConfig()
	config.Consumer.Offsets.AutoCommit.Enable = false

	cg, err := NewConsumerGroup(kafkaAddr, consumerGroupId, config)
	if err != nil {
		t.Fatal(err)
	}

	handler := &MessageHandlerManualCommit{}
	err = ConsumeGroup(context.Background(), cg, []string{topic}, handler)
	if err != nil {
		t.Fatal(err)
	}
}
