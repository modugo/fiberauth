package services

import "context"

type FactorsRepository interface {
	// IsEnabled check if a 2FA is enabled on this login
	IsEnabled(ctx context.Context, login string) bool

	// GetAvailableFactors return a list of active 2FA
	GetAvailableFactors(ctx context.Context, login string) []string

	// GetDefaultFactor return default factor for given login
	GetDefaultFactor(ctx context.Context, login string) string
}
