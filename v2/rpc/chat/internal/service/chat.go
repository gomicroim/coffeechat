package service

import (
	"chat/internal/biz"
	"context"
	"errors"
	"strconv"

	pb "chat/api/chat"
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
		ResCode: pb.IMResCode(msg.MsgResCode),
	}, nil
}

func (s *ChatService) GetRecentContactSession(ctx context.Context, req *pb.GetRecentSessionRequest) (*pb.GetRecentSessionReply, error) {
	list, err := s.sessionBiz.GetSessionList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	session := &pb.GetRecentSessionReply{UserId: req.UserId, ContactSessionList: make([]*pb.IMContactSessionInfo, len(list))}
	for k, v := range list {
		session.ContactSessionList[k] = &pb.IMContactSessionInfo{
			SessionId:     v.Id,
			PeerId:        v.PeerId,
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
	if req.TimeSpan != nil {
		return nil, errors.New("not support timespan filter")
	}
	if req.Filter == nil {
		return nil, errors.New("miss filter filed")
	}
	result, err := s.messageBiz.GetMessageList(ctx, req.UserId, strconv.FormatInt(req.PeerId, 10), req.SessionType,
		req.Filter.IsForward, req.Filter.MsgSeq, int(req.LimitCount))
	if err != nil {
		return nil, err
	}
	out := &pb.GetMsgListReply{
		EndMsgSeq: result[len(result)-1].ServerMsgSeq,
		MsgList:   make([]*pb.IMMsgInfo, len(result)),
	}
	for k, v := range result {
		out.MsgList[k] = &pb.IMMsgInfo{
			FromUserId:   v.From,
			To:           v.To,
			SessionType:  pb.IMSessionType(v.SessionType),
			ClientMsgId:  v.ClientMsgID,
			ServerMsgSeq: v.ServerMsgSeq,
			MsgType:      pb.IMMsgType(v.MsgType),
			MsgData:      v.MsgData,
			MsgResCode:   pb.IMResCode(v.MsgResCode),
			MsgFeature:   pb.IMMsgFeature(v.MsgFeature),
			MsgStatus:    pb.IMMsgStatus(v.MsgStatus),
			CreateTime:   v.Created.Unix(),
		}
	}
	return out, nil
}

func (s *ChatService) MsgReadAck(ctx context.Context, req *pb.MsgReadAckRequest) (*pb.MsgReadAckReply, error) {
	return &pb.MsgReadAckReply{}, nil
}
