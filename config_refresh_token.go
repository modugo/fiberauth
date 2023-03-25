package fiberauth

import (
	"github.com/modugo/fiberauth/services"
	"time"
)

type RefreshTokenConfig struct {
	IsEnabled  bool
	Duration   time.Duration
	Repository services.RefreshTokenRepository
}

func (r RefreshTokenConfig) GetIsEnabled() bool {
	return r.IsEnabled
}

func (r RefreshTokenConfig) GetDuration() time.Duration {
	return r.Duration
}

func (r RefreshTokenConfig) GetRepository() services.RefreshTokenRepository {
	return r.Repository
}
