package biz

import (
	pb "chat/api/chat"
	"chat/internal/data"
	"chat/internal/data/cache"
	"chat/internal/data/ent"
	"context"
	"errors"
	"github.com/gomicroim/gomicroim/v2/pkg/log"
	"strconv"
	"time"
)

type MessageUseCase struct {
	msgRepo     data.MessageRepo
	sessionRepo data.SessionRepo
	seqCache    cache.MsgSeq

	log *log.Logger
}

func NewMessageUseCase(repo data.MessageRepo, seq cache.MsgSeq, sessionRepo data.SessionRepo) *MessageUseCase {
	return &MessageUseCase{
		msgRepo:     repo,
		seqCache:    seq,
		sessionRepo: sessionRepo,
		log:         log.L,
	}
}

func (m *MessageUseCase) Send(ctx context.Context, from int64, to string, sessionType int, clientMsgID string,
	msgType int8, msgData string) (*data.Message, error) {
	// 幂等，如果由于网络等问题，ack客户端没有收到，则下次重发不必再插入数据库
	msg, err := m.msgRepo.FindByClientMsgId(ctx, clientMsgID)
	if err != nil && !ent.IsNotFound(err) {
		return msg, err
	}

	if sessionType == int(pb.IMSessionType_kCIM_SESSION_TYPE_SINGLE) {
		return m.send(ctx, from, to, clientMsgID, msgType, msgData)
	} else if sessionType == int(pb.IMSessionType_kCIM_SESSION_TYPE_GROUP) {
		return m.sendGroup(ctx, from, to, clientMsgID, msgType, msgData)
	} else {
		return nil, errors.New("invalid sessionType")
	}
}

func (m *MessageUseCase) GetMessageList(ctx context.Context, userId int64, peerId string,
	sessionType pb.IMSessionType, isForward bool, msgSeq int64, limit int) ([]*data.Message, error) {
	if isForward {
		endMsgSeq := msgSeq
		return m.msgRepo.ListByEndMsgSeq(ctx, m.msgRepo.GetSessionKey(userId, peerId, sessionType), endMsgSeq, limit)
	}
	startMsgSeq := msgSeq
	return m.msgRepo.ListByStartMsgSeq(ctx, m.msgRepo.GetSessionKey(userId, peerId, sessionType), startMsgSeq, limit)
}

func (m *MessageUseCase) send(ctx context.Context, from int64, to string, clientMsgID string,
	msgType int8, msgData string) (*data.Message, error) {

	sessionType := pb.IMSessionType_kCIM_SESSION_TYPE_SINGLE
	fromStr := strconv.FormatInt(from, 10)

	// check session
	sessions, err := m.sessionRepo.FindSingleSession(ctx, fromStr, to, sessionType)
	if err != nil {
		return nil, err
	}
	if len(sessions) <= 0 {
		err = m.sessionRepo.Create(ctx, &data.Session{
			UserId:        strconv.FormatInt(from, 10),
			PeerId:        to,
			SessionType:   sessionType,
			SessionStatus: pb.IMSessionStatusType_kCIM_SESSION_STATUS_OK,
		})
		if err != nil {
			return nil, err
		}
	} else {
		// update session latest chat time
		if len(sessions) == 1 {
			m.log.Warn("single session miss row")
		}
		for _, v := range sessions {
			if v.SessionStatus == pb.IMSessionStatusType_kCIM_SESSION_STATUS_OK {
				continue
			}
			_, err = m.sessionRepo.UpdateUpdated(ctx, v.Id, time.Now(), pb.IMSessionStatusType_kCIM_SESSION_STATUS_OK)
			if err != nil {
				return nil, err
			}
		}
	}

	// get session autoincrement msg_seq
	msgSeq, err := m.seqCache.IncrSingle(ctx, fromStr, to)
	if err != nil {
		return nil, err
	}

	return m.msgRepo.Create(ctx, &data.Message{
		From:         from,
		To:           to,
		SessionKey:   m.msgRepo.GetSessionKey(from, to, sessionType),
		SessionType:  int(sessionType),
		ClientMsgID:  clientMsgID,
		ServerMsgSeq: msgSeq,
		MsgType:      int(msgType),
		MsgData:      msgData,
		MsgResCode:   int(pb.IMResCode_kCIM_RES_CODE_OK),
		MsgFeature:   int(pb.IMMsgFeature_kCIM_MSG_FEATURE_ROAM_MSG),
		MsgStatus:    int(pb.IMMsgStatus_kCIM_MSG_STATUS_NONE),
	})
}

func (m *MessageUseCase) sendGroup(ctx context.Context, from int64, groupId string,
	clientMsgID string, msgType int8, msgData string) (*data.Message, error) {
	sessionType := pb.IMSessionType_kCIM_SESSION_TYPE_GROUP
	fromStr := strconv.FormatInt(from, 10)

	session, err := m.sessionRepo.FindOne(ctx, fromStr, groupId, pb.IMSessionType_kCIM_SESSION_TYPE_GROUP)
	if err != nil {
		if ent.IsNotFound(err) {
			err = m.sessionRepo.Create(ctx, &data.Session{
				UserId:        fromStr,
				PeerId:        groupId,
				SessionType:   pb.IMSessionType_kCIM_SESSION_TYPE_GROUP,
				SessionStatus: pb.IMSessionStatusType_kCIM_SESSION_STATUS_OK,
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// update session last chat time
		if _, err = m.sessionRepo.UpdateUpdated(ctx, session.Id, time.Now(), pb.IMSessionStatusType_kCIM_SESSION_STATUS_OK); err != nil {
			return nil, err
		}
	}

	// get session autoincrement msg_seq
	msgSeq, err := m.seqCache.IncrGroup(ctx, groupId)
	if err != nil {
		return nil, err
	}

	return m.msgRepo.Create(ctx, &data.Message{
		From:         from,
		To:           groupId,
		SessionType:  int(sessionType),
		ClientMsgID:  clientMsgID,
		ServerMsgSeq: msgSeq,
		MsgType:      int(msgType),
		MsgData:      msgData,
		MsgResCode:   int(pb.IMResCode_kCIM_RES_CODE_OK),
		MsgFeature:   int(pb.IMMsgFeature_kCIM_MSG_FEATURE_ROAM_MSG),
		MsgStatus:    int(pb.IMMsgStatus_kCIM_MSG_STATUS_NONE),
	})
}
