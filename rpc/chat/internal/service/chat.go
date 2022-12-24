package service

import (
	pb "chat/api"
	"chat/internal/biz"
	"chat/internal/data"
	"chat/internal/mq"
	"context"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/gomicroim/gomicroim/protos/chat"
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
		ResCode: chat.IMResCode(msg.MsgResCode),
		MsgInfo: pbMsgInfo,
	}, nil
}

func (s ChatService) toPbMessage(msg *data.Message) *pb.IMBaseMsg {
	return &pb.IMBaseMsg{
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
	}
}

// SyncMessage timeline 同步消息
func (s *ChatService) SyncMessage(context.Context, *pb.SyncMessageRequest) (*pb.SyncMessageReply, error) {
	return nil, nil
}
