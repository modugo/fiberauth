package fiberauth

import (
	internalService "github.com/modugo/fiberauth/internal/services"
	"github.com/modugo/fiberauth/services"
	"time"
)

type FactorsConfig struct {
	IsEnabled  bool
	SecretKey  []byte
	Duration   time.Duration
	Repository services.FactorsRepository
	TOTP       internalService.TOTPConfiger
	SMS        internalService.SMSConfiger
}

func (f FactorsConfig) Init() internalService.FactorsConfiger {
	if f.TOTP == nil {
		f.TOTP = TOTPConfig{}
	}

	if f.SMS == nil {
		f.SMS = SMSConfig{}
	}

	return f
}

func (f FactorsConfig) GetIsEnabled() bool {
	return f.IsEnabled
}

func (f FactorsConfig) GetSecretKey() []byte {
	return f.SecretKey
}

func (f FactorsConfig) GetDuration() time.Duration {
	return f.Duration
}

func (f FactorsConfig) GetRepository() services.FactorsRepository {
	return f.Repository
}

func (f FactorsConfig) GetTOTP() internalService.TOTPConfiger {
	return f.TOTP
}

func (f FactorsConfig) GetSMS() internalService.SMSConfiger {
	return f.SMS
}
