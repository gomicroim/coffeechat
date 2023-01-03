package data

import (
	"chat/internal/data/ent"
	"chat/internal/data/ent/message"
	"context"
	"fmt"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/gomicroim/gomicroim/protos/wspush"
	"strconv"
	"time"
)

type Message struct {
	ID           int64     // 自增ID
	Created      time.Time // 创建时间
	Updated      time.Time // 更新时间
	From         int64
	To           string
	SessionKey   string // 会话标志，用来查询历史聊天消息。私聊是smallId:bigId,群聊是groupId
	SessionType  int
	ClientMsgID  string
	ServerMsgSeq int64
	MsgType      int
	MsgData      string
	MsgResCode   int
	MsgFeature   int
	MsgStatus    int
}

func (m Message) String() string {
	return fmt.Sprintf("from:%d,to:%s,sessionType:%d,clientMsgId:%s,msgSeq:%d,msgType:%d,msgData:%s",
		m.From, m.To, m.SessionType, m.ClientMsgID, m.ServerMsgSeq, m.MsgType, m.MsgData)
}

type MessageRepo interface {
	// GetSessionKey 根据 from&to 生产一个唯一key，以方便查询历史聊天消息
	GetSessionKey(from int64, to string, sessionType wspush.IMSessionType) string
	Create(context.Context, *Message) (*Message, error)
	FindByClientMsgId(ctx context.Context, clientMsgId string) (*Message, error)
	ListByStartMsgSeq(ctx context.Context, sessionKey string, startMsgSeq int64, limit int) ([]*Message, error)
	ListByEndMsgSeq(ctx context.Context, sessionKey string, endMsgSeq int64, limit int) ([]*Message, error)
}

type messageRepo struct {
	log    *log.Logger
	client *ent.Client
}

func NewMessageRepo(data *Data, logger *log.Logger) MessageRepo {
	return &messageRepo{
		client: data.EntClient,
		log:    logger,
	}
}

func (m *messageRepo) Create(ctx context.Context, msg *Message) (*Message, error) {
	result, err := m.client.Message.Create().
		SetSessionKey(msg.SessionKey).SetFrom(msg.From).SetTo(msg.To).SetSessionType(int8(msg.SessionType)).
		SetClientMsgID(msg.ClientMsgID).SetServerMsgSeq(msg.ServerMsgSeq).
		SetMsgType(int8(msg.MsgType)).SetMsgData(msg.MsgData).
		SetMsgResCode(int8(msg.MsgResCode)).SetMsgStatus(int8(msg.MsgStatus)).SetMsgFeature(int8(msg.MsgFeature)).Save(ctx)
	if err != nil {
		return nil, err
	}
	msg.ID = result.ID
	msg.Created = result.Created
	msg.Updated = result.Updated
	return msg, nil
}

func (m *messageRepo) ent2model(old *ent.Message) *Message {
	return &Message{
		ID:           old.ID,
		Created:      old.Created,
		Updated:      old.Updated,
		From:         old.From,
		To:           old.To,
		SessionType:  int(old.SessionType),
		ClientMsgID:  old.ClientMsgID,
		ServerMsgSeq: old.ServerMsgSeq,
		MsgType:      int(old.MsgType),
		MsgData:      old.MsgData,
		MsgResCode:   int(old.MsgResCode),
		MsgFeature:   int(old.MsgFeature),
		MsgStatus:    int(old.MsgStatus),
	}
}

func (m *messageRepo) FindByClientMsgId(ctx context.Context, clientMsgId string) (*Message, error) {
	result, err := m.client.Message.Query().Where(message.ClientMsgID(clientMsgId)).Only(ctx)
	if err != nil {
		return nil, err
	}
	return m.ent2model(result), nil
}

func (m *messageRepo) ListByStartMsgSeq(ctx context.Context, sessionKey string, startMsgSeq int64, limit int) ([]*Message, error) {
	arr, err := m.client.Message.Query().
		Where(message.SessionKey(sessionKey), message.ServerMsgSeqGT(startMsgSeq)).
		Order(ent.Asc(message.FieldServerMsgSeq)).
		Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	messages := make([]*Message, len(arr))
	for k, v := range arr {
		messages[k] = m.ent2model(v)
	}
	return messages, nil
}

func (m *messageRepo) ListByEndMsgSeq(ctx context.Context, sessionKey string, endMsgSeq int64, limit int) ([]*Message, error) {
	arr, err := m.client.Message.Query().
		Where(message.SessionKey(sessionKey), message.ServerMsgSeqLT(endMsgSeq)).
		Order(ent.Desc(message.FieldServerMsgSeq)).
		Limit(limit).All(ctx)
	if err != nil {
		return nil, err
	}
	messages := make([]*Message, len(arr))
	for k, v := range arr {
		messages[k] = m.ent2model(v)
	}
	return messages, nil
}

func (m *messageRepo) GetSessionKey(from int64, to string, sessionType wspush.IMSessionType) string {
	if sessionType == wspush.IMSessionType_SessionTypeSingle {
		fromStr := strconv.FormatInt(from, 10)
		small, big := fromStr, to
		if fromStr > to {
			small, big = to, fromStr
		}
		return fmt.Sprintf("single:%s:%s", small, big)
	}
	// 群直接根据groupId查即可
	return to
}
