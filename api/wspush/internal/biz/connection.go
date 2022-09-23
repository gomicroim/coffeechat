package biz

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/gomicroim/gomicroim/pkg/rescue"
	"go.uber.org/atomic"
	"net"
	"sync"
	"time"
	"wspush/api/wspush"

	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"google.golang.org/protobuf/proto"
)

var (
	errorUnknownClientMessage = errors.New("unknown client message")
)

const (
	KWriteWaitTime          = 5 * time.Second // 写超时
	KWriteChannelSize       = 10
	KHeartbeatTimeOut       = 2 * 60 * time.Second // 120s
	KMaxWaitAckTime         = 10 * time.Second
	KBatchSendTimeInterval  = 500 * time.Millisecond // 批量推送等待间隔 500 ms
	KBatchSendPacketMaxSize = 10 * 1024              // 批量推送等待最大消息包数
)

var (
	ErrWriteChannelFullMsgDropped = errors.New("write buffer channel is full, drop msg")
	ErrHeartBeatTimeout           = errors.New("heartbeat timeout")
	ErrReasonAckTimeout           = errors.New("wait ack timeout")
)

// MsgRequest notify and wait ack response
type MsgRequest struct {
	Data     *wspush.S2CWebsocketMessage
	Timeout  time.Duration
	Callback ResponseCallback

	createTime time.Time
}

// Response client ack response
type Response struct {
	Req       *MsgRequest
	IsSuccess bool
	ErrReason error
}

type ResponseCallback = func(res *Response)

type Connection interface {
	// Close the websocket connection
	Close() error
	// Write thread safe, write websocket.BinaryMessage to the remote
	Write(msg []byte) error
	// SendRequest thread safe, write websocket.BinaryMessage to the remote and wait ack response
	SendRequest(data *wspush.ServerMessage, timeout time.Duration, callback ResponseCallback) error
	// SendNotify thread safe, write websocket.BinaryMessage to the remote without wait ack response
	SendNotify(data *wspush.ServerMessage) error
	ClientInfo() ClientInfo
}

type connectionImpl struct {
	conn        *websocket.Conn
	clientInfo  ClientInfo
	connPoolKey string
	//msgPool *sync.Pool

	notifyWriter chan []byte      // 不需要ack的channel
	reqWriter    chan *MsgRequest // 需要ack确认
	logger       *log.Logger

	heartbeatCh       chan struct{}
	lastHeartbeatTime time.Time

	seq        atomic.Int32
	ackMutex   sync.RWMutex
	waitAckMsg map[int32]*MsgRequest
}

// FD get socket fd
func FD(c *websocket.Conn) int32 {
	fd := int32(0)
	if tcpConn, ok := c.UnderlyingConn().(*net.TCPConn); ok {
		rawConn, err := tcpConn.SyscallConn()
		if err != nil {
			return fd
		}

		_ = rawConn.Control(func(f uintptr) { fd = int32(f) })
	}
	return fd
}

// NewConnection new connection
func NewConnection(conn *websocket.Conn, clientInfo ClientInfo, logger *log.Logger) *connectionImpl {
	connection := &connectionImpl{
		conn:              conn,
		clientInfo:        clientInfo,
		notifyWriter:      make(chan []byte, KWriteChannelSize),
		reqWriter:         make(chan *MsgRequest, KWriteChannelSize),
		heartbeatCh:       make(chan struct{}),
		waitAckMsg:        make(map[int32]*MsgRequest),
		lastHeartbeatTime: time.Now(),
		logger:            logger,
	}
	//connection.msgPool = &sync.Pool{New: func() interface{} {
	//	logger.Debug("pool request a new *MsgRequest instance")
	//	return &MsgRequest{
	//		Data: &wspush.S2CServerMessage{
	//			Header:   &wspush.WebSocketHeader{},
	//			DataList: make([]*wspush.ServerMessage, 0),
	//		},
	//		Timeout:    0,
	//		Callback:   nil,
	//		createTime: time.Time{},
	//	}
	//}}
	return connection
}

func (c *connectionImpl) Close() error {
	return c.conn.Close()
}

func (c *connectionImpl) Write(data []byte) error {
	defer func() { recover() }()

	// if buffer full, return error immediate
	select {
	case c.notifyWriter <- data:
	default:
		return ErrWriteChannelFullMsgDropped
	}
	return nil
}

