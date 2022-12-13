package service

import (
	pb "chat/api"
	"chat/internal/biz"
	"context"
	"github.com/gomicroim/gomicroim/protos/chat"
)

type ChatService struct {
	pb.UnimplementedChatServer

	messageBiz *biz.MessageUseCase
	sessionBiz *biz.RecentSessionUseCase
}

func NewChatService(msgBiz *biz.MessageUseCase, sessionBiz *biz.RecentSessionUseCase) *ChatService {
	return &ChatService{messageBiz: msgBiz, sessionBiz: sessionBiz}
}

func (s *ChatService) SendMsg(ctx context.Context, req *pb.SendMsgRequest) (*pb.SendMsgReply, error) {
	msg, err := s.messageBiz.Send(ctx, req.FromUserId, req.To, int(req.SessionType), req.ClientMsgId, int8(req.MsgType), req.MsgData)
	if err != nil {
		return nil, err
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
