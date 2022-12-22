package mq

import (
	pb "chat/api"
	"chat/internal/conf"
	"context"
	"encoding/json"
	"github.com/Shopify/sarama"
	"github.com/gomicroim/gomicroim/pkg/kafka"
	"google.golang.org/protobuf/encoding/protojson"
)

const (
	BizTypeChatMsg = "chat"
)

type DeadChatMsg struct {
	BizType string
	Content string
}

type MsgProducer interface {
	// SendOutboxMsg 发送消息
	SendOutboxMsg(ctx context.Context, key string, msg *pb.SendMsgRequest) (partition int32, offset int64, err error)
	// SendDeadMsg 发送到死信队列
	SendDeadMsg(ctx context.Context, key string, deadMsg DeadChatMsg) (partition int32, offset int64, err error)
}

type msgProducer struct {
	producer sarama.SyncProducer
	config   *conf.Data_Kafka
}

func NewMsgProducer(data *conf.Data_Kafka) MsgProducer {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewHashPartitioner // 确保同一个group的消息有序

	producer, err := kafka.NewSyncProducer(data.Brokers, config)
	if err != nil {
		panic(err)
	}

	return &msgProducer{producer: producer, config: data}
}

func (m msgProducer) SendOutboxMsg(ctx context.Context, key string, msg *pb.SendMsgRequest) (partition int32, offset int64, err error) {
	buffer, err := protojson.Marshal(msg)
	if err != nil {
		return 0, 0, err
	}

	return m.producer.SendMessage(&sarama.ProducerMessage{
		Topic: m.config.ProducerTopics.OutboxWrite,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(buffer),
	})
}

func (m msgProducer) SendDeadMsg(ctx context.Context, key string, deadMsg DeadChatMsg) (partition int32, offset int64, err error) {
	buffer, err := json.Marshal(deadMsg)
	if err != nil {
		return 0, 0, err
	}
	return m.producer.SendMessage(&sarama.ProducerMessage{
		Topic: m.config.ProducerTopics.Dead,
		Key:   sarama.StringEncoder(key),
		Value: sarama.StringEncoder(buffer),
	})
}
