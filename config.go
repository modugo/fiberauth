package fiberauth

import "time"

type Config struct {
	Entity               Entity
	JWTSecret            []byte
	AccessTokenDuration  time.Duration
	RefreshTokenDuration time.Duration
	WithRefreshToken     bool
	WithCookie           bool
	WithSignIn           bool
	WithSignUp           bool
	WithRefresh          bool
	WithLogout           bool
}
