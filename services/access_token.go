package services

import (
	"context"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type AccessTokenRepository interface {
	// IsValidCredentials lets know if given credentials are valid
	IsValidCredentials(ctx context.Context, login, password string) bool

	// GetClaims return data how must be store in JWT token
	GetClaims(ctx context.Context, login string) jwt.MapClaims

	// IsValidAccessToken is an optional function that allows you
	// to add additional check for validate access token
	IsValidAccessToken(ctx context.Context, token string) bool

	// ManageAccessToken is an optional function that allows you
	// to perform actions after generating an access token (like store it)
	ManageAccessToken(ctx context.Context, token string, exp *time.Time)

	// SetCookie is an optional function that allows you to set cookie if
	// you enabled cookie
	SetCookie(ctx context.Context, name, value string, exp *time.Time) fiber.Cookie
}
