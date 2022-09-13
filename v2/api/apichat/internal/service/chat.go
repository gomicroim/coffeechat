package service

import (
	"chat/api/chat"
	"context"
	"github.com/gomicroim/gomicroim/v2/pkg/jwt"
	"github.com/gomicroim/gomicroim/v2/pkg/log"
	"go.uber.org/zap"
	"strings"

	v1 "apichat/api/chat/v1"
	pb "chat/api/chat"
)

type ChatService struct {
	v1.UnimplementedChatServer

	chatClient chat.ChatClient
	log        *log.Logger
}

func NewChatService(client chat.ChatClient) *ChatService {
	return &ChatService{chatClient: client, log: log.L}
}

func (s *ChatService) SendMsg(ctx context.Context, req *v1.SendMsgRequest) (*v1.SendMsgReply, error) {
	sessionType := s.parseSession(req.To)
	userId := jwt.GetUserId(ctx)

	r, err := s.chatClient.SendMsg(ctx, &pb.SendMsgRequest{
		FromUserId:  userId,
		To:          req.To,
		SessionType: sessionType,
		ClientMsgId: req.ClientMsgId,
		MsgType:     req.MsgType,
		MsgData:     req.MsgData,
	})
	if err != nil {
		return nil, err
	}
	s.log.Info("SendMsg", zap.Int64("from", userId), zap.String("to", req.To),
		zap.Any("msgType", req.MsgType), zap.String("msgData", req.MsgData),
		zap.Int64("msgSeq", r.MsgSeq))

	return &v1.SendMsgReply{
		ServerMsgSeq: r.MsgSeq,
		ResCode:      r.ResCode,
	}, nil
}
func (s *ChatService) GetSyncMessage(ctx context.Context, req *v1.SyncMessageRequest) (*v1.SyncMessageReply, error) {
	return &v1.SyncMessageReply{}, nil
}
func (s *ChatService) GetRecentContactSession(ctx context.Context, req *v1.GetRecentSessionRequest) (*v1.GetRecentSessionReply, error) {
	return &v1.GetRecentSessionReply{}, nil
}
func (s *ChatService) GetMsgList(ctx context.Context, req *v1.GetMsgListRequest) (*v1.GetMsgListReply, error) {
	return &v1.GetMsgListReply{}, nil
}
func (s *ChatService) MsgReadAck(ctx context.Context, req *v1.MsgReadAckRequest) (*v1.MsgReadAckReply, error) {
	return &v1.MsgReadAckReply{}, nil
}

func (s *ChatService) parseSession(to string) pb.IMSessionType {
	const (
		kGroupPrefix  = "group-"
		kSinglePrefix = "user-"
	)
	if strings.Index(to, kGroupPrefix) != -1 {
		return pb.IMSessionType_kCIM_SESSION_TYPE_GROUP
	} else if strings.Index(to, kSinglePrefix) != -1 {
		return pb.IMSessionType_kCIM_SESSION_TYPE_SINGLE
	}
	return pb.IMSessionType_kCIM_SESSION_TYPE_Invalid
}
