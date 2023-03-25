package services

import (
	"github.com/modugo/fiberauth/services"
	"time"
)

type Configer interface {
	Init() Configer
	GetAccessToken() AccessTokenConfiger
	GetRefreshToken() RefreshTokenConfiger
	GetFactors() FactorsConfiger
	GetRegister() RegisterConfiger
	GetWithCookie() bool
	GetWithLogout() bool
}

type AccessTokenConfiger interface {
	GetSecretKey() []byte
	GetDuration() time.Duration
	GetRepository() services.AccessTokenRepository
}

type RefreshTokenConfiger interface {
	GetIsEnabled() bool
	GetDuration() time.Duration
	GetRepository() services.RefreshTokenRepository
}

type RegisterConfiger interface {
	GetIsEnabled() bool
	GetRepository() services.RegisterRepository
}

type FactorsConfiger interface {
	Init() FactorsConfiger
	GetIsEnabled() bool
	GetSecretKey() []byte
	GetDuration() time.Duration
	GetRepository() services.FactorsRepository
	GetTOTP() TOTPConfiger
	GetSMS() SMSConfiger
}

type TOTPConfiger interface {
	GetIsEnabled() bool
	GetIssuer() string
	GetRepository() services.TOTPRepository
}

type SMSConfiger interface {
	GetIsEnabled() bool
	GetRepository() services.SMSRepository
}
