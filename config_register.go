package fiberauth

import (
	"github.com/modugo/fiberauth/services"
)

type RegisterConfig struct {
	IsEnabled  bool
	Repository services.RegisterRepository
}

func (r RegisterConfig) GetIsEnabled() bool {
	return r.IsEnabled
}

func (r RegisterConfig) GetRepository() services.RegisterRepository {
	return r.Repository
}