func (c *connectionImpl) ClientInfo() ClientInfo {
	return c.clientInfo
}

func (c *connectionImpl) SendRequest(wsMsg *wspush.ServerMessage, timeout time.Duration, callback ResponseCallback) error {
	// To avoid the panic "send on closed channel"
	defer func() { recover() }()

	//req := c.msgPool.Get().(*MsgRequest)
	req := &MsgRequest{}
	req.Data = &wspush.S2CWebsocketMessage{
		Header:   &wspush.WebSocketHeader{Seq: 1},
		DataList: []*wspush.ServerMessage{wsMsg},
	}
	req.Timeout = timeout
	req.createTime = time.Now()
	req.Callback = callback

	// if buffer full, return error immediate
	select {
	case c.reqWriter <- req:
	default:
		return ErrWriteChannelFullMsgDropped
	}
	return nil
}

func (c *connectionImpl) SendNotify(data *wspush.ServerMessage) error {
	buffer, err := json.Marshal(data)
	if err != nil {
		return err
	}
	return c.Write(buffer)
}

func (c *connectionImpl) setConnPoolKey(key string) {
	c.connPoolKey = key
}

func (c *connectionImpl) getConnPoolKey() string {
	return c.connPoolKey
}

func (c *connectionImpl) onHeartBeat(pingMsg []byte) error {
	err := c.conn.WriteControl(websocket.PongMessage, pingMsg, time.Now().Add(KWriteWaitTime))
	if err == websocket.ErrCloseSent {
		return nil
	} else if e, ok := err.(net.Error); ok && e.Timeout() {
		c.logger.Debug("user connection timeout")
		return nil
	}
	c.heartbeatCh <- struct{}{}
	return err
}

func (c *connectionImpl) cleanup() {
	close(c.reqWriter)
	close(c.notifyWriter)
	close(c.heartbeatCh)
}

func (c *connectionImpl) onWriteError(err error) {
	c.logger.Warn("user connection write data failed.", zap.Error(err))
}

func (c *connectionImpl) onHandleAckMessage(wsMessage *wspush.C2SWebsocketMessage) {
	c.logger.Debug("receive ack msg", zap.Int32("seq", wsMessage.Header.Seq),
		zap.Any("wsMsgType", wsMessage.Body.Type))

	var (
		ok                 = false
		ackSeq             = wsMessage.Header.Seq
		item   *MsgRequest = nil
	)

	c.ackMutex.Lock()
	item, ok = c.waitAckMsg[ackSeq]
	if ok {
		delete(c.waitAckMsg, ackSeq)
	}
	c.ackMutex.Unlock()

	if ok {
		if item.Callback != nil {
			item.Callback(&Response{
				Req:       item,
				IsSuccess: true,
				ErrReason: nil,
			})
		}
	} else {
		c.logger.Warn("receive invalid ack msg", zap.Int32("seq", wsMessage.Header.Seq),
			zap.Any("wsMsgType", wsMessage.Body.Type))
	}
}

func (c *connectionImpl) OnLoop() {
	defer c.cleanup()

	// ping/pong
	c.conn.SetPingHandler(func(message string) error {
		return c.onHeartBeat([]byte(message))
	})

	// write routine
	go rescue.WithRecover(func() {
		c.onWrite()
	})

	for {
		msgType, data, err := c.conn.ReadMessage()
		if err != nil {
			c.logger.Debug("Connection closed", zap.Error(err))
			break
		}

		if msgType != websocket.BinaryMessage {
			c.logger.Warn("invalid msgType, close the connection", zap.Int("msgType", msgType))
			break
		}

		var wsMessage = wspush.C2SWebsocketMessage{}
		err = proto.Unmarshal(data, &wsMessage)
		if err != nil {
			c.logger.Error("proto.Unmarshal error", zap.Error(err))
			continue
		}

		c.lastHeartbeatTime = time.Now()
		if wsMessage.Body.Type == wspush.ClientWSMsgType_AckMessage {
			c.onHandleAckMessage(&wsMessage)
			continue
		}

		// 同步回调，耗时操作如redis更新等，应当自己起routine 处理
		if err = c.onHandle(context.Background(), &wsMessage); err != nil {
			c.logger.Warn("onHandle message error", zap.Any("clientMsgType", wsMessage.Body.Type), zap.Error(err))
		}
	}
}

