package service

import (
	pb "chat/api"
	"chat/internal/biz"
	"chat/internal/data"
	"chat/internal/mq"
	"context"
	"github.com/gomicroim/gomicroim/pkg/log"
	"go.uber.org/zap"
)

type ChatService struct {
	pb.UnimplementedChatServer

	msgHisBiz    biz.MessageHistoryUseCase // 消息持久化到mysql，读扩散模型
	msgWriterBiz biz.MessageSyncUseCase    // 写入到用户的消息收件箱

	logger *log.Logger

	sessionBiz *biz.RecentSessionUseCase
	consumer   mq.ChatConsumerGroup
	producer   mq.MsgProducer
}

func NewChatService(history biz.MessageHistoryUseCase, sync biz.MessageSyncUseCase, sessionBiz *biz.RecentSessionUseCase,
	l *log.Logger, group mq.ChatConsumerGroup, producer mq.MsgProducer) *ChatService {

	chatSrv := &ChatService{
		msgHisBiz:    history,
		msgWriterBiz: sync,
		sessionBiz:   sessionBiz,
		logger:       l,
		consumer:     group,
		producer:     producer,
	}
	return chatSrv
}

func (s *ChatService) SendMsg(ctx context.Context, req *pb.SendMsgRequest) (*pb.SendMsgReply, error) {
	// persistent to mysql (inbox, read diffusion)
	msg, err := s.msgHisBiz.Send(ctx, req.FromUserId, req.To, req.SessionType, req.ClientMsgId, req.MsgType, req.MsgData)
	if err != nil {
		return nil, err
	}

	pbMsgInfo := s.toPbMessage(msg)

	// persistent to mongo use kafka (outbox, write diffusion)
	go func() {
		err = s.msgWriterBiz.WriteMsg(ctx, pbMsgInfo)
		if err != nil {
			log.L.Error("WriteMsg err", zap.Error(err))
		}
	}()

	return &pb.SendMsgReply{
		MsgSeq:  msg.ServerMsgSeq,
		ResCode: wspush.IMResCode(msg.MsgResCode),
		MsgInfo: pbMsgInfo,
	}, nil
}

func (s ChatService) toPbMessage(msg *data.Message) *pb.IMBaseMsg {
	return &pb.IMBaseMsg{
		FromUserId:   msg.From,
		To:           msg.To,
		SessionType:  wspush.IMSessionType(msg.SessionType),
		ClientMsgId:  msg.ClientMsgID,
		ServerMsgSeq: msg.ServerMsgSeq,
		MsgType:      wspush.IMMsgType(msg.MsgType),
		MsgData:      msg.MsgData,
		MsgResCode:   wspush.IMResCode(msg.MsgResCode),
		MsgFeature:   wspush.IMMsgFeature(msg.MsgFeature),
		MsgStatus:    wspush.IMMsgStatus(msg.MsgStatus),
		CreateTime:   msg.Created.Unix(),
	}
}

// SyncMessage timeline 同步消息
func (s *ChatService) SyncMessage(context.Context, *pb.SyncMessageRequest) (*pb.SyncMessageReply, error) {
	return nil, nil
}
