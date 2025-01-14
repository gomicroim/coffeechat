package mq

import (
	"context"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/gomicroim/gomicroim/pkg/kafka"
	"github.com/gomicroim/gomicroim/pkg/log"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/anypb"
	pb "pushjob/api/pushjob"
)

type PushMsgProducer interface {
	BuildBatchUserReceiver(ctx context.Context, userIds []string, msgType int32, msg *anypb.Any) *pb.InternalWsPushMessage
	BuildAllUserReceiver(ctx context.Context, msgType int32, msg *anypb.Any) *pb.InternalWsPushMessage

	PushWsMsg(ctx context.Context, msg *pb.InternalWsPushMessage) (partition int32, offset int64, err error)
}

// implement MsgProducer interface check
var (
	_ PushMsgProducer = (*msgProducer)(nil)
)

type msgProducer struct {
	producer    sarama.SyncProducer
	wsPushTopic string
}

func NewMsgProducer(addr []string, topic string) (PushMsgProducer, error) {
	client, p, err := kafka.NewSyncProducerClient(addr, nil)
	if err != nil {
		return nil, err
	}

	brokers := make([]string, 0)
	for _, b := range client.Brokers() {
		brokers = append(brokers, fmt.Sprintf("id:%d, addr:%s", b.ID(), b.Addr()))
	}
	partitions, err := client.Partitions(topic)
	if err != nil {
		return nil, err
	}

	log.L.Info("start kafka producer",
		zap.Strings("brokers", brokers),
		zap.String("topic", topic),
		zap.Int32s("topic partitions", partitions))

	return &msgProducer{
		producer:    p,
		wsPushTopic: topic,
	}, nil
}

func (m *msgProducer) BuildBatchUserReceiver(ctx context.Context, userIds []string, msgType int32, msg *anypb.Any) *pb.InternalWsPushMessage {
	return m.buildReceiver(ctx, userIds, pb.InternalWsPushMessage_ReceiverTypeBatchUser, msgType, msg)
}

func (m *msgProducer) BuildAllUserReceiver(ctx context.Context, msgType int32, msg *anypb.Any) *pb.InternalWsPushMessage {
	return m.buildReceiver(ctx, nil, pb.InternalWsPushMessage_ReceiverTypeAllUser, msgType, msg)
}

func (m *msgProducer) buildReceiver(ctx context.Context, recvIds []string,
	toType pb.InternalWsPushMessage_ToType, msgType int32, msg *anypb.Any) *pb.InternalWsPushMessage {

	traceId := ""
	sc := trace.SpanContextFromContext(ctx)
	if scRaw, err := sc.MarshalJSON(); err == nil {
		traceId = string(scRaw)
	}

	return &pb.InternalWsPushMessage{
		ToType:    toType,
		ToIds:     recvIds,
		WsMsgType: msgType,
		AnyData:   msg,
		TraceId:   traceId,
	}
}

// PushWsMsg 发送消息到ws网关
func (m *msgProducer) PushWsMsg(ctx context.Context, msg *pb.InternalWsPushMessage) (partition int32, offset int64, err error) {
	// proto convert to json
	data, err := protojson.Marshal(msg)
	if err != nil {
		return 0, 0, err
	}

	mq := sarama.ProducerMessage{
		Key:   sarama.StringEncoder(""),
		Value: sarama.StringEncoder(data),
		Topic: m.wsPushTopic,
	}

	partition, offset, err = m.producer.SendMessage(&mq)

	log.Trace(ctx).Debug("SendMsgToMQ", zap.Int32("partition", partition), zap.Int64("offset", offset),
		zap.Error(err))

	return
}
