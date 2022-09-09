package jwt

import (
	"context"
)

const (
	kUserIdCtxKey     = "tkUserId"
	kDeviceIdCtxKey   = "tkDeviceId"
	kClientTypeCtxKey = "tkClientType"
	kDomainCtxKey     = "tkDomain"
)

// WithToken save client info to context
func WithToken(ctx context.Context, info *ClientInfo) context.Context {
	newCtx := context.WithValue(ctx, kUserIdCtxKey, info.UserId)
	newCtx = context.WithValue(newCtx, kDeviceIdCtxKey, info.DeviceId)
	newCtx = context.WithValue(newCtx, kClientTypeCtxKey, info.ClientType)
	newCtx = context.WithValue(newCtx, kDomainCtxKey, info.Domain)
	return newCtx
}

// FromJwtContext extract client info from jwt context
func FromJwtContext(ctx context.Context) ClientInfo {
	return ClientInfo{
		UserId:     GetUserId(ctx),
		DeviceId:   ctx.Value(kDeviceIdCtxKey).(string),
		ClientType: ctx.Value(kClientTypeCtxKey).(string),
		Domain:     ctx.Value(kDomainCtxKey).(string),
	}
}

// GetUserId try get userId from jwt context
func GetUserId(ctx context.Context) int64 {
	return ctx.Value(kUserIdCtxKey).(int64)
}
