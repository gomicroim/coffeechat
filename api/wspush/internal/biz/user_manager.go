package biz

import (
	"context"
	"sync"
	"wspush/api/wspush"
)

// UserManager manger user<-->connection relations,
// all the function is thread-safe
type UserManager interface {
	Add(conn Connection) (User, int)
	Remove(connection Connection) int

	BroadcastMsgToUser(ctx context.Context, msg *wspush.ServerMessage, users []int64)
	BroadcastMsgToAllByAppVersion(ctx context.Context, msg *wspush.ServerMessage)
}

var (
	Users             = NewUserManager()
	_     UserManager = (*userManager)(nil)
)

type userManager struct {
	sync.RWMutex
	users map[int64]User
}

func NewUserManager() UserManager {
	return &userManager{
		users: make(map[int64]User),
	}
}

func (m *userManager) Add(conn Connection) (User, int) {
	uid := conn.ClientInfo().UserId

	m.Lock()
	defer m.Unlock()

	curUser, ok := m.users[uid]
	if !ok {
		curUser = NewUser(uid)
		m.users[uid] = curUser
	}

	connectCount := curUser.Add(conn)
	return curUser, connectCount
}

func (m *userManager) Remove(conn Connection) int {
	var remain int
	uid := conn.ClientInfo().UserId

	m.Lock()
	defer m.Unlock()

	curUser, ok := m.users[uid]
	if !ok {
		return remain
	}

	remain = curUser.Remove(conn)
	if remain <= 0 {
		delete(m.users, uid)
	}

	return remain
}

func (m *userManager) BroadcastMsgToUser(ctx context.Context, msg *wspush.ServerMessage, users []int64) {
	for _, uid := range users {
		m.RLock()
		curUser, ok := m.users[uid]
		m.RUnlock()

		if ok {
			curUser.BroadcastMsg(ctx, msg)
		}
	}
}

func (m *userManager) BroadcastMsgToAllByAppVersion(ctx context.Context, msg *wspush.ServerMessage) {
	m.RLock()
	defer m.RUnlock()

	for _, u := range m.users {
		u.BroadcastMsg(ctx, msg)
	}
}
