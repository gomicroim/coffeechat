package service

import (
	pb "chat/api"
	"chat/internal/biz"
	"chat/internal/conf"
	"chat/internal/mq"
	"context"
	v1 "github.com/gomicroim/go-timeline/api/v1"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/gomicroim/gomicroim/protos/chat"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"time"
)

type ChatService struct {
	pb.UnimplementedChatServer

	persistentBiz *biz.MessageHistoryUseCase // 消息持久化到mysql，读扩散模型
	timeline      v1.TimelineClient          // 写扩散存储到mongo，目前是grpc方式，可以考虑优化为mq异步提升吞吐
	logger        *log.Logger

	sessionBiz *biz.RecentSessionUseCase
	consumer   mq.ChatConsumerGroup
	producer   mq.MsgProducer
}

func NewChatService(msgBiz *biz.MessageHistoryUseCase, sessionBiz *biz.RecentSessionUseCase,
	l *log.Logger, config *conf.Bootstrap, group mq.ChatConsumerGroup, producer mq.MsgProducer) *ChatService {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*1)
	defer cancel()
	conn, err := grpc.DialContext(ctx, config.Data.Timeline.GrpcAddr)
	if err != nil {
		panic(err)
	}
	chatSrv := &ChatService{
		persistentBiz: msgBiz,
		timeline:      v1.NewTimelineClient(conn),
		sessionBiz:    sessionBiz,
		logger:        l,
		consumer:      group,
		producer:      producer,
	}
	group.ObserverOutboxWriteMsg(chatSrv.onHandleOutboxWriteMsg)
	return chatSrv
}

func (s *ChatService) SendMsg(ctx context.Context, req *pb.SendMsgRequest) (*pb.SendMsgReply, error) {
	// persistent to mysql (inbox, read diffusion)
	msg, err := s.persistentBiz.Send(ctx, req.FromUserId, req.To, req.SessionType, req.ClientMsgId, req.MsgType, req.MsgData)
	if err != nil {
		return nil, err
	}

	// persistent to mongo use kafka (outbox, write diffusion)
	_, _, err = s.producer.SendOutboxMsg(ctx, req.To, req)
	if err != nil {
		log.L.Error("SendOutboxMsg err", zap.Error(err))
	}

	return &pb.SendMsgReply{
		MsgSeq:  msg.ServerMsgSeq,
		ResCode: chat.IMResCode(msg.MsgResCode),
		MsgInfo: &pb.IMBaseMsg{
			FromUserId:   msg.From,
			To:           msg.To,
			SessionType:  chat.IMSessionType(msg.SessionType),
			ClientMsgId:  msg.ClientMsgID,
			ServerMsgSeq: msg.ServerMsgSeq,
			MsgType:      chat.IMMsgType(msg.MsgType),
			MsgData:      msg.MsgData,
			MsgResCode:   chat.IMResCode(msg.MsgResCode),
			MsgFeature:   chat.IMMsgFeature(msg.MsgFeature),
			MsgStatus:    chat.IMMsgStatus(msg.MsgStatus),
			CreateTime:   msg.Created.Unix(),
		},
	}, nil
}

// SyncMessage timeline 同步消息
func (s *ChatService) SyncMessage(context.Context, *pb.SyncMessageRequest) (*pb.SyncMessageReply, error) {
	return nil, nil
}

func (s *ChatService) onHandleOutboxWriteMsg(ctx context.Context, req *pb.SendMsgRequest) {
	s.logger.Info("onHandleOutboxWriteMsg", zap.String("req", req.String()))

	var (
		err            error
		sendReply      *v1.SendMessageReply
		sendGroupReply *v1.SendGroupReply
	)

	switch req.SessionType {
	case chat.IMSessionType_SessionTypeSingle:
		sendReply, err = s.timeline.Send(context.Background(), &v1.SendMessageRequest{
			From:    "",
			To:      "",
			Message: "",
		})
		if sendReply != nil {
			s.logger.Info("timeline.Send reply", zap.Any("reply", sendReply.String()))
		}
	case chat.IMSessionType_SessionTypeNormalGroup:
		sendGroupReply, err = s.timeline.SendGroup(context.Background(), &v1.SendGroupRequest{})
		if sendGroupReply != nil {
			s.logger.Info("timeline.SendGroup reply", zap.Any("reply", sendReply.String()))
		}
	default:
		s.logger.Warn("error write to mongodb: unknown session type")
		return
	}

	// 写入到死信队列，手动处理
	if err != nil {
		buffer, err := protojson.Marshal(req)
		if err != nil {
			s.logger.Error("SendDeadMsg Marshal error", zap.Error(err))
			return
		}
		_, _, err = s.producer.SendDeadMsg(ctx, req.To, mq.DeadChatMsg{
			BizType: mq.BizTypeChatMsg,
			Content: string(buffer),
		})
		if err != nil {
			s.logger.Error("SendDeadMsg error", zap.Error(err))
		}
	}
}
