package service

import (
	"chat/internal/biz"
	"context"
	"errors"
	"github.com/gomicroim/gomicroim/protos/chat"

	pb "chat/api"
)

type MsgListService struct {
	pb.UnimplementedMsgListServer

	msgList *biz.MessageUseCase
}

func NewMsgListService(msgList *biz.MessageUseCase) *MsgListService {
	return &MsgListService{msgList: msgList}
}

func (s *MsgListService) GetMsgList(ctx context.Context, req *pb.GetMsgListRequest) (*pb.GetMsgListReply, error) {
	if req.Filter == nil {
		return nil, errors.New("miss filter filed")
	}
	result, err := s.msgList.GetMessageList(ctx, req.UserId, req.PeerId, req.SessionType,
		req.Filter.IsForward, req.Filter.MsgSeq, int(req.LimitCount))
	if err != nil {
		return nil, err
	}

	out := &pb.GetMsgListReply{
		EndMsgSeq: req.Filter.MsgSeq,
		MsgList:   make([]*pb.IMBaseMsg, len(result)),
	}
	if len(result) > 0 {
		out.EndMsgSeq = result[len(result)-1].ServerMsgSeq
		for k, v := range result {
			out.MsgList[k] = &pb.IMBaseMsg{
				FromUserId:   v.From,
				To:           v.To,
				SessionType:  chat.IMSessionType(v.SessionType),
				ClientMsgId:  v.ClientMsgID,
				ServerMsgSeq: v.ServerMsgSeq,
				MsgType:      chat.IMMsgType(v.MsgType),
				MsgData:      v.MsgData,
				MsgResCode:   chat.IMResCode(v.MsgResCode),
				MsgFeature:   chat.IMMsgFeature(v.MsgFeature),
				MsgStatus:    chat.IMMsgStatus(v.MsgStatus),
				CreateTime:   v.Created.Unix(),
			}
		}
	}
	return out, nil
}
