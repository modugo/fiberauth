package fiberauth

import "github.com/modugo/fiberauth/services"

type SMSConfig struct {
	IsEnabled  bool
	Repository services.SMSRepository
}

func (s SMSConfig) GetIsEnabled() bool {
	return s.IsEnabled
}

func (s SMSConfig) GetRepository() services.SMSRepository {
	return s.Repository
}
