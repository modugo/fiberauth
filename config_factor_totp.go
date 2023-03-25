package fiberauth

import (
	"github.com/modugo/fiberauth/services"
)

type TOTPConfig struct {
	IsEnabled  bool
	Issuer     string
	Repository services.TOTPRepository
}

func (t TOTPConfig) GetIsEnabled() bool {
	return t.IsEnabled
}

func (t TOTPConfig) GetIssuer() string {
	return t.Issuer
}

func (t TOTPConfig) GetRepository() services.TOTPRepository {
	return t.Repository
}