func (c *connectionImpl) onWrite() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	msgBufferSize := 0
	msgBuffers := make([]*MsgRequest, 0)
	batchSendTicker := time.NewTicker(KBatchSendTimeInterval)
	defer batchSendTicker.Stop()

	sendBatch := func(requests []*MsgRequest) {
		c.logger.Debug("sendBatch", zap.Int("item size", len(requests)))

		// merge
		var mergeReq *MsgRequest
		for k := range requests {
			if mergeReq == nil {
				mergeReq = &MsgRequest{
					Data:       requests[k].Data,
					Timeout:    requests[k].Timeout,
					Callback:   requests[k].Callback,
					createTime: requests[k].createTime,
				}
				copy(mergeReq.Data.DataList, requests[k].Data.DataList)
			} else {
				mergeReq.Data.DataList = append(mergeReq.Data.DataList, requests[k].Data.DataList...)
			}
			// release MsgRequest
			//c.msgPool.Put(requests[k])
		}

		// overwrite seq
		seq := c.seq.Inc()
		mergeReq.Data.Header.Seq = seq

		buffer, err := proto.Marshal(mergeReq.Data)
		if err != nil {
			c.logger.Warn("onWrite Marshal error", zap.Error(err))
			return
		}

		_ = c.conn.SetWriteDeadline(time.Now().Add(KWriteWaitTime))
		if err = c.conn.WriteMessage(websocket.BinaryMessage, buffer); err != nil {
			_ = c.conn.Close()
			c.onWriteError(err)
			return
		} else {
			// insert to wait ack list
			c.ackMutex.Lock()
			c.waitAckMsg[seq] = mergeReq
			c.ackMutex.Unlock()
		}
	}

	for {
		select {
		case _, ok := <-c.heartbeatCh:
			if !ok {
				return
			}
			c.lastHeartbeatTime = time.Now()
			ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
			c.onDeviceHeartbeat(ctx)
			cancel()

		case buffer, ok := <-c.notifyWriter:
			if !ok {
				return
			}
			_ = c.conn.SetWriteDeadline(time.Now().Add(KWriteWaitTime))
			if err := c.conn.WriteMessage(websocket.BinaryMessage, buffer); err != nil {
				_ = c.conn.Close()
				c.onWriteError(err)
				return
			}

		case msg, ok := <-c.reqWriter:
			if !ok {
				return
			}

			msgBuffers = append(msgBuffers, msg)
			msgBufferSize += proto.Size(msg.Data)
			if msgBufferSize >= KBatchSendPacketMaxSize {
				c.logger.Info("msg buffer is full, sendBatch immediately")
				sendBatch(msgBuffers)
				msgBuffers = msgBuffers[:0]
				msgBufferSize = 0
			}

		case <-batchSendTicker.C:
			if len(msgBuffers) > 0 {
				sendBatch(msgBuffers)
				msgBuffers = msgBuffers[:0]
				msgBufferSize = 0
			}

		case <-ticker.C:
			if err := c.onTimer(); err != nil {
				c.conn.Close()
				c.logger.Debug(err.Error())
				return
			}
		}
	}
}

func (c *connectionImpl) onTimer() error {
	// Checking heartbeat timeout
	if time.Now().Sub(c.lastHeartbeatTime) > KHeartbeatTimeOut {
		return ErrHeartBeatTimeout
	}

	// check ack timeout
	timeoutSeq := make([]int32, 0)
	c.ackMutex.RLock()
	for _, v := range c.waitAckMsg {
		if time.Now().Sub(v.createTime) >= KMaxWaitAckTime {
			timeoutSeq = append(timeoutSeq, v.Data.Header.Seq)
		}
	}
	c.ackMutex.RUnlock()

	// remove timeout msg, insert to table store
	var (
		req *MsgRequest
	)
	for _, v := range timeoutSeq {
		c.ackMutex.Lock()
		req = c.waitAckMsg[v]
		delete(c.waitAckMsg, v)
		c.ackMutex.Unlock()

		if req.Callback != nil {
			req.Callback(&Response{
				Req:       req,
				IsSuccess: false,
				ErrReason: ErrReasonAckTimeout,
			})
		}
	}

	return nil
}
