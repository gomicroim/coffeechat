package service

import (
	"chat/api"
	"context"
	"github.com/gomicroim/gomicroim/pkg/jwt"
	"github.com/gomicroim/gomicroim/pkg/log"
	v1 "github.com/gomicroim/gomicroim/protos/api"
	pb "github.com/gomicroim/gomicroim/protos/wspush"
	"go.uber.org/zap"
)

type ApiChatService struct {
	v1.UnimplementedChatServer

	chatClient chat.ChatClient
	log        *log.Logger
}

func NewApiChatService(client chat.ChatClient) *ApiChatService {
	return &ApiChatService{chatClient: client, log: log.L}
}

func (s *ApiChatService) SendMsg(ctx context.Context, req *v1.SendMsgRequest) (*v1.SendMsgReply, error) {
	sessionType := req.To.SessionType
	userId := jwt.GetUserId(ctx)

	r, err := s.chatClient.SendMsg(ctx, &chat.SendMsgRequest{
		FromUserId:  userId,
		To:          req.To.PeerId,
		SessionType: sessionType,
		ClientMsgId: req.MsgUuid,
		MsgType:     req.MsgType,
		MsgData:     req.MsgData,
	})
	if err != nil {
		return nil, err
	}
	s.log.Info("SendMsg", zap.Int64("from", userId), zap.Any("to", req.To),
		zap.Any("msgType", req.MsgType), zap.String("msgData", req.MsgData),
		zap.Int64("msgSeq", r.MsgSeq))

	return &v1.SendMsgReply{
		MsgSeq:  r.MsgSeq,
		ResCode: r.ResCode,
		MsgInfo: s.baseMsg2ImMsg(r.MsgInfo),
	}, nil
}
func (s *ApiChatService) GetSyncMessage(ctx context.Context, req *v1.SyncMessageRequest) (*v1.SyncMessageReply, error) {
	result, err := s.chatClient.SyncMessage(ctx, &chat.SyncMessageRequest{
		Member:   req.Member,
		LastRead: req.LastRead,
		Count:    req.Count,
	})
	if err != nil {
		return nil, err
	}
	reply := &v1.SyncMessageReply{
		LatestSeq:       result.LatestSeq,
		EntrySetLastSeq: result.EntrySetLastSeq,
		EntrySet:        make([]*v1.SyncMessageReply_TimelineEntry, 0),
	}
	for _, v := range result.EntrySet {
		entry := &v1.SyncMessageReply_TimelineEntry{
			Sequence: v.Sequence,
			Message:  s.baseMsg2ImMsg(v.Message),
		}
		reply.EntrySet = append(reply.EntrySet, entry)
	}
	return reply, nil
}
func (s ApiChatService) baseMsg2ImMsg(msg *chat.IMBaseMsg) *pb.IMMsg {
	out := &pb.IMMsg{
		From: &pb.IMMsg_From{
			UserId:   msg.FromUserId,
			NickName: "",
		},
		Peer: &pb.IMMsg_Peer{
			SessionType: msg.SessionType,
			PeerId:      msg.To,
			PeerName:    nil,
		},
		MsgType:    msg.MsgType,
		MsgData:    msg.MsgData,
		MsgId:      msg.ClientMsgId,
		MsgSeq:     msg.ServerMsgSeq,
		MsgResCode: msg.MsgResCode,
		MsgFeature: msg.MsgFeature,
		MsgStatus:  msg.MsgStatus,
		CreateTime: msg.CreateTime,
	}
	return out
}

func (s *ApiChatService) GetRecentContactSession(ctx context.Context, req *v1.GetRecentSessionRequest) (*v1.GetRecentSessionReply, error) {
	return &v1.GetRecentSessionReply{}, nil
}

//func (s *ApiChatService) GetMsgList(ctx context.Context, req *v1.GetMsgListRequest) (*v1.GetMsgListReply, error) {
//	sessionType := s.parseSession(req.PeerId)
//	userId := jwt.GetUserId(ctx)
//
//	result, err := s.chatClient.GetMsgList(ctx, &chat.GetMsgListRequest{
//		UserId:      userId,
//		SessionType: sessionType,
//		PeerId:      req.PeerId,
//		Filter: &chat.GetMsgListRequest_GetMsgBySeq{
//			MsgSeq:    req.MsgSeq,
//			IsForward: req.IsForward,
//		},
//		LimitCount: req.LimitCount,
//	})
//	if err != nil {
//		return nil, err
//	}
//
//	msgList := make([]*v1.IMMsgInfo, len(result.MsgList))
//
//	// FIXME fill BFF message
//	for k, v := range result.MsgList {
//		msgList[k] = &v1.IMMsgInfo{
//			ClientMsgId: v.ClientMsgId,
//			ServerMsgId: v.ServerMsgSeq,
//			MsgResCode:  v.MsgResCode,
//			MsgFeature:  v.MsgFeature,
//			SessionType: v.SessionType,
//			FromUserId:  v.FromUserId,
//			PeerId:      v.To,
//			CreateTime:  v.CreateTime,
//			MsgType:     v.MsgType,
//			MsgStatus:   v.MsgStatus,
//			MsgData:     v.MsgData,
//		}
//	}
//
//	return &v1.GetMsgListReply{
//		EndMsgSeq: result.EndMsgSeq,
//		MsgList:   msgList,
//	}, nil
//}

func (s *ApiChatService) MsgReadAck(ctx context.Context, req *v1.MsgReadAckRequest) (*v1.MsgReadAckReply, error) {
	return &v1.MsgReadAckReply{}, nil
}
