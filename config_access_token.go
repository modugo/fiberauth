package fiberauth

import (
	"github.com/modugo/fiberauth/services"
	"time"
)

type AccessTokenConfig struct {
	SecretKey  []byte
	Duration   time.Duration
	Repository services.AccessTokenRepository
}

func (a AccessTokenConfig) GetSecretKey() []byte {
	return a.SecretKey
}

func (a AccessTokenConfig) GetDuration() time.Duration {
	return a.Duration
}

func (a AccessTokenConfig) GetRepository() services.AccessTokenRepository {
	return a.Repository
}
