package service

import (
	"chat/internal/biz"
	"context"

	pb "chat/api/chat"
)

type ChatService struct {
	pb.UnimplementedChatServer

	messageBiz *biz.MessageUseCase
}

func NewChatService(msgBiz *biz.MessageUseCase) *ChatService {
	return &ChatService{messageBiz: msgBiz}
}

func (s *ChatService) SendMsg(ctx context.Context, req *pb.SendMsgRequest) (*pb.SendMsgReply, error) {
	msg, err := s.messageBiz.Send(ctx, req.FromUserId, req.To, int(req.SessionType), req.ClientMsgId, int8(req.MsgType), req.MsgData)
	if err != nil {
		return nil, err
	}
	return &pb.SendMsgReply{
		MsgSeq:  uint64(msg.ServerMsgSeq),
		ResCode: pb.IMResCode(msg.MsgResCode),
	}, nil
}

func (s *ChatService) GetSyncMessage(ctx context.Context, req *pb.SyncMessageRequest) (*pb.SyncMessageReply, error) {
	return &pb.SyncMessageReply{}, nil
}

func (s *ChatService) GetRecentContactSession(ctx context.Context, req *pb.GetRecentSessionRequest) (*pb.GetRecentSessionReply, error) {
	list, err := s.messageBiz.GetSessionList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	session := &pb.GetRecentSessionReply{UserId: req.UserId, ContactSessionList: make([]*pb.IMContactSessionInfo, len(list))}
	for k, v := range list {
		session.ContactSessionList[k] = &pb.IMContactSessionInfo{
			SessionId:     uint64(v.Id),
			SessionType:   v.SessionType,
			SessionStatus: v.SessionStatus,
			UnreadCnt:     0,
			UpdatedTime:   0,
			LatestMsgId:   "",
			LatestMsgSeq:  0,
			MsgTimeStamp:  0,
			MsgData:       "",
			MsgType:       0,
			MsgFromUserId: 0,
			MsgStatus:     0,
		}
	}
	return session, nil
}

func (s *ChatService) GetMsgList(ctx context.Context, req *pb.GetMsgListRequest) (*pb.GetMsgListReply, error) {
	return &pb.GetMsgListReply{}, nil
}

func (s *ChatService) MsgReadAck(ctx context.Context, req *pb.MsgReadAckRequest) (*pb.MsgReadAckReply, error) {
	return &pb.MsgReadAckReply{}, nil
}
