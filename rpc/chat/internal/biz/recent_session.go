package biz

import (
	"chat/internal/data"
	"context"
	"strconv"
)

type RecentSessionUseCase struct {
	sessionRepo data.SessionRepo
}

func NewRecentSessionUseCase(repo data.SessionRepo) *RecentSessionUseCase {
	return &RecentSessionUseCase{sessionRepo: repo}
}

func (r *RecentSessionUseCase) GetSessionList(ctx context.Context, userId int64) ([]*data.Session, error) {
	return r.sessionRepo.ListAll(ctx, strconv.FormatInt(userId, 10))
}
