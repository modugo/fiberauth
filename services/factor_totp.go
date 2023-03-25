package services

import "context"

type TOTPRepository interface {
	// IsEnabled check if TOTP is already enabled on this login
	IsEnabled(ctx context.Context, login string) bool

	// GetSharedSecretKey retrieves the shared secret key based on a login string
	GetSharedSecretKey(ctx context.Context, login string) string

	// Enable allows to enable TOTP
	Enable(ctx context.Context, key string)

	// Disable allows to disable TOTP
	Disable(ctx context.Context, login string)

	// StoreRecoveryCodes stores a list of recovery codes,
	// which can be used to recover access to the account
	// in case of loss or damage to the primary authentication mechanism.
	StoreRecoveryCodes(ctx context.Context, codes []string)
}
