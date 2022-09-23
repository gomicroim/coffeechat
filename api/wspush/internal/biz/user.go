package biz

import (
	"chat/api/chat"
	"context"
	"fmt"
	"go.uber.org/atomic"
	"sync"
	"wspush/api/wspush"
)

type ClientInfo struct {
	UserId     int64
	DeviceId   string
	ClientType chat.IMClientType
	Domain     string
	AppVersion int32
}

type User interface {
	Add(connection Connection) int
	Remove(connection Connection) int

	Connections() int
	BroadcastMsg(ctx context.Context, msg *wspush.ServerMessage)

	SetOffline(ctx context.Context)
	SetOnline(ctx context.Context) bool
}

type user struct {
	mu          sync.RWMutex
	connections map[string]Connection
	userId      int64

	connPoolCounter atomic.Int32
}

func NewUser(userId int64) User {
	return &user{userId: userId, connections: make(map[string]Connection, 0)}
}

func (u *user) Add(connection Connection) int {
	deviceId := connection.ClientInfo().DeviceId
	key := deviceId

	// 避免设备ID相同，无法建立多个ws连接
	if rawConn, ok := connection.(*connectionImpl); ok {
		key = u.buildConnectionKey(u.connPoolCounter.Inc(), deviceId)
		rawConn.setConnPoolKey(key)
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	u.connections[key] = connection

	return len(u.connections)
}

func (u *user) Remove(connection Connection) int {
	key := connection.ClientInfo().DeviceId
	if rawConn, ok := connection.(*connectionImpl); ok {
		key = rawConn.getConnPoolKey()
	}

	u.mu.Lock()
	defer u.mu.Unlock()

	delete(u.connections, key)
	return len(u.connections)
}

func (u *user) buildConnectionKey(poolIndex int32, deviceId string) string {
	key := fmt.Sprintf("%d-%s", poolIndex, deviceId)
	return key
}

func (u *user) Connections() int {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return len(u.connections)
}

func (u *user) BroadcastMsg(ctx context.Context, msg *wspush.ServerMessage) {
	u.mu.RLock()
	defer u.mu.RUnlock()

	for _, conn := range u.connections {
		_ = conn.SendRequest(msg, KMaxWaitAckTime, nil)
	}
}

func (u *user) SetOffline(ctx context.Context) {

}

func (u *user) SetOnline(ctx context.Context) bool {
	return true
}
