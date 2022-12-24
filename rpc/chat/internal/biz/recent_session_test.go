package biz

import (
	"chat/internal/data"
	"context"
	"github.com/gomicroim/gomicroim/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecentSession_GetSessionList(t *testing.T) {
	_, dat := setupBiz(t)
	session := NewRecentSessionUseCase(data.NewSessionRepo(dat, log.L))
	r, err := session.GetSessionList(context.Background(), 1)
	assert.NoError(t, err)
	t.Log(r)
}
