package services

import "context"

type RegisterRepository interface {
	// CreateAccount allows to create account,
	// the return must return the final object for response
	CreateAccount(ctx context.Context, payload map[string]interface{}) (p interface{}, err error)
}
