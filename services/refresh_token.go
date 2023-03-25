package services

import (
	"context"
	"time"
)

type RefreshTokenRepository interface {
	// IsValidRefreshToken lets know if refresh token is valid
	IsValidRefreshToken(ctx context.Context, token string) bool

	// DeleteRefreshToken allows to delete refresh token
	DeleteRefreshToken(ctx context.Context, token string)

	// ManageRefreshToken allows you to perform actions after
	// generating a refresh token (like store it)
	ManageRefreshToken(ctx context.Context, token string, exp *time.Time)
}
