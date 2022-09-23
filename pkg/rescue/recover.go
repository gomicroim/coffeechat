package rescue

import (
	"github.com/gomicroim/gomicroim/pkg/log"
	"go.uber.org/zap"
)

func WithRecover(fn func()) {
	defer func() {
		if err := recover(); err != nil {
			log.L.Error("recover", zap.Any("error", err))
		}
	}()

	fn()
}
