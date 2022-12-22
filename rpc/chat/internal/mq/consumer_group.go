package mq

import (
	pb "chat/api"
	"chat/internal/conf"
	"context"
	"github.com/Shopify/sarama"
	"github.com/gomicroim/gomicroim/pkg/kafka"
	"github.com/gomicroim/gomicroim/pkg/log"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

type OutboxWriteMsgCallback func(ctx context.Context, msg *pb.SendMsgRequest)

type ChatConsumerGroup interface {
	// Consume 启动消费组
	Consume(ctx context.Context)
	// Close 关闭消费组
	Close()

	// ObserverOutboxWriteMsg 设置写消息监听器
	ObserverOutboxWriteMsg(cb OutboxWriteMsgCallback)
}

type chatConsumerGroup struct {
	logger *log.Logger
	group  sarama.ConsumerGroup
	config *conf.Data_Kafka

	outboxCb OutboxWriteMsgCallback
}

var _ ChatConsumerGroup = (*chatConsumerGroup)(nil)

func NewChatConsumerGroup(config *conf.Data_Kafka) ChatConsumerGroup {
	group, err := kafka.NewConsumerGroup(config.Brokers, config.ConsumerGroup, nil)
	if err != nil {
		panic(err)
	}
	return &chatConsumerGroup{group: group, config: config, logger: log.L}
}

func (c *chatConsumerGroup) Consume(ctx context.Context) {
	c.logger.Info("Start Consume")
	go func() {
		for {
			err := c.group.Consume(ctx, []string{c.config.ConsumerTopics.OutboxWrite}, c)
			c.logger.Info("consumer stopped, restart after 3s", zap.Error(err))
			time.Sleep(time.Second * 3)
		}
	}()
}

func (c chatConsumerGroup) Close() {
	if err := c.group.Close(); err != nil {
		log.L.Warn("close error", zap.Error(err))
	}
}

func (c *chatConsumerGroup) ObserverOutboxWriteMsg(cb OutboxWriteMsgCallback) {
	c.outboxCb = cb
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (c *chatConsumerGroup) Setup(session sarama.ConsumerGroupSession) error {
	log.L.Info("Setup")
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (c *chatConsumerGroup) Cleanup(session sarama.ConsumerGroupSession) error {
	log.L.Info("Cleanup")
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (c *chatConsumerGroup) ConsumeClaim(session sarama.ConsumerGroupSession, msg sarama.ConsumerGroupClaim) error {
	for mqMsg := range msg.Messages() {
		switch mqMsg.Topic {
		case c.config.ConsumerTopics.OutboxWrite:
			c.onHandleOutboxWriteMsg(mqMsg)
		default:
			log.L.Warn("ConsumeClaim unknown topic", zap.String("topic", mqMsg.Topic))
		}

		// 标记后自动提交offset
		session.MarkMessage(mqMsg, "")
	}

	return nil
}

func (c chatConsumerGroup) onHandleOutboxWriteMsg(message *sarama.ConsumerMessage) {
	req := pb.SendMsgRequest{}
	err := protojson.Unmarshal(message.Value, &req)
	if err != nil {
		log.L.Warn("ConsumeClaim Unmarshal error", zap.Error(err))
		return
	}

	log.L.Info("onHandleOutboxWriteMsg", zap.Any("body", string(message.Value)))

	if c.outboxCb != nil {
		c.outboxCb(context.Background(), &req)
	}
}
