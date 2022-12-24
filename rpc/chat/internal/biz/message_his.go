package biz

import (
	"chat/internal/data"
	"chat/internal/data/cache"
	"chat/internal/data/ent"
	"context"
	"errors"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/gomicroim/gomicroim/protos/chat"
	"strconv"
	"time"
)

// MessageHistoryUseCase 读扩散，历史聊天消息管理，存储到mysql
type MessageHistoryUseCase interface {
	Send(ctx context.Context, from int64, to string, sessionType chat.IMSessionType, clientMsgID string,
		msgType chat.IMMsgType, msgData string) (*data.Message, error)
	GetMessageList(ctx context.Context, userId int64, peerId string,
		sessionType chat.IMSessionType, isForward bool, msgSeq int64, limit int) ([]*data.Message, error)
}

type messageHistoryUseCase struct {
	msgRepo     data.MessageRepo
	sessionRepo data.SessionRepo
	seqCache    cache.MsgSeq

	log *log.Logger
}

func NewMessageHistoryUseCase(repo data.MessageRepo, seq cache.MsgSeq, sessionRepo data.SessionRepo) MessageHistoryUseCase {
	return &messageHistoryUseCase{
		msgRepo:     repo,
		seqCache:    seq,
		sessionRepo: sessionRepo,
		log:         log.L,
	}
}

func (m *messageHistoryUseCase) Send(ctx context.Context, from int64, to string, sessionType chat.IMSessionType, clientMsgID string,
	msgType chat.IMMsgType, msgData string) (*data.Message, error) {

	// 幂等，如果由于网络等问题，ack客户端没有收到，则下次重发不必再插入数据库
	msg, err := m.msgRepo.FindByClientMsgId(ctx, clientMsgID)
	if err != nil && !ent.IsNotFound(err) {
		return msg, err
	}

	switch sessionType {
	case chat.IMSessionType_SessionTypeSingle:
		return m.sendSingle(ctx, from, to, clientMsgID, int8(msgType), msgData)
	case chat.IMSessionType_SessionTypeNormalGroup:
		return m.sendGroup(ctx, from, to, clientMsgID, int8(msgType), msgData)
	case chat.IMSessionType_SessionTypeSuperGroup:
		return nil, errors.New("not support super group")
	default:
		return nil, errors.New("invalid sessionType")
	}
}

func (m *messageHistoryUseCase) GetMessageList(ctx context.Context, userId int64, peerId string,
	sessionType chat.IMSessionType, isForward bool, msgSeq int64, limit int) ([]*data.Message, error) {
	if isForward {
		endMsgSeq := msgSeq
		return m.msgRepo.ListByEndMsgSeq(ctx, m.msgRepo.GetSessionKey(userId, peerId, sessionType), endMsgSeq, limit)
	}
	startMsgSeq := msgSeq
	return m.msgRepo.ListByStartMsgSeq(ctx, m.msgRepo.GetSessionKey(userId, peerId, sessionType), startMsgSeq, limit)
}

func (m *messageHistoryUseCase) sendSingle(ctx context.Context, from int64, to string, clientMsgID string,
	msgType int8, msgData string) (*data.Message, error) {

	sessionType := chat.IMSessionType_SessionTypeSingle
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
			SessionStatus: chat.IMSessionStatus_SessionStatusOk,
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
			if v.SessionStatus == chat.IMSessionStatus_SessionStatusDelete {
				continue
			}
			_, err = m.sessionRepo.UpdateUpdated(ctx, v.Id, time.Now(), chat.IMSessionStatus_SessionStatusOk)
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
		MsgResCode:   int(chat.IMResCode_ResCodeOk),
		MsgFeature:   int(chat.IMMsgFeature_MsgFeatureRoamMsg),
		MsgStatus:    int(chat.IMMsgStatus_MsgStatusNone),
	})
}

func (m *messageHistoryUseCase) sendGroup(ctx context.Context, from int64, groupId string,
	clientMsgID string, msgType int8, msgData string) (*data.Message, error) {

	sessionType := chat.IMSessionType_SessionTypeNormalGroup
	fromStr := strconv.FormatInt(from, 10)

	session, err := m.sessionRepo.FindOne(ctx, fromStr, groupId, chat.IMSessionType_SessionTypeNormalGroup)
	if err != nil {
		if ent.IsNotFound(err) {
			err = m.sessionRepo.Create(ctx, &data.Session{
				UserId:        fromStr,
				PeerId:        groupId,
				SessionType:   chat.IMSessionType_SessionTypeNormalGroup,
				SessionStatus: chat.IMSessionStatus_SessionStatusOk,
			})
			if err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else {
		// update session last chat time
		if _, err = m.sessionRepo.UpdateUpdated(ctx, session.Id, time.Now(), chat.IMSessionStatus_SessionStatusOk); err != nil {
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
		MsgResCode:   int(chat.IMResCode_ResCodeOk),
		MsgFeature:   int(chat.IMMsgFeature_MsgFeatureRoamMsg),
		MsgStatus:    int(chat.IMMsgStatus_MsgStatusNone),
	})
}
