package service

import (
	"chat/internal/biz"
	"context"

	pb "chat/api"
)

type SessionService struct {
	pb.UnimplementedSessionServer
	sessionBiz *biz.RecentSessionUseCase
}

func NewSessionService(sessionBiz *biz.RecentSessionUseCase) *SessionService {
	return &SessionService{sessionBiz: sessionBiz}
}

func (s *SessionService) GetSession(ctx context.Context, req *pb.GetSessionRequest) (*pb.GetRecentSessionReply, error) {
	list, err := s.sessionBiz.GetSessionList(ctx, req.UserId)
	if err != nil {
		return nil, err
	}

	session := &pb.GetRecentSessionReply{UserId: req.UserId, ContactSessionList: make([]*pb.IMSessionInfo, len(list))}
	for k, v := range list {
		session.ContactSessionList[k] = &pb.IMSessionInfo{
			SessionId:     v.Id,
			PeerId:        v.PeerId,
			SessionType:   v.SessionType,
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

func (s *SessionService) ReadMsgNotify(ctx context.Context, req *pb.MsgReadAckRequest) (*pb.MsgReadAckReply, error) {
	return &pb.MsgReadAckReply{}, nil
}
