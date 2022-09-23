package biz

import (
	"context"
	"wspush/api/wspush"
)

type MsgHandler func(ctx context.Context, data *wspush.C2SWebsocketMessage) error

func (c *connectionImpl) onHandle(ctx context.Context, msg *wspush.C2SWebsocketMessage) error {
	switch msg.Body.Type {
	case 1:
		return c.onHandleBackgroundActiveMessage(ctx, msg)
	case 2:
		return c.onHandleBackgroundActiveMessage(ctx, msg)
	default:
		return errorUnknownClientMessage
	}
}

// app 从后台进入前台
func (c *connectionImpl) onHandleFrontActiveMessage(ctx context.Context, data *wspush.C2SWebsocketMessage) error {
	c.logger.Debug("onHandleFrontActiveMessage")
	return nil
}

// app 进入后台
func (c *connectionImpl) onHandleBackgroundActiveMessage(ctx context.Context, data *wspush.C2SWebsocketMessage) error {
	c.logger.Debug("onHandleBackgroundActiveMessage")
	return nil
}

func (c *connectionImpl) onDeviceHeartbeat(ctx context.Context) {

}
