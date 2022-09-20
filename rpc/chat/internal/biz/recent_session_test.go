package biz

import (
	"chat/internal/data"
	"context"
	"github.com/gomicroim/gomicroim/v2/pkg/log"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRecentSession_GetSessionList(t *testing.T) {
	dat, _ := setupBiz()
	session := NewRecentSessionUseCase(data.NewSessionRepo(dat, log.L))
	r, err := session.GetSessionList(context.Background(), 1)
	assert.NoError(t, err)
	t.Log(r)
}
