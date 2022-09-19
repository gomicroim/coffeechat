package service

import (
	"chat/api/chat"
	"context"
	"github.com/gomicroim/gomicroim/v2/pkg/jwt"
	"github.com/gomicroim/gomicroim/v2/pkg/log"
	"go.uber.org/zap"
	"strconv"
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
		MsgInfo: &v1.IMMsgInfo{
			ClientMsgId: r.MsgInfo.ClientMsgId,
			ServerMsgId: r.MsgInfo.ServerMsgSeq,
			MsgResCode:  r.MsgInfo.MsgResCode,
			MsgFeature:  r.MsgInfo.MsgFeature,
			SessionType: r.MsgInfo.SessionType,
			FromUserId:  r.MsgInfo.FromUserId,
			PeerId:      r.MsgInfo.To,
			CreateTime:  r.MsgInfo.CreateTime,
			MsgType:     r.MsgInfo.MsgType,
			MsgStatus:   r.MsgInfo.MsgStatus,
			MsgData:     r.MsgInfo.MsgData,
		},
	}, nil
}
func (s *ChatService) GetSyncMessage(ctx context.Context, req *v1.SyncMessageRequest) (*v1.SyncMessageReply, error) {
	return &v1.SyncMessageReply{}, nil
}
func (s *ChatService) GetRecentContactSession(ctx context.Context, req *v1.GetRecentSessionRequest) (*v1.GetRecentSessionReply, error) {
	return &v1.GetRecentSessionReply{}, nil
}
func (s *ChatService) GetMsgList(ctx context.Context, req *v1.GetMsgListRequest) (*v1.GetMsgListReply, error) {
	sessionType := s.parseSession(req.PeerId)
	userId := jwt.GetUserId(ctx)

	result, err := s.chatClient.GetMsgList(ctx, &pb.GetMsgListRequest{
		UserId:      userId,
		SessionType: sessionType,
		PeerId:      req.PeerId,
		Filter: &pb.GetMsgListRequest_GetMsgBySeq{
			MsgSeq:    req.MsgSeq,
			IsForward: req.IsForward,
		},
		LimitCount: req.LimitCount,
	})
	if err != nil {
		return nil, err
	}

	msgList := make([]*v1.IMMsgInfo, len(result.MsgList))

	// FIXME fill BFF message
	for k, v := range result.MsgList {
		msgList[k] = &v1.IMMsgInfo{
			ClientMsgId: v.ClientMsgId,
			ServerMsgId: v.ServerMsgSeq,
			MsgResCode:  v.MsgResCode,
			MsgFeature:  v.MsgFeature,
			SessionType: v.SessionType,
			FromUserId:  v.FromUserId,
			PeerId:      v.To,
			CreateTime:  v.CreateTime,
			MsgType:     v.MsgType,
			MsgStatus:   v.MsgStatus,
			MsgData:     v.MsgData,
		}
	}

	return &v1.GetMsgListReply{
		EndMsgSeq: result.EndMsgSeq,
		MsgList:   msgList,
	}, nil
}
func (s *ChatService) MsgReadAck(ctx context.Context, req *v1.MsgReadAckRequest) (*v1.MsgReadAckReply, error) {
	return &v1.MsgReadAckReply{}, nil
}

func (s *ChatService) parseSession(to string) pb.IMSessionType {
	const (
		kGroupPrefix = "group-"
	)
	if strings.Index(to, kGroupPrefix) != -1 {
		return pb.IMSessionType_kCIM_SESSION_TYPE_GROUP
	} else if _, err := strconv.ParseInt(to, 10, 64); err == nil {
		return pb.IMSessionType_kCIM_SESSION_TYPE_SINGLE
	}
	return pb.IMSessionType_kCIM_SESSION_TYPE_Invalid
}
